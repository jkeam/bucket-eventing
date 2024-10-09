# S3 Bucket Eventing

## Setup

```shell
oc new-app bucket-eventing

# fix values in ./openshift/pipe.yaml

oc apply -k ./openshift

cd ./app
func subscribe --filter type=org.apache.camel.event.aws.s3.getObject --source bucket-eventing-broker
```
