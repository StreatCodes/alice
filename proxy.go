package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func matchingProxy(proxies []Proxy, headerAddress string) *Proxy {
	for _, proxy := range proxies {
		if strings.HasPrefix(headerAddress, proxy.Address) {
			return &proxy
		}
	}

	return nil
}

func proxyHandler(config *Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerAddress := r.Header.Get("X-Alice-Address")
		log.Printf("Proxying request: %s %s\n", r.Method, headerAddress)

		req, err := http.NewRequest(r.Method, headerAddress, r.Body)
		if err != nil {
			http.Error(w, "Failed to proxy request "+err.Error(), http.StatusBadGateway)
			return
		}

		// Include the configured headers if we have a matching config for the address
		req.Header = r.Header.Clone()
		match := matchingProxy(config.Proxies, headerAddress)
		if match != nil {
			log.Println("Matching URL, appending headers:")
			for k, v := range match.Headers {
				log.Printf("\t%s\n", k)
				req.Header.Set(k, v)
			}
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, "Failed to proxy response "+err.Error(), http.StatusBadGateway)
			return
		}

		// Copy the response headers into our response
		for k, values := range res.Header {
			for _, v := range values {
				w.Header().Add(k, v)
			}
		}

		w.WriteHeader(res.StatusCode)
		_, err = io.Copy(w, res.Body)
		if err != nil {
			http.Error(w, "Failed to proxy response body "+err.Error(), http.StatusBadGateway)
			return
		}

	})
}
