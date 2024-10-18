package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/redis/go-redis/v9"
)

var version string

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ping", pingHandler)

	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Printf("服务器启动失败: %v\n", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "version: %s\n", version)
	redisURL := os.Getenv("REDIS_URL")
	fmt.Fprintf(w, "redisURL: %s\n", redisURL)
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		fmt.Fprintf(w, "解析Redis URL失败: %v\n", err)
		http.Error(w, "内部服务器错误", http.StatusInternalServerError)
		return
	}

	client := redis.NewClient(opts)

	defer client.Close()

	ctx := r.Context()
	info, err := client.Info(ctx).Result()
	if err != nil {
		fmt.Fprintf(w, "获取Redis信息失败: %v\n", err)
		http.Error(w, "内部服务器错误", http.StatusInternalServerError)
		return
	}

	// 从info字符串中提取Redis版本信息
	var version string
	for _, line := range strings.Split(info, "\n") {
		if strings.HasPrefix(line, "redis_version:") {
			version = strings.TrimPrefix(line, "redis_version:")
			break
		}
	}

	if version == "" {
		fmt.Println("无法获取Redis版本信息")
		http.Error(w, "无法获取Redis版本信息", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Redis版本: %s", version)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
