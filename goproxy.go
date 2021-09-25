package main

import (
	"net/http"
	"os"

	"github.com/goproxy/goproxy"
	"github.com/goproxy/goproxy/cacher"
)

func main() {
	g := goproxy.New()
	g.GoBinEnv = append(
		os.Environ(),
		"GOPROXY=https://goproxy.cn,direct", // 使用 goproxy.cn 作为上游代理
	)
	g.ProxiedSUMDBs = []string{"sum.golang.org https://goproxy.cn/sumdb/sum.golang.org"} // 代理默认的校验和数据库
	g.Cacher = &cacher.Disk{
		Root: "/cache",
	}
	http.ListenAndServe("0.0.0.0:8080", g)
}
