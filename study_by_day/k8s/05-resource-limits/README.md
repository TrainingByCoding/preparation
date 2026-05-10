# k8s/05: Resource Limits & Requests

```yaml
resources:
  requests:           # minimum guaranteed — used for scheduling
    cpu: "100m"       # 100 millicores = 0.1 CPU
    memory: "128Mi"
  limits:             # hard cap
    cpu: "500m"
    memory: "256Mi"
```

**What happens when limits are exceeded:**
- CPU: throttled (slowed down, not killed)
- Memory: OOMKilled (pod is killed immediately)

## Exercises
1. Deploy pod with memory limit 64Mi. Run `stress --vm 1 --vm-bytes 100M` inside — watch OOMKilled.
2. Set CPU request=1000m on a node with 1 CPU. Add a second pod — it'll be Pending (insufficient CPU).
3. `kubectl top pods` (requires metrics-server) — see actual CPU/memory usage.
4. Set no limits — then simulate memory leak — watch node pressure.

```bash
# Install stress tool in a pod
kubectl run stress-test --image=polinux/stress --limits='memory=64Mi' \
  --command -- stress --vm 1 --vm-bytes 100M --vm-keep

kubectl get pod stress-test   # watch for OOMKilled
kubectl describe pod stress-test | grep -A5 "Last State"
```

## Interview facts
- Request = what scheduler uses to place pods on nodes
- Limit = runtime enforcement by kubelet/cgroup
- OOMKilled exit code = 137
- `LimitRange` can set defaults per namespace
- `ResourceQuota` caps total usage per namespace
