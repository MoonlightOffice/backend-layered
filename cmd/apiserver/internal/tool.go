package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type H map[string]interface{}

func writeResponse(w http.ResponseWriter, status int, body any) {
	w.WriteHeader(status)

	b, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, string(b))
}

func bindReqData(r *http.Request, v any) bool {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return false
	}
	defer r.Body.Close()

	return json.Unmarshal(data, &v) == nil
}

func ClientIP(r *http.Request) string {
	// Get the forwarded-for IP addresses
	forwardedFor := r.Header.Get("X-Forwarded-For")

	// If there are multiple IP addresses, take the last one
	// (the one closest to the client)
	ips := strings.Split(forwardedFor, ",")
	if len(ips) > 0 {
		clientIP := strings.TrimSpace(ips[len(ips)-1])
		if len(clientIP) != 0 {
			return clientIP
		}
	}

	// If there are no forwarded IPs, use the remote address
	splitted := strings.Split(r.RemoteAddr, ":")
	n := len(splitted)

	// IPv4
	if n == 2 {
		return splitted[0]
	}

	// IPv6
	if n > 2 {
		return strings.Join(splitted[:n-1], ":")
	}

	// Others
	return r.RemoteAddr
}
