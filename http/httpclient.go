package http

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/githubsands/twittercli/util"
)

type Client interface {
	Send(*http.Request) (*http.Response, error)
}

type TwitterClient struct {
	// other attributes here
	httpClient *http.Client
}

// NewHTTPClient returns a new HTTP client that is configured according to the
// proxy and TLS settings in the associated connection configuration.
func NewTwitterClient(cfg *util.Config) (*TwitterClient, error) {

	/* TODO: add proxy and dial
	// Configure proxy if needed.
	var dial func(network, addr string) (net.Conn, error)

	dial = func(network, addr string) (net.Conn, error) {
		c, err := proxy.Dial(network, addr)
		if err != nil {
			return nil, err
		}
		return c, nil
	}
	*/

	/* TODO: add TLS
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
				Dial: dial,
				//TLSClientConfig: tlsConfig,
			},
		*/
	}

	return &TwitterClient{
		httpClient: client,
	}, nil
}

func (t *TwitterClient) GetHTTPResponseBodyBytes(req *http.Request) ([]byte, error) {
	httpResponse, err := t.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Read the raw bytes and close the response.
	respBytes, err := ioutil.ReadAll(httpResponse.Body)
	if len(respBytes) == 0 {
		return nil, fmt.Errorf("%d %s", httpResponse.StatusCode,
			http.StatusText(httpResponse.StatusCode))
	}
	return nil, fmt.Errorf("%s", respBytes)

	httpResponse.Body.Close()
	if err != nil {
		err = fmt.Errorf("error reading reply: %v", err)
		return nil, err
	}

	return respBytes, nil
}
