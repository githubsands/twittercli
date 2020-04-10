package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type config struct {
	TLSSkipVerify bool
}

type Client interface {
	Send(*http.Request) (*http.Response, error)
}

type TwitterClient struct {
	marshalResponse bool
	// other attributes here
	httpClient *http.Client
}

// NewHTTPClient returns a new HTTP client that is configured according to the
// proxy and TLS settings in the associated connection configuration.
func NewTwitterClient(mr bool) (*TwitterClient, error) {

	// Configure proxy if needed.
	/*
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
	*/

	// Create and return the new HTTP client potentially configured with a
	// proxy and TLS.
	client := &http.Client{
		/*
			Transport: &http.Transport{
				Dial:            dial,
				TLSClientConfig: tlsConfig,
			},
		*/
	}

	return &TwitterClient{
		marshalResponse: mr,
		httpClient:      client,
	}, nil
}

func (t *TwitterClient) NewTwitterRequest() (*http.Request, error) {
	req, err := http.NewRequest("POST", "www.twitter.com", nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// GetHTTPResponseBody takes a request, sends it, and returns the response
func (t *TwitterClient) GetHTTPResponse(req *http.Request) (*http.Response, error) {
	res, err := t.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Read the raw bytes and close the response.
	respBytes, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		err = fmt.Errorf("error reading reply: %v", err)
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		if len(respBytes) == 0 {
			return nil, fmt.Errorf("%d %s", res.StatusCode,
				http.StatusText(res.StatusCode))
		}
		return nil, fmt.Errorf("%s", respBytes)
	}

	return res, nil
}

func (t *TwitterClient) GetHTTPResponseBodyBytes(req *http.Request) ([]byte, error) {
	res, err := t.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Read the raw bytes and close the response.
	respBytes, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		err = fmt.Errorf("error reading reply: %v", err)
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		if len(respBytes) == 0 {
			return nil, fmt.Errorf("%d %s", res.StatusCode,
				http.StatusText(res.StatusCode))
		}
		return nil, fmt.Errorf("%s", respBytes)
	}

	return respBytes, nil
}

func (t *TwitterClient) UnmarshalResponse() bool { return t.marshalResponse }
