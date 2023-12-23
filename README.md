# tdarr_exporter

A prometheus exporter for the [Tdarr](https://home.tdarr.io/) distributed transcoding system.

## Configuration

All configuration is done via environment variables. See `.env-sample` for all available options. "Sensible defaults" are set for all options, so you really only need to set the `TDARR_HOST` variable to point to your Tdarr instance, eg `TDARR_HOST=http://tdarr.example.com:8265`. If running in Kubernetes, and assuming you've deployed this exporter in the same namespace as your Tdarr instance, you don't even need to set that, as it will default to `http://tdarr:8265`.

## Running

### Docker

```bash
docker run -d --name tdarr_exporter -p 9082:9082 -e TDARR_HOST=http://tdarr.example.com:8265 robertlestak/tdarr_exporter:latest
```

### Kubernetes

```bash
kubectl apply -f k8s/deploy.yaml
```