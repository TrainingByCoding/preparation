# k8s/09: Troubleshooting Real Issues

These are the actual problems you'll hit and be asked about in interviews.

---

## Quick Reference — 30 Essential Commands

| Problem | Check | Command |
|---------|-------|---------|
| **1. Pod is not running** | Check pod status | `kubectl get pods -A` |
| **2. Pod in CrashLoopBackOff** | Check pod logs | `kubectl logs <pod-name> -n <ns>` |
| **3. Previous container logs** | Check previous logs | `kubectl logs <pod-name> -n <ns> --previous` |
| **4. Pod is stuck in ContainerCreating** | Describe pod | `kubectl describe pod <pod-name> -n <ns>` |
| **5. ImagePullBackOff error** | Check events | `kubectl describe pod <pod-name> -n <ns> \| grep -i image`
| **6. Pod is Pending** | See why pod is pending | `kubectl describe pod <pod-name> -n <ns>` |
| **7. Insufficient resources (CPU/Memory)** | Check cluster resources | `kubectl top nodes` |
| **8. Node is Not Ready** | Check node status | `kubectl get nodes` |
| **9. High CPU/Memory usage** | Check usage | `kubectl top pods -A` |
| **10. Pod OOMKilled** | Check pod events | `kubectl describe pod <pod-name> -n <ns> \| grep -i oom` |
| **11. Pod keeps restarting** | Check restart count | `kubectl get pod <pod-name> -n <ns>` |
| **12. Service not reachable** | Check service | `kubectl get svc -A` |
| **13. DNS not resolving** | Check CoreDNS pods | `kubectl get pods -n kube-system \| grep -i coredns` |
| **14. DNS resolution failed from pod** | Test DNS from pod | `kubectl exec -it <pod-name> -n <ns> -- nslookup kubernetes.default` |
| **15. Ingress not working** | Check ingress | `kubectl get ingress -A` |
| **16. Ingress returns 404/502** | Check ingress controller logs | `kubectl logs -n ingress-nginx <kubectl get pods -n ingress-nginx \| head -1>` |
| **17. No endpoints found** | Check endpoints | `kubectl get endpoints <svc-name> -n <ns>` |
| **18. ConfigMap not found** | Check configmap | `kubectl get configmap <cm-name> -n <ns> -o yaml` |
| **19. Secret not found** | Check secret | `kubectl get secret <secret-name> -n <ns> -o yaml` |
| **20. Volume mount issues** | Check pod describe | `kubectl describe pod <pod-name> -n <ns> \| grep -i volume` |
| **21. PVC is Pending** | Check PVC | `kubectl get pvc -A` |
| **22. PV not available** | Check PV | `kubectl get pv` |
| **23. Storage full on node** | Check node disk | `kubectl describe node <node-name> \| grep -i "disk\|storage"` |
| **24. Network policy blocking access** | Check network policies | `kubectl get networkpolicy -A` |
| **25. Pod can't reach external service** | Test connectivity from pod | `kubectl exec -it <pod-name> -n <ns> -- curl -I http://<external-url>` |
| **26. Deployment rollback** | Check rollout history | `kubectl rollout history deployment <deploy-name> -n <ns>` |
| **27. Undo bad deployment** | Rollback deployment | `kubectl rollout undo deployment <deploy-name> -n <ns>` |
| **28. Check rollout status** | See rollout status | `kubectl rollout status deployment <deploy-name> -n <ns>` |
| **29. Troubleshoot quickly** | Get all events | `kubectl get events -A --sort-by=.metadata.creationTimestamp` |
| **30. Verify everything** | Get all resources | `kubectl get all -A` |

**Pro Tips:**
- Always check events first: `kubectl get events -n <namespace>`
- Logs are your best friend: `kubectl logs <pod> --previous` for crashed containers
- Describe is a superpower: `kubectl describe <resource> <name>`
- Stay calm — Kubernetes will make sense eventually!

**Handy Shortcuts:**
- `-A` = All namespaces
- `-n` = Specific namespace
- `-o wide` = More details
- `-o yaml` = Full YAML output
- `--watch` = Watch changes live

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
