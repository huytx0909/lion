FROM golang:1.14 as builder

WORKDIR /appbuilder

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make build

FROM centos:8

WORKDIR /app

COPY --from=builder /appbuilder/bin/lion /usr/bin/
