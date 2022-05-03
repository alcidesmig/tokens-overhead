FROM golang:1.18 as builder

WORKDIR /tokens-overhead

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ./ /tokens-overhead

RUN CGO_ENABLED=0 GOOS=linux go install -ldflags "-w -s" .

ENTRYPOINT ["/go/bin/tokens-overhead"]
