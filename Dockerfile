FROM golang:1.16 as builder

RUN mkdir -p $GOPATH/src/github.com/abdullokh-mukhammadjonov/template_api_gateway
WORKDIR $GOPATH/src/github.com/abdullokh-mukhammadjonov/template_api_gateway

COPY . ./

RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    make build && \
    mv ./bin/ek_admin_api_gateway /

FROM alpine
COPY --from=builder ek_admin_api_gateway .
ENTRYPOINT [ "/ek_admin_api_gateway" ]
