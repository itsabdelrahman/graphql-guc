package util

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func HttpPostWithFormData(api, resource, username, password, clientVersion, appOS, osVersion string) string {
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

func SendJsonResponse(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func BasicAuthentication(r *http.Request) (string, string) {
	auth := strings.SplitN(r.Header["Authorization"][0], " ", 2)
	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	return pair[0], pair[1]
}

func JsonToStruct(j string, v interface{}) error {
	return json.Unmarshal([]byte(j), v)
}

func XmlToStruct(x string, v interface{}) error {
	return xml.Unmarshal([]byte(x), v)
}

func httpResponseBodyToString(responseBody io.ReadCloser) string {
	responseBodyRead, _ := ioutil.ReadAll(responseBody)
	return string(responseBodyRead)
}
