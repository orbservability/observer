# Observer

Meant to be paired with Pixie's [standalone PEM](https://github.com/pixie-io/pixie/tree/main/src/experimental/standalone_pem).

## Usage

Install via [Helm](https://github.com/orbservability/helm-charts/tree/main/charts/observer) if you're using Kubernetes.

If you're loading this manually, add your PxL script at $PXL_FILE_PATH, and point to the PEM via $PIXIE_URL.

## Development

```sh
docker compose build
docker compose run --rm go mod tidy
```

## Reading

Learn about the various tech powering this application:

- [Pixie](https://docs.px.dev/about-pixie/what-is-pixie/)
- [eBPF](https://ebpf.io/)
