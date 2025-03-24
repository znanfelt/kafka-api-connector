# Configuration

Values come from `config.yaml`, env vars, or Helm chart.

## Example

```yaml
kafka_broker: localhost:9092
api_endpoint: https://api.example.com
schema_path: ./schema.json
```

Secrets can be set as:
- `file://path/to/secret`
- `$ENV_VAR`
