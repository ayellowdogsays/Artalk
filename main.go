package main

import (
	"context"
	"net/http"
	"os"

	"github.com/artalkjs/artalk/v2/cmd"
	"github.com/artalkjs/artalk/v2/internal/pkged"
	"github.com/vercel/vercel-go/http"
)

//go:embed public/*
//go:embed i18n/*
//go:embed conf/*
var embedFS embed.FS

func init() {
	pkged.SetFS(embedFS)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app := cmd.New()
	app.SetPort("") // 禁用默认端口监听
	app.SetDomain(r.Host) // 设置当前域名

	// 将 HTTP 请求转发给 Artalk
	app.ServeHTTP(w, r)
}

func main() {
	// 本地开发时直接启动（`go run main.go`）
	if os.Getenv("VERCEL") == "" {
		app := cmd.New()
		if err := app.Launch(); err != nil {
			panic(err)
		}
		return
	}

	// Vercel 环境使用 Serverless 函数
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
