package africastalking

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

func SetUrl(prod bool, url string) string {
	if prod {
		return strings.Replace(url, "sandbox.", "", 1)
	} else {
		return url
	}
}

// BasicPostClient/1 - Created client will be dealing with
// json only.
func JsonRequest(url, api_key string, data []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return req, fmt.Errorf("'http.NewRequest/3' got the error: %q", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("apiKey", api_key)
	return req, nil
}