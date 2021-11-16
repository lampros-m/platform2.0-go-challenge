package test

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

// HTTPClient : Struct to call go service.
type HTTPClient struct {
	serviceURL string
	t          *testing.T
	f          *FileReader
}

// HTTPExpectedResponse struct.
type HTTPExpectedResponse struct {
	ExpectedCode     int
	ExpectedResponse interface{}
	LogResponse      bool
}

// JSONFile alias.
type JSONFile string

// Regex alias.
type Regex string

// NewHTTPClient : HTTP client constructor.
func NewHTTPClient(t *testing.T) *HTTPClient {
	t.Helper()
	return &HTTPClient{
		serviceURL: "http://localhost:8080",
		t:          t,
		f:          NewFileReader(t),
	}
}

// Post executes a post request to `serviceURL` service and expects the requested executeCode and response.
func (h *HTTPClient) PostAndCompare(payload string, path string, headers map[string]string, exp HTTPExpectedResponse) (map[string]interface{}, string) {
	h.t.Helper()

	res := h.PostJSONExpectCode(payload, path, headers, exp.ExpectedCode)
	if exp.LogResponse {
		m := new(interface{})
		json.Unmarshal([]byte(res), &m)
		ls, _ := json.MarshalIndent(m, "", "    ")
		h.t.Logf("Response: \n%s\n", string(ls))
	}

	switch r := exp.ExpectedResponse.(type) {
	case JSONFile:
		LooseJSONEq(h.t, h.f.ReadFile(string(r)), res)
	case string:
		if r != "" {
			LooseJSONEq(h.t, r, res)
		}
	case nil:
	default:
		h.t.Fatalf("Wrong expected response for post: %#v", exp.ExpectedResponse)
	}

	if res == "" {
		return nil, ""
	}

	m := make(map[string]interface{})
	json.Unmarshal([]byte(res), &m)
	return m, res
}

// PostJSONExpectCode executes a post request to `serviceURL` service and returns response as string and expects the given code.
func (h *HTTPClient) PostJSONExpectCode(jsonString string, path string, headers map[string]string, expectCode int) string {
	h.t.Helper()

	return string(h.PostJSONGetBytes(jsonString, path, headers, expectCode))
}

// PostJSONGetBytes executes a post request to `serviceURL` service and returns response as byte slice and expects the given code.
func (h *HTTPClient) PostJSONGetBytes(jsonString string, path string, headers map[string]string, expectCode int) []byte {
	h.t.Helper()

	req, err := http.NewRequest("POST", h.serviceURL+path, strings.NewReader(jsonString))
	if err != nil {
		h.t.Fatalf("Error in http client request creation %s", err.Error())
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		h.t.Fatalf("Error in http client call %s", err.Error())
		return nil
	}
	defer res.Body.Close()

	if res.StatusCode != expectCode {
		body, _ := io.ReadAll(res.Body)
		h.t.Fatalf("Error in http client. Expected code %d got %d \n\tResponse: %s", expectCode, res.StatusCode, string(body))
		return nil
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		h.t.Fatalf("Error in http client response read %s", err.Error())
		return nil
	}

	return body
}
