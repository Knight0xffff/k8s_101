package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

const versionKey = "VERSION"

func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		for key, values := range r.Header {
			for _, v := range values {
				w.Header().Add(key, v)
			}
		}

		clientIp := getClientIp(r)
		version := os.Getenv(versionKey)
		w.Header().Set(versionKey, version)
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "ok")
		fmt.Printf("%s client ip: %s path: %s status code: %d", time.Now().Format(time.RFC822), clientIp, "/healthz", http.StatusOK)
	})

	err := http.ListenAndServe(":3888", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getClientIp(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err != nil {
		fmt.Printf("get client ip err: %s", err.Error())
	}

	if ip != "" {
		return ip
	}

	return "unknown ip"
}
