package client

import (
	"crypto/tls"
	"net/http"
	"github.com/quic-go/quic-go/http3"
)

type GRPCClientFactory[T any] func(client *http.Client, url string) T

func NewGRPCClient[T any](skipVerify bool, factory GRPCClientFactory[T]) T {
	roundTripper := &http3.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: skipVerify,
		},
	}
	client := &http.Client{
		Transport: roundTripper,
	}

	return factory(client, "https://grpc-server.app.lan:4446")

}