FROM golang:1.18.5-alpine3.16 as build
 
WORKDIR /source
COPY . .
 
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o bin/main main.go
 
FROM alpine:3.12
 
COPY --from=build /source/bin/main /main
 
ENTRYPOINT ["/main"]