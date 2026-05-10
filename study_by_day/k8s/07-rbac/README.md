# k8s/07: RBAC

RBAC = Role-Based Access Control. Who can do what to which resources.

**4 objects:**
- `Role` — permissions in one namespace
- `ClusterRole` — permissions cluster-wide
- `RoleBinding` — binds Role to a user/serviceaccount in namespace
- `ClusterRoleBinding` — binds ClusterRole cluster-wide

```yaml
# Role: can only read pods in 'default' namespace
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: pod-reader
  namespace: default
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "list", "watch"]
---
# Bind it to a ServiceAccount
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: pod-reader-binding
  namespace: default
subjects:
- kind: ServiceAccount
  name: my-app-sa
  namespace: default
roleRef:
  kind: Role
  name: pod-reader
  apiGroup: rbac.authorization.k8s.io
---
# ServiceAccount the pod uses
apiVersion: v1
kind: ServiceAccount
metadata:
  name: my-app-sa
  namespace: default
```

## Key commands
```bash
kubectl auth can-i get pods --as=system:serviceaccount:default:my-app-sa
kubectl auth can-i delete pods --as=system:serviceaccount:default:my-app-sa   # no
kubectl create serviceaccount my-sa
kubectl create role pod-reader --verb=get,list --resource=pods
kubectl create rolebinding my-binding --role=pod-reader --serviceaccount=default:my-sa
```

## Exercises
1. Create SA + Role + RoleBinding from the YAML. Test with `kubectl auth can-i`.
2. Try to give it `delete` pods — test it's denied.
3. Attach the SA to a pod: `spec.serviceAccountName: my-app-sa`. Exec in and use the mounted token to query the API.
4. Create a ClusterRole that can read all nodes. Bind it to a SA.

## Interview facts
- Every pod gets a default ServiceAccount (with minimal permissions)
- Token is auto-mounted at `/var/run/secrets/kubernetes.io/serviceaccount/token`
- Principle of least privilege — give only what the pod needs
