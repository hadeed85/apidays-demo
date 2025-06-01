#!/bin/bash

go run ./go-http3-fileserver/server.go &
go run ./go-graphql-gateway/server.go -address=0.0.0.0 -port=4444 -cert=_wildcard.app.lan+3.pem -key=_wildcard.app.lan+3-key.pem &
go run ./aircraft-seats-service-go/server/main.go &
cd ./aircraft-seats-service-sse-ui/ && npm install && npm run start



