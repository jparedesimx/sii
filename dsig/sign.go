package dsig

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Sign xml files using xmlsec1 library
func Sign(certBase64 string, password string, xmlData string) ([]byte, error) {
	// Decode certificate base64
	pfxData, err := base64.StdEncoding.DecodeString(certBase64)
	if err != nil {
		log.Println("Error decoding certificate. ", err.Error())
		return nil, err
	}
	tmpFolder, err := ioutil.TempDir("", "tmp")
	if err != nil {
		log.Println("TempDir error", err.Error())
	}
	pfxFile := filepath.Join(tmpFolder, "cert.pfx")
	log.Println(pfxFile)
	xmlFile := filepath.Join(tmpFolder, "file.xml")
	log.Println(xmlFile)
	signedFile := filepath.Join(tmpFolder, "file_signed.xml")
	// Save certificate to disk
	err = ioutil.WriteFile(pfxFile, pfxData, 0755)
	if err != nil {
		log.Println("cert.pfx", err.Error())
	}
	// Save xml to disk
	err = ioutil.WriteFile(xmlFile, []byte(xmlData), 0755)
	if err != nil {
		log.Println("file.xml", err.Error())
	}
	// Generate signed file
	cmd := fmt.Sprintf("xmlsec1 --sign --output %s --pkcs12 %s --pwd %s %s", signedFile, pfxFile, password, xmlFile)
	log.Println(cmd)
	_, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Println("Error in xmlsec1 command. ", err.Error())
		return nil, err
	}
	fileSigned, err := ioutil.ReadFile(signedFile)
	if err != nil {
		log.Println("Error reading file_signed.xml. ", err.Error())
		return nil, err
	}
	// Remove temporary files
	os.Remove(pfxFile)
	os.Remove(xmlFile)
	os.Remove(signedFile)
	return fileSigned, nil
}
