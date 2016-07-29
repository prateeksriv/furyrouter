package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
	"github.com/gofury/furyrouter"
)

var basicAuthPrefix = []byte("Basic ")

func BasicAuth(h fasthttp.RequestHandler, user, pass []byte) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		// Get the Basic Authentication credentials
		auth := ctx.Request.Header.Peek("Authorization")
		if bytes.HasPrefix(auth, basicAuthPrefix) {
			// Check credentials
			payload, err := base64.StdEncoding.DecodeString(string(auth[len(basicAuthPrefix):]))
			if err == nil {
				pair := bytes.SplitN(payload, []byte(":"), 2)
				if len(pair) == 2 &&
					bytes.Equal(pair[0], user) &&
					bytes.Equal(pair[1], pass) {
					// Delegate request to the given handle
					h(ctx)
					return
				}
			}
		}

		// Request Basic Authentication otherwise
		ctx.Response.Header.Set("WWW-Authenticate", "Basic realm=Restricted")
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusUnauthorized), fasthttp.StatusUnauthorized)
	}
}

func NotProtected(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Not protected!\n")
}

func Protected(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Protected!\n")
}

func main() {
	user := []byte("gordon")
	pass := []byte("secret!")

	router := furyrouter.New()
	router.GET("/", NotProtected)
	router.GET("/protected/", BasicAuth(Protected, user, pass))

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
