# Study Map

```
study_by_day/
├── go-core/          ← Go language fundamentals
│   ├── 01-pointers/
│   ├── 02-arrays-slices/
│   ├── 03-algorithms/
│   ├── 04-two-sum/
│   ├── 05-channels/
│   ├── 06-producer-consumer/
│   └── 07-worker-pool/
│
├── go-patterns/      ← Concurrency & design patterns
│   ├── 01-singleton/
│   ├── 02-factory/
│   ├── 03-fan-in/
│   ├── 04-fan-out/
│   ├── 05-pipeline/
│   ├── 06-circuit-breaker/
│   ├── 07-rate-limiter/
│   ├── 08-retry-backoff/
│   ├── 09-context-cancel/
│   └── 10-review/
│
├── dsa/              ← Data structures & algorithms
│   ├── 01-sliding-window/
│   ├── 02-two-pointers/
│   ├── 03-linked-list/
│   ├── 04-linked-list-advanced/
│   ├── 05-stack-queue/
│   ├── 06-binary-search/
│   ├── 07-trees/
│   ├── 08-trees-bst/
│   ├── 09-graphs/
│   └── 10-dp/
│
└── k8s/              ← Kubernetes
    ├── 01-pods/
    ├── 02-deployments/
    ├── 03-services/
    ├── 04-config-secrets/
    ├── 05-resource-limits/
    ├── 06-probes/
    ├── 07-rbac/
    ├── 08-storage/
    ├── 09-troubleshoot/
    └── 10-helm-ingress/
```

## Each folder has
- `README.md` — pattern + commands + exercises (no fluff)
- `practice.go` (Go) — TODOs to implement, solutions in comments at bottom
- `*.yaml` (k8s) — ready to apply against minikube

## Tools
- Go: already installed
- K8s: `minikube start` + `kubectl`
- Cross-device: Replit (replit.com, login with Google) — run code in browser, no install

## Quick start
```powershell
cd go-patterns\03-fan-in
code README.md    # read
code practice.go  # implement TODOs
go run practice.go
```
| [Day 28](day28/) | Final Project | Combine concepts | [ ] |

### Week 5: Interview Prep
| Day | Topic | Focus | Status |
|-----|-------|-------|--------|
| [Day 29](day29/) | Pattern Review | All concurrency patterns | [ ] |
| [Day 30](day30/) | Mock Interview | Timed problem-solving | [ ] |

