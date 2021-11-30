package handler

import (
	"net"
	"net/http"
	"os"
	"strings"

	"go.uber.org/zap"
)

func New(lg *zap.Logger) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthz)

	return logMiddleware(headerMiddleware(mux), lg)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func logMiddleware(next http.Handler, lg *zap.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := wrapperResponseWriter(w)

		next.ServeHTTP(rw, r)

		lg.Sugar().Infof("请求IP：%v，HTTP返回码：%v\n", ipAddress(r), rw.statusCode)
	})
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for key := range r.Header {
			w.Header().Add(key, r.Header.Get(key))
		}

		version := os.Getenv("VERSION")
		w.Header().Add("Version", version)

		next.ServeHTTP(w, r)
	})
}

// ipAddress 获取客户端 IP
func ipAddress(r *http.Request) string {
	if ip := strings.TrimSpace(strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0]); ip != "" {
		return ip
	}

	if ip := strings.TrimSpace(r.Header.Get("X-Real-Ip")); ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
