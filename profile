- apiVersion: v1
  kind: profile
  metadata:
    name: backend
    tags:
    - backend
  spec:
    egress:
    - action: allow
      destination: {}
      source: {}
    ingress:
    - action: allow
      destination: {}
      source:
        tag: backend
- apiVersion: v1
  kind: profile
  metadata:
    name: frontend
    tags:
    - frontend
  spec:
    egress:
    - action: allow
      destination: {}
      source: {}
    ingress:
    - action: allow
      destination: {}
      source:
        tag: frontend
