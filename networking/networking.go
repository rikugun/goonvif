package networking

import (
	"bytes"
	"net/http"
	"time"
)

func SendSoap(endpoint, message string) (*http.Response, error) {
	//httpClient := new(http.Client)
	//
	//resp, err := httpClient.Post(endpoint, "application/soap+xml; charset=utf-8", bytes.NewBufferString(message))
	//if err != nil {
	//	return resp, err
	//}
	//
	//return resp,nil
	return SendSoapWithTimeout(endpoint, message, time.Second*3)
}

/**
  Add timeout
*/
func SendSoapWithTimeout(endpoint string, message string, timeout time.Duration) (*http.Response, error) {
	httpClient := &http.Client{
		Timeout: timeout,
	}
	return httpClient.Post(endpoint, "application/soap+xml; charset=utf-8", bytes.NewBufferString(message))
}
