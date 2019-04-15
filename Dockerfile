FROM golang:1.12.4-alpine as builder

# standard workdir after building from golang image:
# WORKDIR = /go
#

WORKDIR /go/src/github.com/fraenky8/go-echo-server
COPY . .

RUN go build .

FROM alpine:3.9
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/github.com/fraenky8/go-echo-server/go-echo-server .

EXPOSE 80

CMD ["./go-echo-server", "80"]
