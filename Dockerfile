FROM golang:1.21.13-alpine3.20

WORKDIR /app
COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg
COPY examples ./examples
COPY go.mod go.sum ./
EXPOSE 8088

RUN go install -C /app/cmd/wamock
ENTRYPOINT [ "wamock" ]
CMD /app/examples/single_files_config/hello.json