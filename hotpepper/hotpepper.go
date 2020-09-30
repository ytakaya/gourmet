package hotpepper

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const baseURL = "https://webservice.recruit.co.jp/hotpepper/"

type APIClient struct {
	key        string
	httpClient *http.Client
}

func New(key string) *APIClient {
	apiClient := &APIClient{key, &http.Client{}}
	return apiClient
}

func (api *APIClient) doRequest(method, urlPath string, query map[string]string, data []byte) (body []byte, err error) {
	baseURL, err := url.Parse(baseURL)
	if err != nil {
		return
	}
	apiURL, err := url.Parse(urlPath)
	if err != nil {
		return
	}
	endpoint := baseURL.ResolveReference(apiURL).String()
	log.Printf("action=doRequest endpoint=%s", endpoint)
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	q := req.URL.Query()
	q.Add("key", api.key)
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (api *APIClient) GetGourmet() ([]Gourmet, error) {
	url := "gourmet/v1"
	resp, err := api.doRequest("GET", url, map[string]string{"lat": "34.67", "lng": "135.52"}, nil)
	log.Printf("url=%s resp=%s", url, string(resp))
	if err != nil {
		log.Printf("action=GetGourmet err=%s", err.Error())
		return nil, err
	}
	var gourmet []Gourmet
	err = xml.Unmarshal(resp, &gourmet)
	if err != nil {
		log.Printf("action=GetGourmet err=%s", err.Error())
		return nil, err
	}
	return gourmet, nil
}
