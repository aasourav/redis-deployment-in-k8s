for redis deployment we used [Ot-redis-operator](https://ot-redis-operator.netlify.app/) Operator

[Operator Source Code](https://github.com/OT-CONTAINER-KIT/redis-operator/tree/master)

This operator have multiple architecture
    * Replication
    * Standalone
    * Sentinel

There is an issue I have faced when I deploy it in DigitalOcean. 
The issue was
```text
Can't open or create append-only dir appendonlydir: Permission denied
```

I have found the solution from redis repo (not from OT-CONTAINER-KIT/redis-operator)
The solution link https://github.com/helm/charts/issues/5041#issuecomment-421608472

```sh
redis.master.securityContext.enabled=true
redis.master.securityContext.runAsUser=0
redis.master.securityContext.fsGroup=2000
```

So I modify my `securityContext`:
```yaml
  securityContext:
    runAsUser: 0
    fsGroup: 2000
```

I have created a Issues on OT-CONTAINER-KIT/redis-operator : https://github.com/OT-CONTAINER-KIT/redis-operator/issues/1115
