# k8s/03: Services

Service = stable DNS name + IP for a set of pods (selected by label). Pods come and go, Service stays.

## Types
| Type | Accessible from | Use case |
|------|----------------|----------|
| ClusterIP | Inside cluster only | Default. Internal microservice |
| NodePort | Node IP + port (30000-32767) | Dev/test, direct external access |
| LoadBalancer | Cloud LB (AWS ELB etc.) | Production external traffic |

## service.yaml (ClusterIP)
```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-app-svc
spec:
  selector:
    app: my-app          # matches pods with this label
  ports:
  - port: 80             # service port
    targetPort: 80       # container port
  type: ClusterIP
```

## Key commands
```bash
kubectl apply -f service.yaml
kubectl get svc
kubectl describe svc my-app-svc
# test from inside cluster:
kubectl run tmp --image=busybox --rm -it -- wget -qO- http://my-app-svc
# NodePort — find the port:
kubectl get svc my-app-svc -o=jsonpath='{.spec.ports[0].nodePort}'
```

## Exercises
1. Deploy `02-deployments/deployment.yaml` then apply ClusterIP service. Curl it from a busybox pod.
2. Change type to NodePort. Access via `localhost:<nodePort>` (if using minikube: `minikube service my-app-svc`)
3. Break the selector (change `app: my-app` to `app: wrong`) — curl fails. Fix it.
4. Add a second port (port 8080 → targetPort 80). Both should work.

## Interview facts
- Service does NOT know about pods directly — it watches Endpoints (automatically managed)
- `kubectl get endpoints my-app-svc` shows which pod IPs are currently behind the service
- kube-proxy implements the load balancing (round-robin by default)
- DNS: `my-app-svc.default.svc.cluster.local`
