# Kyma k8s native lambda POC

## Instructions

1. Apply ns.yaml first
1. Use `kubens` or any other tool to apply other yamls automatically to `lambda-test` namespace
1. Apply `sbu.yaml` and `trigger.yaml`
1. Update APIRule's spec.service.host in `istio-native-lambda.yaml` and `non-istio.yaml`
1. Apply either istio or non-istio yaml, test connection via gateway by link in APIRule
1. Look at lambda pod logs -> see REDIS_PORT being properly injected
1. Run: `k run -it --rm --restart=Never alpine --image=alpine sh`
1. `apk add curl`, next copy demo/curl.sh and execute this command
1. This sends event to event-mesh, which should be forwarded to lambda
1. Look at lambda logs to see that event hit lambda
1. `k delete -f demo/{previous yaml}` and apply the other one, either istio or non-istio one
1. Be happy that everything works wonderfully
1. If you want you can load-test lambda to trigger HPA using e.g. `fortio`:

```bash
fortio load -a -qps 0 -t 3m -c 800 -r 0.0001 {lambda-url}
```
