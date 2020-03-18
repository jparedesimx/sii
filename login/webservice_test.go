package login

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jparedesimx/sii/model"
)

func TestAuthWebService(t *testing.T) {
	response, err := http.Get("http://localhost:9494/v1/companies/4.json?expand=[certificate]")
	if err != nil {
		t.Fatalf("Test error. %s", err.Error())
	}
	companyData, _ := ioutil.ReadAll(response.Body)
	var company model.Company
	json.Unmarshal([]byte(string(companyData)), &company)
	actualResult, err := AuthWebService(company.Certificate.Base64, company.Certificate.Pass)
	if err != nil {
		t.Fatalf("Test error. %s", err.Error())
	}
	var expectedResult = "V21MU9PTCD7A9"
	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}
