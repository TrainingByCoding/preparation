# k8s/09: Troubleshooting Real Issues

These are the actual problems you'll hit and be asked about in interviews.

---

## 1. CrashLoopBackOff
**Cause:** App keeps crashing. Backoff grows: 10s, 20s, 40s, 80s...

```bash
kubectl get pods                        # see CrashLoopBackOff
kubectl logs <pod>                      # current logs
kubectl logs <pod> --previous          # logs from crashed container
kubectl describe pod <pod>             # see exit code, events
```
**Common causes:** bad config, missing env var, OOMKilled, app bug, wrong entrypoint.

---

## 2. ImagePullBackOff / ErrImagePull
**Cause:** Can't pull the image.

```bash
kubectl describe pod <pod>  # look at Events section
```
**Common causes:**
- Wrong image name/tag → fix the image reference
- Private registry, no pull secret → `kubectl create secret docker-registry`
- No internet on node (air-gapped)

---

## 3. Pod Stuck in Pending
**Cause:** Scheduler can't find a node to place it.

```bash
kubectl describe pod <pod>  # Events: "Insufficient cpu/memory", "no nodes match"
kubectl get nodes           # are nodes Ready?
kubectl describe node <node>  # see Allocatable vs Requests
```
**Common causes:** not enough CPU/memory, nodeSelector doesn't match any node, PVC unbound.

---

## 4. Service Not Reachable
```bash
kubectl get svc                          # service exists?
kubectl get endpoints my-app-svc        # any IPs? Empty = selector mismatch
kubectl describe svc my-app-svc         # check selector
kubectl get pods --show-labels          # do pod labels match selector?
```
**Common cause:** `selector: app: my-app` in Service but pods have `app: myapp` (typo).

---

## 5. OOMKilled
```bash
kubectl describe pod <pod>  # Last State: OOMKilled, exit code 137
kubectl top pods            # see actual memory usage
```
**Fix:** Increase memory limit OR fix memory leak in app.

---

## 6. Rolling Update Stuck
```bash
kubectl rollout status deployment/my-app  # shows stuck
kubectl describe deployment my-app        # see events
```
**Cause:** New pods failing readiness probe → old pods never terminated.  
**Fix:** Fix the app/probe, or `kubectl rollout undo`.

---

## 7. Node NotReady
```bash
kubectl get nodes             # see NotReady
kubectl describe node <node>  # see conditions, events
kubectl cordon <node>         # stop scheduling new pods here
kubectl drain <node>          # evict pods (for maintenance)
```

---

## Practice Scenarios — do these with minikube

```bash
# Scenario 1: Break an app
kubectl set image deployment/my-app app=nginx:doesnotexist
# → ImagePullBackOff. Fix: rollout undo

# Scenario 2: OOM
kubectl run oom --image=polinux/stress --limits='memory=32Mi' -- stress --vm 1 --vm-bytes 64M
# → OOMKilled. Fix: increase limit

# Scenario 3: Selector mismatch
# Edit service selector to wrong label → endpoints empty → curl fails

# Scenario 4: Pending due to CPU
kubectl run hungry --image=nginx --requests='cpu=99'
# → Pending (no node has 99 CPUs). Fix: lower request
```
