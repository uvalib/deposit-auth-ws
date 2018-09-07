package client

import (
	"github.com/uvalib/deposit-auth-ws/depositauthws/api"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var debugHTTP = false
var serviceTimeout = 5

//
// HealthCheck -- calls the service health check method
//
func HealthCheck(endpoint string) int {

	url := fmt.Sprintf("%s/healthcheck", endpoint)
	//fmt.Printf( "%s\n", url )

	resp, _, errs := gorequest.New().
		SetDebug(debugHTTP).
		Get(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	return resp.StatusCode
}

//
// VersionCheck -- calls the service version check method
//
func VersionCheck(endpoint string) (int, string) {

	url := fmt.Sprintf("%s/version", endpoint)
	//fmt.Printf( "%s\n", url )

	resp, body, errs := gorequest.New().
		SetDebug(false).
		Get(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError, ""
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	r := api.VersionResponse{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return http.StatusInternalServerError, ""
	}

	return resp.StatusCode, r.Version
}

//
// MetricsCheck -- calls the service metrics method
//
func MetricsCheck(endpoint string) (int, string) {

	url := fmt.Sprintf("%s/metrics", endpoint)
	//fmt.Printf( "%s\n", url )

	resp, body, errs := gorequest.New().
		SetDebug(false).
		Get(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError, ""
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	return resp.StatusCode, body
}

//
// GetDepositAuthorization -- calls the service get authorization method
//
func GetDepositAuthorization(endpoint string, id string, token string) (int, []*api.Authorization) {

	url := fmt.Sprintf("%s/%s?auth=%s", endpoint, id, token)
	//fmt.Printf( "%s\n", url )

	resp, body, errs := gorequest.New().
		SetDebug(debugHTTP).
		Get(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError, nil
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	r := api.StandardResponse{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return resp.StatusCode, r.Details
}

//
// GetInboundDepositAuthorization -- gets new authorizations
//
func GetInboundDepositAuthorization(endpoint string, after string, token string) (int, []*api.Authorization) {

	url := fmt.Sprintf("%s/inbound?auth=%s&after=%s", endpoint, token, after)
	//fmt.Printf( "%s\n", url )
	return (getDepositAuthorization(url))
}

//
// SearchDepositAuthorizationByCid -- calls the service search by cid method
//
func SearchDepositAuthorizationByCid(endpoint string, cid string, token string) (int, []*api.Authorization) {

	url := fmt.Sprintf("%s?auth=%s&cid=%s", endpoint, token, cid)
	//fmt.Printf( "%s\n", url )
	return (getDepositAuthorization(url))
}

//
// SearchDepositAuthorizationByCreated -- calls the service search by create date method
//
func SearchDepositAuthorizationByCreated(endpoint string, created string, token string) (int, []*api.Authorization) {

	url := fmt.Sprintf("%s?auth=%s&created=%s", endpoint, token, created)
	//fmt.Printf( "%s\n", url )
	return (getDepositAuthorization(url))
}

//
// SearchDepositAuthorizationByExported -- calls the service search by exported date method
//
func SearchDepositAuthorizationByExported(endpoint string, exported string, token string) (int, []*api.Authorization) {

	url := fmt.Sprintf("%s?auth=%s&exported=%s", endpoint, token, exported)
	//fmt.Printf( "%s\n", url )
	return (getDepositAuthorization(url))
}

//
// ImportDepositAuthorization -- calls the service import method
//
func ImportDepositAuthorization(endpoint string, token string) (int, int, int, int, int) {

	url := fmt.Sprintf("%s/import?auth=%s", endpoint, token)
	//fmt.Printf( "%s\n", url )

	resp, body, errs := gorequest.New().
		SetDebug(debugHTTP).
		Post(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError, 0, 0, 0, 0
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	r := api.ImportResponse{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return http.StatusInternalServerError, 0, 0, 0, 0
	}

	return resp.StatusCode, r.NewCount, r.UpdatedCount, r.DuplicateCount, r.ErrorCount
}

//
// ExportDepositAuthorization -- calls the service export method
//
func ExportDepositAuthorization(endpoint string, token string) (int, int, int) {

	url := fmt.Sprintf("%s/export?auth=%s", endpoint, token)
	//fmt.Printf( "%s\n", url )

	resp, body, errs := gorequest.New().
		SetDebug(debugHTTP).
		Post(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError, 0, 0
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	r := api.ExportResponse{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return http.StatusInternalServerError, 0, 0
	}

	return resp.StatusCode, r.ExportCount, r.ErrorCount
}

//
// FulfillDepositAuthorization -- calls the service fulfil method
//
func FulfillDepositAuthorization(endpoint string, id string, depositID string, token string) int {

	url := fmt.Sprintf("%s/%s?deposit=%s&auth=%s", endpoint, id, depositID, token)
	//fmt.Printf( "%s\n", url )

	resp, _, errs := gorequest.New().
		SetDebug(debugHTTP).
		Put(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	return resp.StatusCode
}

func getDepositAuthorization(url string) (int, []*api.Authorization) {

	resp, body, errs := gorequest.New().
		SetDebug(debugHTTP).
		Get(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError, nil
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	r := api.StandardResponse{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return resp.StatusCode, r.Details
}

//
// end of file
//
