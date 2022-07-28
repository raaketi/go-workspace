FROM golang:1.18.4-alpine3.16 as builder

WORKDIR /go/src/Userapiservice
COPY . .

RUN  go build .

FROM alpine:3.7 as runtime

ENV PORT=8080
WORKDIR "/app"
RUN apk add wget curl
#HEALTHCHECK --interval=10s --timeout=5s CMD wget --no-verbose --tries=1 --spider http://localhost:${PORT}/users

HEALTHCHECK --interval=10s --timeout=5s CMD curl http://localhost:${PORT}/users

COPY --from=builder /go/src/Userapiservice .
CMD ["./Userapiservice"]
