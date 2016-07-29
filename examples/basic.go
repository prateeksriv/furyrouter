package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
	"github.com/gofury/furyrouter"
)

func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello, %s!\n", ctx.UserValue("name"))
}

func MultiParams(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hi, %s, %s!\n", ctx.UserValue("name"), ctx.UserValue("word"))
}

func main() {
	router := furyrouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/multi/:name/:word", MultiParams)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
