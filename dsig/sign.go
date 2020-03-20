package dsig

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// Sign xml files using xmlsec1 library
func Sign(certBase64 string, password string, xmlData string) ([]byte, error) {
	// Decode certificate base64
	pfxData, err := base64.StdEncoding.DecodeString(certBase64)
	if err != nil {
		log.Println("Error decoding certificate. ", err.Error())
		return nil, err
	}
	// Save certificate to disk
	err = ioutil.WriteFile("cert.pfx", pfxData, 0755)
	log.Println("cert.pfx", err.Error())
	// Save xml to disk
	err = ioutil.WriteFile("file.xml", []byte(xmlData), 0755)
	log.Println("file.xml", err.Error())
	// Generate signed file
	cmd := fmt.Sprintf("xmlsec1 --sign --output file_signed.xml --pkcs12 cert.pfx --pwd %s file.xml", password)
	log.Println(cmd)
	_, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Println("Error in xmlsec1 command. ", err.Error())
		return nil, err
	}
	fileSigned, err := ioutil.ReadFile("file_signed.xml")
	if err != nil {
		log.Println("Error reading file_signed.xml. ", err.Error())
		return nil, err
	}
	// Remove temporary files
	os.Remove("cert.pfx")
	os.Remove("file.xml")
	os.Remove("file_signed.xml")
	return fileSigned, nil
}
