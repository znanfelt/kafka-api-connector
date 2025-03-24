# Helm Deployment

This app supports deployment via Helm.

```bash
helm install kafka-api ./charts/kafka-api
```

## Configurable values

See `values.yaml` for all options:
- `replicaCount`
- `resources`
- `envSecrets`
