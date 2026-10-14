/*
Copyright 2026.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

// DualStackCIDRs is a list of at most one IPv4 and one IPv6 CIDR.
//
// +kubebuilder:validation:MinItems=1
// +kubebuilder:validation:MaxItems=2
// +kubebuilder:validation:XValidation:rule="self.all(c, isCIDR(c))",message="all entries must be valid CIDRs"
// +kubebuilder:validation:XValidation:rule="self.filter(c, isCIDR(c) && cidr(c).ip().family() == 4).size() <= 1",message="at most one IPv4 CIDR is allowed"
// +kubebuilder:validation:XValidation:rule="self.filter(c, isCIDR(c) && cidr(c).ip().family() == 6).size() <= 1",message="at most one IPv6 CIDR is allowed"
// +listType=atomic
type DualStackCIDRs []string
