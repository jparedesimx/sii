package login

import (
	"encoding/xml"
	"log"
	"sii/sii/config"
	"sii/sii/dsig"
	"sii/sii/soap"
	"strings"

	"github.com/antchfx/xmlquery"
)

type seed struct {
	XMLName xml.Name
	Body    struct {
		XMLName         xml.Name
		GetSeedResponse struct {
			XMLName       xml.Name
			GetSeedReturn string `xml:"getSeedReturn"`
		} `xml:"getSeedResponse"`
	}
}
type token struct {
	XMLName xml.Name
	Body    struct {
		XMLName          xml.Name
		GetTokenResponse struct {
			XMLName        xml.Name
			GetTokenReturn string `xml:"getTokenReturn"`
		} `xml:"getTokenResponse"`
	}
}

var response []byte
var err error

// AuthWebService implements SII authentication using soap webservices
func AuthWebService(certBase64 string, password string) (string, error) {
	body := []byte(strings.TrimSpace(config.SeedTemplate))
	retries := 100
	for retries > 0 {
		response, err = soap.Request(config.SeedWsdl, body)
		if err != nil {
			retries--
			if retries == 0 {
				return "", err
			}
		} else {
			break
		}
	}
	log.Println("Tries: ", retries)
	// Parse response to xml struct
	var seed seed
	err = xml.Unmarshal([]byte(string(response)), &seed)
	if err != nil {
		log.Println("Error unmarshalling xml. ", err.Error())
		return "", err
	}
	responseNode, err := xmlquery.Parse(strings.NewReader(seed.Body.GetSeedResponse.GetSeedReturn))
	if err != nil {
		log.Println("Error reading seed. ", err.Error())
		return "", err
	}
	seedNode := xmlquery.FindOne(responseNode, "//SEMILLA")
	log.Println("SEMILLA:", seedNode.InnerText())
	pszXML := strings.Replace(config.PszXML, "@seed", seedNode.InnerText(), 1)
	// Sign pszXML and return the generated file like a byte array
	pszSigned, err := dsig.Sign(certBase64, password, pszXML)
	if err != nil {
		return "", err
	}
	pszXML = strings.Replace(config.TokenTemplate, "@pszXML", string(pszSigned), 1)
	body = []byte(strings.TrimSpace(pszXML))
	retries = 100
	for retries > 0 {
		response, err = soap.Request(config.TokenWsdl, body)
		if err != nil {
			retries--
			if retries == 0 {
				return "", err
			}
		} else {
			break
		}
	}
	log.Println("Tries: ", retries)
	// Parse response to xml struct
	var token token
	err = xml.Unmarshal([]byte(string(response)), &token)
	if err != nil {
		log.Println("Error unmarshalling xml. ", err.Error())
		return "", err
	}
	responseNode, err = xmlquery.Parse(strings.NewReader(token.Body.GetTokenResponse.GetTokenReturn))
	if err != nil {
		log.Println("Error reading token. ", err.Error())
		return "", err
	}
	tokenNode := xmlquery.FindOne(responseNode, "//TOKEN")
	log.Println("TOKEN:", tokenNode.InnerText())
	return tokenNode.InnerText(), nil
}
