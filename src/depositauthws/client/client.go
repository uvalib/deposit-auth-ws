package client

import (
	"depositauthws/api"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var debugHttp = false
var serviceTimeout = 5

func HealthCheck(endpoint string) int {

	url := fmt.Sprintf("%s/healthcheck", endpoint)
	//fmt.Printf( "%s\n", url )

	resp, _, errs := gorequest.New().
		SetDebug(debugHttp).
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

func RuntimeCheck(endpoint string) (int, *api.RuntimeResponse) {

	url := fmt.Sprintf("%s/runtime", endpoint)
	//fmt.Printf( "%s\n", url )

	resp, body, errs := gorequest.New().
		SetDebug(false).
		Get(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError, nil
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	r := api.RuntimeResponse{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return resp.StatusCode, &r
}

func GetDepositAuthorization(endpoint string, id string, token string) (int, []*api.Authorization) {

	url := fmt.Sprintf("%s/%s?auth=%s", endpoint, id, token)
	//fmt.Printf( "%s\n", url )

	resp, body, errs := gorequest.New().
		SetDebug(debugHttp).
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

func SearchDepositAuthorizationById(endpoint string, id string, token string) (int, []*api.Authorization) {

	url := fmt.Sprintf("%s?auth=%s&later=%s", endpoint, token, id)
	//fmt.Printf( "%s\n", url )
	return (searchDepositAuthorization(url))
}

func SearchDepositAuthorizationByCid(endpoint string, cid string, token string) (int, []*api.Authorization) {

	url := fmt.Sprintf("%s?auth=%s&cid=%s", endpoint, token, cid)
	//fmt.Printf( "%s\n", url )
	return (searchDepositAuthorization(url))
}

func SearchDepositAuthorizationByCreated(endpoint string, created string, token string) (int, []*api.Authorization) {

	url := fmt.Sprintf("%s?auth=%s&created=%s", endpoint, token, created)
	//fmt.Printf( "%s\n", url )
	return (searchDepositAuthorization(url))
}

func SearchDepositAuthorizationByExported(endpoint string, exported string, token string) (int, []*api.Authorization) {

	url := fmt.Sprintf("%s?auth=%s&exported=%s", endpoint, token, exported)
	//fmt.Printf( "%s\n", url )
	return (searchDepositAuthorization(url))
}

func searchDepositAuthorization(url string) (int, []*api.Authorization) {

	resp, body, errs := gorequest.New().
		SetDebug(debugHttp).
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

func ImportDepositAuthorization(endpoint string, token string) (int, int) {

	url := fmt.Sprintf("%s/import?auth=%s", endpoint, token)
	//fmt.Printf( "%s\n", url )

	resp, body, errs := gorequest.New().
		SetDebug(debugHttp).
		Post(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError, 0
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	r := api.ImportExportResponse{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return http.StatusInternalServerError, 0
	}

	return resp.StatusCode, r.Count
}

func ExportDepositAuthorization(endpoint string, token string) (int, int) {

	url := fmt.Sprintf("%s/export?auth=%s", endpoint, token)
	//fmt.Printf( "%s\n", url )

	resp, body, errs := gorequest.New().
		SetDebug(debugHttp).
		Post(url).
		Timeout(time.Duration(serviceTimeout) * time.Second).
		End()

	if errs != nil {
		return http.StatusInternalServerError, 0
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	r := api.ImportExportResponse{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return http.StatusInternalServerError, 0
	}

	return resp.StatusCode, r.Count
}

func FulfillDepositAuthorization(endpoint string, id string, depositId string, token string) int {

	url := fmt.Sprintf("%s/%s?deposit=%s&auth=%s", endpoint, id, depositId, token)
	//fmt.Printf( "%s\n", url )

	resp, _, errs := gorequest.New().
		SetDebug(debugHttp).
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
