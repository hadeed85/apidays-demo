
Prerequisites: [Buf](https://buf.build/) and [Go](https://golang.org/dl/).

### Generate a self-signed certificate

```sh
mkcert "*.app.lan" "*.app.test" "*.app.invalid" "*.app.example"
```

This will generate certificate and key files (e.g., `_wildcard.app.lan+3.pem` and `_wildcard.app.lan+3-key.pem`).

Then run: ```buf generate```

Followed by: ```go run ./server/main.go```