package main

import (
	"net/http"
	"os"

	"github.com/goproxy/goproxy"
)

func main() {
	g := goproxy.New()
	g.GoBinEnv = append(
		os.Environ(),
		"GOPROXY=https://goproxy.cn,direct",
	)
	g.ProxiedSUMDBs = []string{"sum.golang.org https://sum.golang.google.cn"}
	http.ListenAndServe("0.0.0.0:8080", g)
}