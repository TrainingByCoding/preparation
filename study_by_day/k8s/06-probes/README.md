# k8s/06: Liveness & Readiness Probes

**Liveness** — is the container still alive? If fails → restart container.  
**Readiness** — is the container ready to serve traffic? If fails → removed from Service endpoints (no traffic).

```yaml
spec:
  containers:
  - name: app
    image: my-app
    livenessProbe:
      httpGet:
        path: /healthz
        port: 8080
      initialDelaySeconds: 10   # wait before first check
      periodSeconds: 5          # check every 5s
      failureThreshold: 3       # restart after 3 fails

    readinessProbe:
      httpGet:
        path: /ready
        port: 8080
      initialDelaySeconds: 5
      periodSeconds: 3
```

**Other probe types:**
```yaml
# TCP check
tcpSocket:
  port: 5432

# Command check
exec:
  command: ["pg_isready", "-U", "postgres"]
```

## Exercises
1. Deploy app with liveness probe on `/healthz`. After 30s, exec in and `kill 1` — watch restart.
2. Set liveness probe to a path that returns 500 — watch CrashLoopBackOff accumulate.
3. Make readiness probe fail (wrong path) — pod is Running but gets 0 traffic. Verify with `kubectl get endpoints`.
4. Set `initialDelaySeconds: 0` on a slow-starting app — watch it get killed before it's ready. Fix with higher delay.

## Interview facts
- Pod restarts don't change IP (it's the same pod). Deployment creates a new pod with new IP on crash.
- CrashLoopBackOff = pod keeps failing liveness/startup and restarting (backoff grows: 10s, 20s, 40s...)
- `startupProbe` exists for slow-starting apps — disables liveness until startup succeeds
