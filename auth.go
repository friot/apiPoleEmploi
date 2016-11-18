package apipe

import (
    "bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var Conf = map[string]string{}

func SetCredentialsFromEnv() {
	Conf = map[string]string{
        "PE_API_URI"                : "PE_API_URI",
        "PE_API_GET_URI"            : "PE_API_GET_URI",
    	"PE_AUTH_URL"               : "PE_AUTH_URL",
        "PE_GET_URL"                : "PE_GET_URL",
        "PE_ID_PUBLIC_SERVICE"      : "PE_ID_PUBLIC_SERVICE",
        "PE_ID_PRIVATE_SERVICE"     : "PE_ID_PRIVATE_SERVICE",
        "PE_REALM"                  : "PE_REALM",
        "PE_GRANT_TYPE"             : "PE_GRANT_TYPE",
        "PE_LIST_COLLECTION"        : "PE_LIST_COLLECTION",
        "PE_INFO_HOST_COLLECTION"   : "PE_INFO_HOST_COLLECTION",
        "PE_INFO_COLLECTION"        : "PE_INFO_COLLECTION",
    	"PE_COLLECTION_DATA"        : "PE_COLLECTION_DATA",
	}
    for _, value := range Conf {
        Conf[value] = os.Getenv(value)
    }
}

func getAuthRequest() (result *http.Request, err error) {
	authUrl := Conf["PE_API_URI"] + Conf["PE_AUTH_URL"]
	data := url.Values{}
	data.Set("client_id", Conf["PE_ID_PUBLIC_SERVICE"])
    data.Add("client_secret", Conf["PE_ID_PRIVATE_SERVICE"])
    data.Add("realm", Conf["PE_REALM"])
	data.Add("grant_type", Conf["PE_GRANT_TYPE"])

	if result, err = http.NewRequest("POST", authUrl, bytes.NewBufferString(data.Encode())); err == nil {
		result.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	return
}

func authService() (result *http.Response, err error) {

	req, err := getAuthRequest()
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	return client.Do(req)
}

func Authenticate() (result string, err error) {

	resp, err := authService()
    if err != nil {
        return "", err
    }

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var receiver JSONPoleEmploiToken
	err = json.Unmarshal(body, &receiver)
	if err != nil {
		return "", errors.New("Service response is unidentified")
	}
	return receiver.Access_token, nil
}
