// SPDX-License-Identifier:Apache-2.0

package tests

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/openperouter/openperouter/api/v1alpha1"

	"github.com/openperouter/openperouter/e2etests/pkg/k8s"
	"github.com/openperouter/openperouter/e2etests/pkg/k8sclient"
	"github.com/openperouter/openperouter/e2etests/pkg/openperouter"
)

var _ = FDescribe("Node Router Status", func() {
	const routerNamespace = openperouter.Namespace

	var nodes []corev1.Node
	var cs kubernetes.Interface

	BeforeEach(func() {
		cs = k8sclient.New()

		var err error
		nodes, err = k8s.GetNodes(cs)
		Expect(err).NotTo(HaveOccurred())
		Expect(nodes).NotTo(BeEmpty(), "should have at least one node")
	})

	It("should ensure RouterNodeConfigurationStatus per each node", func() {
		expected := newNodeStatuses(nodes, routerNamespace,
			withConditions(newUnknownConditions()))

		assertNodeStatusesMatchExpected(routerNamespace, expected)

		By("delete all CRs and verify they are recreated")
		Expect(Updater.Client().DeleteAllOf(context.Background(), &v1alpha1.RouterNodeConfigurationStatus{}, client.InNamespace(routerNamespace))).To(Succeed())

		assertNodeStatusesMatchExpected(routerNamespace, expected)
	})
})

func assertNodeStatusesMatchExpected(namespace string, expected []v1alpha1.RouterNodeConfigurationStatus) {
	GinkgoHelper()

	Eventually(func(g Gomega) {
		actual := &v1alpha1.RouterNodeConfigurationStatusList{}
		g.Expect(Updater.Client().List(context.Background(), actual, &client.ListOptions{Namespace: namespace})).To(Succeed())

		g.Expect(actual.Items).To(WithTransform(sanitizeNodeStatus, ConsistOf(expected)))
	}).
		WithTimeout(10 * time.Second).
		WithPolling(1 * time.Second).
		Should(Succeed())
}

func sanitizeNodeStatus(statuses []v1alpha1.RouterNodeConfigurationStatus) []v1alpha1.RouterNodeConfigurationStatus {
	var sanitized []v1alpha1.RouterNodeConfigurationStatus

	for i := range statuses {
		s := statuses[i].DeepCopy()
		s.ObjectMeta.CreationTimestamp = metav1.Time{}
		s.ObjectMeta.ManagedFields = nil
		s.ObjectMeta.ResourceVersion = ""
		s.ObjectMeta.Generation = 0
		s.ObjectMeta.UID = ""

		if s.Status != nil {
			for j := range s.Status.Conditions {
				s.Status.Conditions[j].LastTransitionTime = metav1.Time{}
			}
		}

		sanitized = append(sanitized, *s)
	}

	return sanitized
}

type opt func(*v1alpha1.RouterNodeConfigurationStatus)

func newNodeStatuses(nodes []corev1.Node, namespace string, opts ...opt) []v1alpha1.RouterNodeConfigurationStatus {
	var nodeStatuses []v1alpha1.RouterNodeConfigurationStatus

	for _, node := range nodes {
		nodeStatus := newNodeStatus(node, namespace)
		for _, opt := range opts {
			opt(nodeStatus)
		}
		nodeStatuses = append(nodeStatuses, *nodeStatus)
	}

	return nodeStatuses
}

func newUnknownConditions() []metav1.Condition {
	return []metav1.Condition{
		{
			Type:    "Ready",
			Status:  "Unknown",
			Reason:  "Unknown",
			Message: "Unknown status",
		},
		{
			Type:    "Degraded",
			Status:  "Unknown",
			Reason:  "Unknown",
			Message: "Unknown status",
		},
	}
}

func newNodeStatus(node corev1.Node, namespace string) *v1alpha1.RouterNodeConfigurationStatus {
	nodeOwnerRef := metav1.OwnerReference{
		APIVersion: "v1",
		Kind:       "Node",
		Name:       node.Name,
		UID:        node.UID,
	}
	return &v1alpha1.RouterNodeConfigurationStatus{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "openpe.openperouter.github.io/v1alpha1",
			Kind:       "RouterNodeConfigurationStatus",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:            node.Name,
			Namespace:       namespace,
			OwnerReferences: []metav1.OwnerReference{nodeOwnerRef},
		},
		Status: &v1alpha1.RouterNodeConfigurationStatusStatus{},
	}
}

func withConditions(conditions []metav1.Condition) opt {
	return func(s *v1alpha1.RouterNodeConfigurationStatus) {
		s.Status.Conditions = conditions
	}
}
