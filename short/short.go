package main

import (
	"net/http"
	"net/url"
	"sync"

	"github.com/bwmarrin/snowflake"
	"github.com/valyala/fasthttp"
)

var (
	mp   sync.Map
	node *snowflake.Node
)

func requestHandler(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())
	if path == "/shortUrl" {
		handleShortUrl(ctx)
		return
	}

	p := string(ctx.URI().LastPathSegment())
	if v, ok := mp.Load(p); ok {
		ctx.Redirect(v.(string), http.StatusMovedPermanently)
		return
	}
	ctx.Error("Unsupported path", fasthttp.StatusNotFound)
}

func handleShortUrl(ctx *fasthttp.RequestCtx) {
	if !ctx.IsGet() {
		ctx.Error("Unsupported Method", http.StatusBadRequest)
		return
	}
	if !ctx.QueryArgs().Has("url") {
		ctx.Error("Param url Miss", http.StatusBadRequest)
		return
	}
	rawurl := ctx.QueryArgs().Peek("url")
	u, err := url.Parse(string(rawurl))
	if err != nil {
		ctx.Error("Invalid url", http.StatusBadRequest)
		return
	}
	shortLink := genShortUrl()
	mp.LoadOrStore(shortLink, u.String())
	ctx.WriteString(shortLink)
}

func genShortUrl() string {
	return node.Generate().Base58()
}

func init() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
}

func main() {
	fasthttp.ListenAndServe(":8080", requestHandler)
}
