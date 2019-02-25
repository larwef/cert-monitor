package cert

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/larwef/cert-monitor/pkg/config"
	"golang.org/x/oauth2/clientcredentials"
)

// Client is handling the commnunication with the cert api.
type Client struct {
	conf   *config.Config
	client *http.Client
}

// NewClient returns a new cert client.
func NewClient(conf *config.Config) *Client {
	oauthConfig := clientcredentials.Config{
		ClientID:     conf.ClientID,
		ClientSecret: conf.ClientSecret,
		TokenURL:     conf.TokenURL,
		Scopes:       conf.Scopes,
	}

	oauthClient := oauthConfig.Client(context.Background())

	return &Client{
		conf:   conf,
		client: oauthClient,
	}
}

// Search searches for certificates as specified by the Request
func (c *Client) Search(req *Request) (*Response, error) {
	httpReq, err := c.getRequest(req)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(httpReq)
	if err != nil {
		log.Fatalf("Error performing request: %v", err)
		return nil, err
	}

	return res, nil
}

func (c *Client) getRequest(req *Request) (*http.Request, error) {
	reqBuf := new(bytes.Buffer)
	err := json.NewEncoder(reqBuf).Encode(req)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, c.conf.Endpoint, reqBuf)
}

func (c *Client) doRequest(req *http.Request) (*Response, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		b, _ := ioutil.ReadAll(res.Body)
		log.Printf("Response: %s", string(b))
		return nil, fmt.Errorf("Http status: %d", res.StatusCode)
	}

	var certRes Response
	if err := json.NewDecoder(res.Body).Decode(&certRes); err != nil {
		return nil, err
	}

	return &certRes, nil
}
