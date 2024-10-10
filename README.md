# S3 Bucket Eventing

## Setup

```shell
oc new-app bucket-eventing
# fix values in ./openshift/pipe.yaml
oc apply -k ./openshift
```

## Docs

1. [S3 with Knative](https://knative.dev/blog/articles/consuming_s3_data_with_knative/)
2. [Knative Event Types](https://knative.dev/docs/eventing/event-registry/#about-eventtype-objects)
3. [Python Knative Function](https://developers.redhat.com/articles/2021/09/09/create-event-based-serverless-functions-python)
4. [Kamelet AWS S3 Source](https://camel.apache.org/camel-kamelets/4.8.x/aws-s3-source.html)
