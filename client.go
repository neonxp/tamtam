package tamtam

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type client struct {
	key        string
	version    string
	url        *url.URL
	httpClient *http.Client
}

func newClient(key string, version string, url *url.URL, httpClient *http.Client) *client {
	return &client{key: key, version: version, url: url, httpClient: httpClient}
}

func (cl *client) request(method, path string, query url.Values, body interface{}) (io.ReadCloser, error) {
	j, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return cl.requestReader(method, path, query, bytes.NewReader(j))
}

func (cl *client) requestReader(method, path string, query url.Values, body io.Reader) (io.ReadCloser, error) {
	u := *cl.url
	u.Path = path
	query.Set("access_token", cl.key)
	query.Set("v", cl.version)
	u.RawQuery = query.Encode()
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	resp, err := cl.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		errObj := new(Error)
		err = json.NewDecoder(resp.Body).Decode(errObj)
		if err != nil {
			return nil, err
		}
		return nil, errObj
	}
	return resp.Body, err
}
