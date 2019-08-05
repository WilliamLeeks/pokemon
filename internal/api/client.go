package api

import (
	"net"
	"net/http"
	"time"
)

// getDefaultClient returns an HTTP client.
func getDefaultClient() *http.Client {
	var netTransport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          5,
		IdleConnTimeout:       10 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   5,
	}

	httpClient := &http.Client{
		Timeout:   time.Second * time.Duration(10),
		Transport: netTransport,
	}

	return httpClient
}
