# k8s/04: ConfigMaps & Secrets

ConfigMap = non-sensitive config. Secret = sensitive data (base64 encoded, not encrypted by default).

## configmap.yaml
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
data:
  APP_ENV: production
  LOG_LEVEL: info
  config.json: |
    {"timeout": 30, "retries": 3}
```

## secret.yaml
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: app-secret
type: Opaque
stringData:              # kubectl auto base64-encodes these
  DB_PASSWORD: mysecretpass
  API_KEY: abc123
```

## Inject into pod
```yaml
spec:
  containers:
  - name: app
    image: nginx
    envFrom:
    - configMapRef:
        name: app-config       # all keys become env vars
    - secretRef:
        name: app-secret
    volumeMounts:
    - name: config-vol
      mountPath: /etc/config
  volumes:
  - name: config-vol
    configMap:
      name: app-config         # mounts config.json as a file
```

## Key commands
```bash
kubectl apply -f configmap.yaml
kubectl apply -f secret.yaml
kubectl exec -it <pod> -- env | grep APP_ENV
kubectl exec -it <pod> -- cat /etc/config/config.json
kubectl get secret app-secret -o jsonpath='{.data.DB_PASSWORD}' | base64 -d
```

## Exercises
1. Create ConfigMap + Secret. Deploy a pod that uses both as env vars. `exec` and `env` to verify.
2. Mount the ConfigMap as a volume. Verify file exists at `/etc/config/config.json`.
3. Update the ConfigMap. Pod needs a restart to pick up env changes (volume changes are live).
4. Create a secret imperatively: `kubectl create secret generic db-creds --from-literal=password=abc`

## Interview facts
- Secrets are base64, NOT encrypted at rest by default (use etcd encryption or Sealed Secrets in prod)
- Env vars from ConfigMap/Secret don't update live — need pod restart
- Volume-mounted ConfigMaps DO update live (within ~1 min)
