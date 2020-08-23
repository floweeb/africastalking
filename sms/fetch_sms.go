package sms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/paradox-3arthling/africastalking"
)

func CheckInbox() (map[string]interface{}, error) {
	inbox := make(map[string]interface{})
	// !!change appropiately when the parameters come in
	// https://api.sandbox.africastalking.com/version1/messaging?username=MyAppUsername&lastReceivedId=0
	// params := fmt.Sprintf("?username=%s&lastReceivedId=%v", username, last_received)
	params := fmt.Sprintf("?username=floyd&lastReceivedId=0")
	// !!change the url when the we whether we are in production or nah
	// url := africastalking.SetUrl(prod, africastalking.INBOX_URL)
	url := africastalking.INBOX_URL + params
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return inbox, fmt.Errorf("'http.NewRequest/3' got the error: %q", err)
	}
	req.Header.Add("Accept", "application/json")
	// !!set the API key
	// req.Header.Add("apiKey", api_key)
	req.Header.Add("apiKey", "")

	resp, err := client.Do(req)
	if err != nil {
		return inbox, fmt.Errorf("'client.Do/1' got the error: %q", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return inbox, fmt.Errorf("'ioutil.ReadAll/1' got the error: %q", err)
	}
	err = json.Unmarshal(body, &inbox)
	if err != nil {
		return inbox, fmt.Errorf("'ioutil.ReadAll/1' got the error: %q", err)
	}
	return inbox, nil
}