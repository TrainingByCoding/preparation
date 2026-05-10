# k8s/01: Pods

A Pod is the smallest deployable unit. One or more containers sharing network + storage.

## Key commands
```bash
kubectl run nginx --image=nginx              # create pod
kubectl get pods
kubectl describe pod nginx
kubectl exec -it nginx -- /bin/sh
kubectl logs nginx
kubectl delete pod nginx
```

## pod.yaml
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-app
  labels:
    app: my-app
spec:
  containers:
  - name: app
    image: golang:1.22
    command: ["sleep", "3600"]
    env:
    - name: ENV
      value: "dev"
```

## Exercises
1. Apply pod.yaml, exec into it, run `env` to see ENV variable
2. Change the image to `nginx`, apply, verify it serves a page: `kubectl exec -it my-app -- curl localhost`
3. Delete the pod — notice it does NOT restart (that's why you use Deployments)
4. Try `kubectl get pod my-app -o yaml` — read the full spec k8s generated
5. Add a second container to the pod (sidecar). Both share the same `localhost`.

## Key facts for interview
- Pod IP is ephemeral — changes on restart
- Containers in same pod share network namespace (same localhost, different ports)
- Pod alone has no self-healing — use Deployment for that
