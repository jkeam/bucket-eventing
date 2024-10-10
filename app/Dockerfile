FROM registry.access.redhat.com/ubi9/ubi-micro:9.4-15

WORKDIR /go/src/app
COPY ./bucket-eventing .

USER nobody
EXPOSE 8080
CMD ["./bucket-eventing"]
