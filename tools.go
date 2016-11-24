package apipe

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
//    "log"
	"net/http"
//    "net/http/httputil"
	"strconv"
)

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func doRequest(reqUrl string, reqType string, accessToken string) (result *http.Response, err error) {
	req, err := http.NewRequest(reqType, reqUrl, nil)
    req.Header.Add("Authorization", "Bearer " + accessToken)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	return client.Do(req)
}

func getAPIErrors(resp *http.Response) (err error) {
	if resp == nil {
		return errors.New("Service error response is unidentified")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Service error response is unidentified")
	}
	defer resp.Body.Close()
	var f interface{}
	json.Unmarshal(body, &f)
	errorsResp, ok := f.(map[string]interface{})["errors"].([]interface{})
	if ok == false {
		return errors.New("Service error response is unidentified")
	}
	errMessage := "Response read with " + strconv.Itoa(len(errorsResp)) + " errors:\n"
	for _, val := range errorsResp {
		errMessage += val.(string) + "_\n"
	}
	return errors.New(errMessage)
}

func getWithToken(reqUrl string, accessToken string) (data []byte, err error) {
	resp, err := doRequest(reqUrl, "GET", accessToken)
	if err != nil {
		return nil, err
	}
    // dump, err := httputil.DumpResponse(resp, true)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Printf("%q", dump)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Service response is unidentified")
	}
	defer resp.Body.Close()
	return body, nil
}


func get(reqUrl string) (data []byte, err error) {
	accessToken, err := Authenticate()
	if err != nil {
		fmt.Println("Unable to get token")
		return nil, err
	}
	resp, err := doRequest(reqUrl, "GET", accessToken)
	if err != nil {
		return nil, err
	}
    // dump, err := httputil.DumpResponse(resp, true)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Printf("%q", dump)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Service response is unidentified")
	}
	defer resp.Body.Close()
	return body, nil
}

func post(reqUrl string, accessToken string) (data []byte, err error) {
	fmt.Println(reqUrl)
	resp, err := doRequest(reqUrl, "POST", accessToken)
	if err != nil {
		return nil, err
	}
    // dump, err := httputil.DumpResponse(resp, true)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Printf("%q", dump)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Service response is unidentified")
	}
	defer resp.Body.Close()
	return body, nil
}
