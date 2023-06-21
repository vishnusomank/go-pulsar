FROM golang:alpine3.18 as builder

WORKDIR /pulsar

COPY . .

RUN go build main.go

FROM alpine:3.18

WORKDIR /pulsar

RUN mkdir config

COPY --from=builder /pulsar/main .
COPY --from=builder /pulsar/kmux.yaml .
COPY --from=builder /pulsar/config/app.yaml ./config/app.yaml

ENTRYPOINT [ "sh","-c", "/pulsar/main > log.txt" ]

