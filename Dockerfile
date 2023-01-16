FROM docker.io/library/golang:1.19-alpine

WORKDIR /go/src/app

COPY go.* *.go ./

RUN go build -o /go/src/app/prometheus-demo

CMD [ "/go/src/app/prometheus-demo" ]
