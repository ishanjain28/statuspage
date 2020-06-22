package statuspage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	log "github.com/sirupsen/logrus"
)

// DefaultProdHost points to the default public statuspage.io V1 Endpoint
const DefaultProdHost string = "https://api.statuspage.io"

type Client struct {
	url    string
	apiKey string
}

func WithAPIKey(key string) Client {
	return Client{
		url:    DefaultProdHost,
		apiKey: key,
	}
}

func WithAPIKeyAndHost(ustr, key string) Client {
	return Client{
		url:    ustr,
		apiKey: key,
	}
}

type authType int

const (
	authOauthHeader  authType = 0
	apiKeyQueryParam authType = 1
)

func (c Client) applyAuthParam(req *http.Request, at authType) {
	switch at {
	case apiKeyQueryParam:
		query := req.URL.Query()
		query.Set("api_key", c.apiKey)
		req.URL.RawQuery = query.Encode()
	case authOauthHeader:
		req.Header.Set("Authorization", fmt.Sprintf("Oauth %s", c.apiKey))
	}
}

func (c Client) makeRequest(req *http.Request, at authType, response interface{}, goodRespCode int) error {
	client := &http.Client{}

	// Add Auth parameter
	c.applyAuthParam(req, at)

	// Set Content-Type/Accept
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case goodRespCode:
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			b, _ := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Debugf("error in logging response: %v", err)
			}
			log.Debugf("%s: %v", reflect.TypeOf(response), string(b))
			return fmt.Errorf("error in unmarshalling response: %v", err)
		}
		return nil

	default:
		out := new(FailedResponse)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error in reading response body: %v", err)
		}

		out.StatusCode = resp.StatusCode
		out.Message = string(body)
		return out
	}
}
