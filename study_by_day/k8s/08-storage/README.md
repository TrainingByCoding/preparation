# k8s/08: Storage (PV, PVC, StorageClass)

**Problem:** Pod restarts = data lost. Solution: persistent volumes.

**Flow:** `StorageClass` → provisions `PersistentVolume (PV)` → claimed by `PersistentVolumeClaim (PVC)` → mounted in Pod.

```yaml
# PVC — just say what you need, k8s finds/creates a PV
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-data
spec:
  accessModes:
  - ReadWriteOnce        # only one node can write at a time
  resources:
    requests:
      storage: 1Gi
  storageClassName: standard   # matches StorageClass name
---
# Pod using the PVC
apiVersion: v1
kind: Pod
metadata:
  name: app-with-storage
spec:
  containers:
  - name: app
    image: nginx
    volumeMounts:
    - name: data
      mountPath: /data
  volumes:
  - name: data
    persistentVolumeClaim:
      claimName: my-data
```

## Access modes
| Mode | Meaning |
|------|---------|
| ReadWriteOnce (RWO) | One node, read+write |
| ReadOnlyMany (ROX) | Many nodes, read only |
| ReadWriteMany (RWX) | Many nodes, read+write (NFS, EFS) |

## Key commands
```bash
kubectl get pv,pvc
kubectl describe pvc my-data
# minikube has 'standard' StorageClass — auto-provisions PVs
kubectl get storageclass
```

## Exercises
1. Apply pvc.yaml + pod. Exec in, write file to `/data/test.txt`. Delete pod. Recreate pod — file still there.
2. Delete the PVC — PV may be retained (depends on reclaim policy). Check `kubectl get pv`.
3. Change `accessModes` to `ReadWriteMany` with standard storageClass — it'll fail (standard doesn't support RWX).
4. Deploy a StatefulSet with `volumeClaimTemplates` — each pod gets its own PVC automatically.

## Interview facts
- PVC is namespace-scoped; PV is cluster-scoped
- Reclaim policy: `Retain` (keep data), `Delete` (delete PV when PVC deleted), `Recycle` (deprecated)
- StatefulSets use `volumeClaimTemplates` — each replica gets its own PVC (pod-0, pod-1, ...)
