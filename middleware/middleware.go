package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

type loggingResponseWriter struct {
	w          http.ResponseWriter
	statusCode int
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remoteAddrInfo := strings.Split(r.RemoteAddr, ":")
		fmt.Printf("remote addr: %s, remote port: %s, ", remoteAddrInfo[0], remoteAddrInfo[1])
		fmt.Printf("resp code:\n")
		next.ServeHTTP(w, r)
	})
}
