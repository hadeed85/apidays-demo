# Go HTTP/3 File Server

This project is a static file server written in Go with support for HTTP/3 (QUIC) and HTTP/2, using the quic-go library. In the context of this demo is used to host the react ui application.

## Prerequisites:
- Go 1.20+ installed
- mkcert for generating local trusted certificates

## Setup:

1. Generate a self-signed certificate:

```sh
   mkcert "*.app.lan" "*.app.test" "*.app.invalid" "*.app.example"
```
   This will generate certificate and key files (e.g., _wildcard.app.lan+3.pem and _wildcard.app.lan+3-key.pem).

2. Update your /etc/hosts:
   Add an entry to resolve your test domain to localhost:
   127.0.0.1 aircraft-seat-reservation.app.lan

   Test with:

```sh   
   ping -c 1 aircraft-seat-reservation.app.lan
```
# Running the Server:
```sh
   go run server.go -address=0.0.0.0 -port=4445 -cert=_wildcard.app.lan+3.pem -key=_wildcard.app.lan+3-key.pem -wwww=../aircraft-seats-service-sse-ui/build
```
The server will be available at: https://aircraft-seat-reservation.app.lan:4445/
   
