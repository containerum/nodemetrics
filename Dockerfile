FROM golang:1.10-alpine as builder
COPY . $GOPATH/src/github.com/containerum/nodeMetrics
WORKDIR $GOPATH/src/github.com/containerum/nodeMetrics
RUN go build -v -o /nodeMetrics ./cmd/nodeMetrics

FROM alpine:3.8 as runner
COPY --from=builder /nodeMetrics /nodeMetrics

ARG INFLUX_ADDR_ARG="http://localhost:8086"
ENV INFLUX_ADDR $INFLUX_ADDR_ARG

ARG PROMETHEUS_ADDR_ARG="http://localhost:9090"
ENV PROMETHEUS_ADDR $PROMETHEUS_ADDR_ARG

ARG SERVING_ADDR_ARG="localhost:8090"
ENV SERVING_ADDR $SERVING_ADDR_ARG

ENV USERNAME ""
ENV PASSWORD ""

CMD /nodeMetrics \
    -prometheus-addr $SERVING_ADDR \
    -serving-addr $SERVING_ADDR \
    -username $USERNAME \
    -password $PASSWORD
