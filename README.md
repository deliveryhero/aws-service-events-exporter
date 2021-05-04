# aws-service-events-exporter

This tool allows exporting of the aws events of various AWS services to prometheus for obervability and alerting purposes


## Deployment

You can deploy this tool via the helm chart in this repo or use our public available chart and deliveryhero public repository

## Configuration
- sns topic
- sqs queue

## Exposed stats
This exporter exports counter metric as shown below.
```
aws_service_events{event_id="RDS-xx", event_message="Automated cluster snapshot created", event_source="xx",job="xx"}
```
