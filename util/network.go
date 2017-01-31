package util

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// HTTPPostWithFormData sends an HTTP POST request with form data
func HTTPPostWithFormData(api, resource, username, password, clientVersion, appOS, osVersion string) string {
	data := url.Values{}
	data.Set("username", username)
	data.Add("password", password)
	data.Add("clientVersion", clientVersion)

	if appOS != "" && osVersion != "" {
		data.Add("app_os", appOS)
		data.Add("os_version", osVersion)
	}

	uri, _ := url.ParseRequestURI(api)
	uri.Path = resource
	uriString := fmt.Sprintf("%v", uri)

	client := &http.Client{}
	request, _ := http.NewRequest("POST", uriString, bytes.NewBufferString(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	response, _ := client.Do(request)
	defer response.Body.Close()

	responseBodyString := httpResponseBodyToString(response.Body)
	return responseBodyString
}

// SendJSONResponse sends an HTTP response with JSON content
func SendJSONResponse(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

// BasicAuthentication extracts credentials from an HTTP Authorization header
func BasicAuthentication(r *http.Request) (string, string, error) {
	auth := r.Header["Authorization"]

	if len(auth) == 0 {
		return "", "", errors.New("Unauthorized")
	}

	authSplit := strings.SplitN(r.Header["Authorization"][0], " ", 2)
	payload, _ := base64.StdEncoding.DecodeString(authSplit[1])
	pair := strings.SplitN(string(payload), ":", 2)

	return strings.TrimSpace(pair[0]), strings.TrimSpace(pair[1]), nil
}

// JSONToStruct unmarshals JSON to interface
func JSONToStruct(j string, v interface{}) error {
	return json.Unmarshal([]byte(j), v)
}

// XMLToStruct unmarshals XML to interface
func XMLToStruct(x string, v interface{}) error {
	return xml.Unmarshal([]byte(x), v)
}

func httpResponseBodyToString(responseBody io.ReadCloser) string {
	responseBodyRead, _ := ioutil.ReadAll(responseBody)
	return string(responseBodyRead)
}
