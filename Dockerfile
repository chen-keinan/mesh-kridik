# Use an official golang runtime as a parent image
FROM golang:1.15-alpine as builder

ENV GO111MODULE=on

ADD . /src

RUN apk --no-cache add ca-certificates wget


WORKDIR /src/cmd/mesh

RUN wget https://github.com/chen-keinan/mesh-kridik/releases/download/v1.0.1/mesh-kridik_1.0.1_x64.tar.gz -O mesh-kridik.tar.gz
RUN tar xzf mesh-kridik.tar.gz

FROM golang:1.15-alpine

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /src/cmd/mesh/mesh-kridik .

CMD ["./mesh-kridik"]
