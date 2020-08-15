package sms

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/paradox-3arthling/africastalking"
)

// check how `json.Marshal/1` marshals arrays
type Request_data struct {
	Prod                 bool
	Api_key              string
	Username             string   //`username`
	To                   []string //`to`
	Message              string   //`message`
	From                 string   //`from`
	BulkSMSMode          int      //`bulkSMSMode`
	Enqueue              int      //`enqueue`
	Keyword              string   //`keyword`
	LinkId               string   //`linkId`
	RetryDurationInHours int      //`retryDurationInHours`
}

func (req_data *Request_data) encodeValues() []byte {
	data := url.Values{}
	data.Set("username", req_data.Username)
	data.Set("message", req_data.Message)
	data.Set("to", strings.Join(req_data.To, ","))

	return []byte(data.Encode())
}

// Need to confirm that the nos. are valid
func (req_data *Request_data) ConfirmFields() error {
	if req_data.Username == "" {
		return fmt.Errorf("`ConfirmFields` got `req_data.Username` is blank!")
	}
	if len(req_data.To) == 0 {
		return fmt.Errorf("`ConfirmFields` got `req_data.To` has no numbers to send to!")
	}
	if req_data.Message == "" {
		return fmt.Errorf("`ConfirmFields` got `req_data.Message` is blank!")
	}
	return nil
}

// Return data at final func
func (req_data *Request_data) SendSMS() error {
	prod := req_data.Prod
	if prod == false {
		req_data.Username = "sandbox"
	}
	err := req_data.ConfirmFields()
	if err != nil {
		return fmt.Errorf("'req_data.ConfirmFields/0' got the error: %q", err)
	}
	// data, err := json.Marshal(req_data)
	// if err != nil {
	// 	return fmt.Errorf("'json.Marshal/1' got the error: %q", err)
	// }
	data := req_data.encodeValues()
	url := africastalking.SetUrl(prod, africastalking.SMS_URL)
	req, err := africastalking.EncodedRequest(url, req_data.Api_key, data)
	if err != nil {
		return fmt.Errorf("'africastalking.JsonRequest/3' got the error: %q", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("'client.Do/1' got the error: %q", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("'ioutil.ReadAll/1' got the error: %q", err)
	}
	fmt.Printf("Body received: %q\nURL: %q\ndata: %q", string(body), url, string(data))
	return nil
}
