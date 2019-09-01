package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/githubsands/twittercli/twitterjson"
)

type config struct {
	TLSSkipVerify bool 
}

type Client interface {
	Send(*http.Request) (*http.Response, error)
}

type TwitterClient struct {
	// other attributes here 
	httpClient *http.Client 
}

// NewHTTPClient returns a new HTTP client that is configured according to the
// proxy and TLS settings in the associated connection configuration.
func NewTwitterClient(cfg *config) (*TwitterClient, error) {
	// Configure proxy if needed.
	var dial func(network, addr string) (net.Conn, error)

	dial = func(network, addr string) (net.Conn, error) {
		c, err := proxy.Dial(network, addr)
		if err != nil {
			return nil, err
		}
		return c, nil
	}

	// Configure TLS if needed.
	var tlsConfig *tls.Config
	if !cfg.NoTLS {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: cfg.TLSSkipVerify,
		}
		if !cfg.TLSSkipVerify {
			pem, err := ioutil.ReadFile(cfg.RPCCert)
			if err != nil {
				return nil, err } pool := x509.NewCertPool()
			if ok := pool.AppendCertsFromPEM(pem); !ok {
				return nil, fmt.Errorf("invalid certificate file: %v",
					cfg.RPCCert)
			}
			tlsConfig.RootCAs = pool
		}
	}

	// Create and return the new HTTP client potentially configured with a
	// proxy and TLS.
	client := http.Client{
		Transport: &http.Transport{
			Dial:            dial,
			TLSClientConfig: tlsConfig,
		},
	}

	return &TwitterClient{
		httpClient: client, 
	}, nil 
}

func (t *TwitterClient) SendHTTPRequest(req *http.Request) (*http.Response, error) {
	httpResponse, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if httpResponse.StatusCode < 200 || httpResponse.StatusCode >= 300 {
		if len(respBytes) == 0 {
			return nil, fmt.Errorf("%d %s", httpResponse.StatusCode,
				http.StatusText(httpResponse.StatusCode))
		}
		return nil, fmt.Errorf("%s", respBytes)
	}

	// Read the raw bytes and close the response.
	respBytes, err := ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	if err != nil {
		err = fmt.Errorf("error reading reply: %v", err)
		return nil, err
	}

	return httpResponse, nil
}

