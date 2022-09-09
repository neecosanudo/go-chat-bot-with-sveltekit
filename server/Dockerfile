ARG GO_VERSION=1.18.0

FROM golang:${GO_VERSION}-alpine AS builder

RUN go env -w GOPROXY=direct && \ 
  apk add --no-cache git && \
  apk --no-cache add ca-certificates && update-ca-certificates

WORKDIR /src

COPY [ "go.*", "./"]
RUN go mod download

COPY ["./", "./"]

RUN CGO_ENABLED=0 go build \
  -installsuffix 'static' \
  -o /go-chat-bot-with-sveltekit/server .

EXPOSE 5050

ENTRYPOINT [ "/go-chat-bot-with-sveltekit/server" ]