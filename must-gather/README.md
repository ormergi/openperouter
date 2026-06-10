# Collecting OpenPERouter related data

You can use the `oc adm must-gather` CLI command to collect information about your cluster, including openperotuer
components objects and logs.

To collect openperouter data with must-gather, you must specify the extra image using the `--image` option.
Example command line:
```bash
$ oc adm must-gather --image=quay.io/openperouter/must-gather:main
```
Artifacts example:
```bash
$ tree must-gather.local.1583622035375619718/
├── event-filter.html
├── must-gather.logs
├── quay-io-omergi0-gather-sha256-65a04bf8cf79aecee72ad7f268f034a864075d2f15c9854905b11ed79089fb62
│   ├── aggregated-discovery-apis.yaml
│   ├── aggregated-discovery-api.yaml
│   ├── event-filter.html
│   ├── gather.logs
│   ├── namespaces
│   │   └── openperouter-system
│   │       ├── apps
│   │       │   ├── daemonsets.yaml
│   │       │   ├── deployments.yaml
│   │       │   ├── replicasets.yaml
│   │       │   └── statefulsets.yaml
│   │       ├── autoscaling
│   │       │   └── horizontalpodautoscalers.yaml
│   │       ├── batch
│   │       │   ├── cronjobs.yaml
│   │       │   └── jobs.yaml
│   │       ├── core
│   │       │   ├── configmaps.yaml
│   │       │   ├── endpoints.yaml
│   │       │   ├── events.yaml
│   │       │   ├── persistentvolumeclaims.yaml
│   │       │   ├── pods.yaml
│   │       │   ├── replicationcontrollers.yaml
│   │       │   ├── secrets.yaml
│   │       │   └── services.yaml
│   │       ├── discovery.k8s.io
│   │       │   └── endpointslices.yaml
│   │       ├── networking.k8s.io
│   │       │   ├── ingresses.yaml
│   │       │   └── networkpolicies.yaml
│   │       ├── openperouter-system.yaml
│   │       ├── pods
│   │       │   ├── controller-4qfk6
│   │       │   │   ├── controller
│   │       │   │   │   └── controller
│   │       │   │   │       └── logs
│   │       │   │   │           ├── current.log
│   │       │   │   │           ├── previous.insecure.log
│   │       │   │   │           └── previous.log
│   │       │   │   └── controller-4qfk6.yaml
│   │       │   ├── controller-qps2z
│   │       │   │   ├── controller
│   │       │   │   │   └── controller
│   │       │   │   │       └── logs
│   │       │   │   │           ├── current.log
│   │       │   │   │           ├── previous.insecure.log
│   │       │   │   │           └── previous.log
│   │       │   │   └── controller-qps2z.yaml
│   │       │   ├── nodemarker-7cf554c5b8-r6hrv
│   │       │   │   ├── nodemarker
│   │       │   │   │   └── nodemarker
│   │       │   │   │       └── logs
│   │       │   │   │           ├── current.log
│   │       │   │   │           ├── previous.insecure.log
│   │       │   │   │           └── previous.log
│   │       │   │   └── nodemarker-7cf554c5b8-r6hrv.yaml
│   │       │   ├── router-9zg7w
│   │       │   │   ├── cp-frr-files
│   │       │   │   │   └── cp-frr-files
│   │       │   │   │       └── logs
│   │       │   │   │           ├── current.log
│   │       │   │   │           ├── previous.insecure.log
│   │       │   │   │           └── previous.log
│   │       │   │   ├── frr
│   │       │   │   │   └── frr
│   │       │   │   │       └── logs
│   │       │   │   │           ├── current.log
│   │       │   │   │           ├── previous.insecure.log
│   │       │   │   │           └── previous.log
│   │       │   │   ├── reloader
│   │       │   │   │   └── reloader
│   │       │   │   │       └── logs
│   │       │   │   │           ├── current.log
│   │       │   │   │           ├── previous.insecure.log
│   │       │   │   │           └── previous.log
│   │       │   │   └── router-9zg7w.yaml
│   │       │   └── router-mkjhm
│   │       │       ├── cp-frr-files
│   │       │       │   └── cp-frr-files
│   │       │       │       └── logs
│   │       │       │           ├── current.log
│   │       │       │           ├── previous.insecure.log
│   │       │       │           └── previous.log
│   │       │       ├── frr
│   │       │       │   └── frr
│   │       │       │       └── logs
│   │       │       │           ├── current.log
│   │       │       │           ├── previous.insecure.log
│   │       │       │           └── previous.log
│   │       │       ├── reloader
│   │       │       │   └── reloader
│   │       │       │       └── logs
│   │       │       │           ├── current.log
│   │       │       │           ├── previous.insecure.log
│   │       │       │           └── previous.log
│   │       │       └── router-mkjhm.yaml
│   │       └── policy
│   │           └── poddisruptionbudgets.yaml
│   ├── perouter_crds
│   │   ├── aggregated-discovery-apis.yaml
│   │   ├── aggregated-discovery-api.yaml
│   │   ├── event-filter.html
│   │   ├── namespaces
│   │   │   └── openperouter-system
│   │   │       └── openpe.openperouter.github.io
│   │   │           └── routernodeconfigurationstatuses
│   │   │               ├── pe-kind-control-plane.yaml
│   │   │               └── pe-kind-worker.yaml
│   │   └── timestamp
│   ├── perouter_frr
│   │   ├── dump.log
│   │   ├── openperouter-system_router-9zg7w_frr-bfd-peer.log
│   │   ├── openperouter-system_router-9zg7w_frr-bgp-ipv4.log
│   │   ├── openperouter-system_router-9zg7w_frr-bgp-ipv6.log
│   │   ├── openperouter-system_router-9zg7w_frr-bgp-l2vpn-evpn.log
│   │   ├── openperouter-system_router-9zg7w_frr-bgp-neighbors.log
│   │   ├── openperouter-system_router-9zg7w_frr-config.log
│   │   ├── openperouter-system_router-9zg7w_netns-ip-address.log
│   │   ├── openperouter-system_router-9zg7w_netns-ip-link.log
│   │   ├── openperouter-system_router-9zg7w_netns-ip-neigh.log
│   │   ├── openperouter-system_router-9zg7w_netns-ip-route.log
│   │   ├── openperouter-system_router-9zg7w_netns-ip-vrf.log
│   │   ├── openperouter-system_router-mkjhm_frr-bfd-peer.log
│   │   ├── openperouter-system_router-mkjhm_frr-bgp-ipv4.log
│   │   ├── openperouter-system_router-mkjhm_frr-bgp-ipv6.log
│   │   ├── openperouter-system_router-mkjhm_frr-bgp-l2vpn-evpn.log
│   │   ├── openperouter-system_router-mkjhm_frr-bgp-neighbors.log
│   │   ├── openperouter-system_router-mkjhm_frr-config.log
│   │   ├── openperouter-system_router-mkjhm_netns-ip-address.log
│   │   ├── openperouter-system_router-mkjhm_netns-ip-link.log
│   │   ├── openperouter-system_router-mkjhm_netns-ip-neigh.log
│   │   ├── openperouter-system_router-mkjhm_netns-ip-route.log
│   │   └── openperouter-system_router-mkjhm_netns-ip-vrf.log
│   └── timestamp
└── timestamp
```

## Build and publish  

You can use the project make targets for building and publishing the image: 
```bash
$ make gather-build 
$ make gather-push
```

Using private registry and custom tags:
```bash
IMG_REPO=example.com/asd IMG_TAG=latest make gather-build
```