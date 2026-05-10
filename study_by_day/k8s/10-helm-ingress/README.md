# k8s/10: Helm & Ingress

## Helm — package manager for k8s

Helm chart = collection of YAML templates with variables.

```bash
# Install Helm
choco install kubernetes-helm    # Windows
brew install helm                # Mac

# Basic workflow
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
helm search repo nginx
helm install my-nginx bitnami/nginx           # install
helm list                                      # see installed releases
helm upgrade my-nginx bitnami/nginx --set replicaCount=3
helm rollback my-nginx 1                      # rollback to revision 1
helm uninstall my-nginx
```

**Create your own chart:**
```bash
helm create my-app           # generates chart skeleton
# edit my-app/values.yaml   ← your defaults
# edit my-app/templates/    ← your YAML with {{ .Values.xxx }}
helm install my-release ./my-app
helm install my-release ./my-app --set image.tag=1.26
```

---

## Ingress — HTTP routing into the cluster

Without Ingress: one LoadBalancer per service (expensive).  
With Ingress: one LoadBalancer → Ingress controller → route by host/path.

```bash
# Enable ingress on minikube
minikube addons enable ingress
```

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: my-ingress
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  rules:
  - host: myapp.local
    http:
      paths:
      - path: /api
        pathType: Prefix
        backend:
          service:
            name: api-svc
            port:
              number: 80
      - path: /
        pathType: Prefix
        backend:
          service:
            name: frontend-svc
            port:
              number: 80
```

```bash
# Add to hosts file for local testing (Windows: C:\Windows\System32\drivers\etc\hosts)
# 127.0.0.1  myapp.local

kubectl apply -f ingress.yaml
kubectl get ingress
curl http://myapp.local/api
```

## Exercises
1. Install nginx via Helm: `helm install my-nginx bitnami/nginx`. Check it's running.
2. Create your own chart for the Go app from 02-deployments. Parameterize `replicas` and `image.tag`.
3. Deploy two services (api + frontend). Create Ingress routing `/api` → api-svc, `/` → frontend-svc.
4. Add TLS to Ingress using a self-signed cert (cert-manager in prod).

## Interview facts
- Helm 3 (no Tiller) — runs client-side only
- `values.yaml` = defaults; override with `--set` or `-f custom-values.yaml`
- Ingress requires an Ingress Controller (nginx, traefik, etc.) — not built into k8s
- `pathType: Exact` vs `Prefix` matters — `/api` Exact won't match `/api/users`
