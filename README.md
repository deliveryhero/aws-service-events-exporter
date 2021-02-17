# aws-events-exporter

![Version: 1.0.0](https://img.shields.io/badge/Version-1.0.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

Exports aws events to prometheus via aws SQS service

## How to install this chart

Add Delivery Hero public chart repo:

```console
helm repo add deliveryhero https://charts.deliveryhero.io/
```

A simple install with default values:

```console
helm install deliveryhero/aws-events-exporter
```

To install the chart with the release name `my-release`:

```console
helm install my-release deliveryhero/aws-events-exporter
```

To install with some set values:

```console
helm install my-release deliveryhero/aws-events-exporter --set values_key1=value1 --set values_key2=value2
```

To install with custom values file:

```console
helm install my-release deliveryhero/aws-events-exporter -f values.yaml
```
