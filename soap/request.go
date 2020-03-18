package soap

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
)

var client *http.Client

// Request implements a soap request to a given url
func Request(url string, body []byte) ([]byte, error) {
	// HTTP client
	if client == nil {
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
	}
	// Prepare the request
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		log.Println("Error creating request object. ", err.Error())
		return nil, err
	}
	// Set the Content-type header, as well as the other required headers
	req.Header.Set("Content-type", "text/xml;charset=UTF-8")
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("SOAPAction", "")
	// Dispatch the request
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error dispatching the request. ", err.Error())
		return nil, err
	}
	defer res.Body.Close()
	log.Println("-> Retrieving and parsing the response")
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response. ", err.Error())
		return nil, err
	}
	return data, nil
}
