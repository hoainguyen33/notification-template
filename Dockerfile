FROM golang:alpine AS builder

ENV GIN_MODE=debug
ENV PORT=8000

RUN apk update && apk add --no-cache git
WORKDIR notificaiton-template

COPY . .

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o app-exe ./cmd/...

# FROM scratch

# COPY --from=builder /app-exe /go/bin/app-exe

EXPOSE $PORT

ENTRYPOINT ["./app-exe"]
