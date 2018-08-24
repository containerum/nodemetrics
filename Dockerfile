FROM golang:1.10-alpine as builder
RUN apk add --update make git
WORKDIR src/github.com/containerum/nodeMetrics
COPY . .
RUN VERSION=$(git describe --abbrev=0 --tags) make build-for-docker

FROM alpine:3.8 as runner
COPY --from=builder /tmp/nodeMetrics /

ARG INFLUX_ADDR_ARG="http://localhost:8086"
ENV INFLUX_ADDR $INFLUX_ADDR_ARG

ARG PROMETHEUS_ADDR_ARG="http://localhost:9090"
ENV PROMETHEUS_ADDR $PROMETHEUS_ADDR_ARG

ARG SERVING_ADDR_ARG="localhost:8090"
ENV SERVING_ADDR $SERVING_ADDR_ARG

ENV USERNAME ""
ENV PASSWORD ""

CMD /nodeMetrics \
    -prometheus-addr $PROMETHEUS_ADDR \
    -serving-addr $SERVING_ADDR \
    -username $USERNAME \
    -password $PASSWORD
