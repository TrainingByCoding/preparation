# k8s/02: Deployments

Deployment = manages a ReplicaSet → keeps N pods running, handles rolling updates.

## deployment.yaml
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: my-app
  template:                    # this is a Pod template
    metadata:
      labels:
        app: my-app
    spec:
      containers:
      - name: app
        image: nginx:1.25
        ports:
        - containerPort: 80
```

## Key commands
```bash
kubectl apply -f deployment.yaml
kubectl get deployments
kubectl get pods                         # see 3 pods
kubectl scale deployment my-app --replicas=5
kubectl set image deployment/my-app app=nginx:1.26   # rolling update
kubectl rollout status deployment/my-app
kubectl rollout undo deployment/my-app               # rollback
kubectl rollout history deployment/my-app
```

## Exercises
1. Apply deployment with 3 replicas. Kill one pod manually — watch it restart.
2. Scale to 5. Scale back to 2.
3. Update image to `nginx:1.26` — watch rolling update with `kubectl rollout status`
4. Rollback with `kubectl rollout undo`
5. Deliberately use a bad image (`nginx:doesnotexist`) — watch pods fail. Rollback.

## Interview facts
- Deployment → ReplicaSet → Pods (3 layers)
- Rolling update: brings up new pods before killing old (zero downtime)
- `maxSurge` / `maxUnavailable` control update speed
- selector.matchLabels MUST match template.labels — immutable after creation
