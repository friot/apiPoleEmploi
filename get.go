package apipe

import (
	"encoding/json"
	"errors"
//    "bytes"
//    "os"
	"fmt"
	"net/url"
	"strconv"
)

func GetServicesList() (result JSONListOfServices, err error) {

    dataUrl := Conf["PE_API_GET_URI"] + Conf["PE_GET_URL"] + Conf["PE_LIST_COLLECTION"] + "id=digidata"
	data, err := get(dataUrl)
	if err != nil {
		return result, errors.New("Data doesn't exist")
	}
    // var out bytes.Buffer
	// json.Indent(&out, data, "=", "\t")
	// out.WriteTo(os.Stdout)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, errors.New("Service response is unidentified")
	}
	return
}

func GetListOfCollectionsFromServicesId(id_services string) (result JSONListOfCollectionByServices, err error) {

    dataUrl := Conf["PE_API_GET_URI"] + Conf["PE_GET_URL"] + Conf["PE_INFO_COLLECTION"] + "id=" + id_services
	data, err := get(dataUrl)
	if err != nil {
		return result, errors.New("Data doesn't exist")
	}
    // var out bytes.Buffer
	// json.Indent(&out, data, "=", "\t")
	// out.WriteTo(os.Stdout)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, errors.New("Service response is unidentified")
	}
	return
}

func GetMetaOfCollection(id_collection string) (result interface{}, err error) {

    dataUrl := Conf["PE_API_GET_URI"] + Conf["PE_GET_URL"] + Conf["PE_INFO_HOST_COLLECTION"] + "id=" + id_collection
	data, err := get(dataUrl)
	if err != nil {
		return result, errors.New("Data doesn't exist")
	}
    // var out bytes.Buffer
	// json.Indent(&out, data, "=", "\t")
	// out.WriteTo(os.Stdout)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, errors.New("Service response is unidentified")
	}
	return
}

func GetCollectionData(id_collection string) (result interface{}, err error) {

    dataUrl := Conf["PE_API_GET_URI"] + Conf["PE_GET_URL"] + Conf["PE_COLLECTION_DATA"] + "resource_id=" + id_collection
	data, err := get(dataUrl)
	if err != nil {
		return result, errors.New("Data doesn't exist")
	}
    // var out bytes.Buffer
	// json.Indent(&out, data, "=", "\t")
	// out.WriteTo(os.Stdout)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, errors.New("Service response is unidentified")
	}
	return
}

func getAllData(start string, next string) (result interface{}, err error) {

    dataUrl := Conf["PE_API_GET_URI"] + start
	fmt.Println(start)
	data, err := get(dataUrl)
	if err != nil {
		return result, errors.New("Data doesn't exist")
	}
    // var out bytes.Buffer
	// json.Indent(&out, data, "=", "\t")
	// out.WriteTo(os.Stdout)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, errors.New("Service response is unidentified")
	}
	return
}

func GetCollectionDataWithQuery(id_collection string, query string, toSelect string) (result interface{}, err error) {
	var Url *url.URL
	Url, err = url.Parse("https://api.emploi-store.fr")
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	Url.Path += "/api/action/datastore_search_sql"
	parameters := url.Values{}
	parameters.Add("sql", "SELECT " + toSelect + " from \"" + id_collection + "\" " + query)
	Url.RawQuery = parameters.Encode()
	dataUrl := Url.String()
	data, err := get(dataUrl)
	if err != nil {
		return result, errors.New("Data doesn't exist")
	}
    // var out bytes.Buffer
	// json.Indent(&out, data, "=", "\t")
	// out.WriteTo(os.Stdout)
	var fromJSON interface{}
	err = json.Unmarshal(data, &fromJSON)
	if err != nil {
		return result, errors.New("Service response is unidentified")
	}
	result = fromJSON.(map[string]interface{})["result"].(map[string]interface{})["records"]
	return
}

func GetCollectionJSONWithQueryAndToken(id_collection string, query string, toSelect string, accessToken string) (result []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse("https://api.emploi-store.fr")
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	Url.Path += "/api/action/datastore_search_sql"
	parameters := url.Values{}
	parameters.Add("sql", "SELECT " + toSelect + " from \"" + id_collection + "\" " + query)
	Url.RawQuery = parameters.Encode()
	dataUrl := Url.String()
	data, err := getWithToken(dataUrl, accessToken)
	if err != nil {
		return result, errors.New("Data doesn't exist")
	}
    // var out bytes.Buffer
	// json.Indent(&out, data, "=", "\t")
	// out.WriteTo(os.Stdout)
	return data, nil
}

func CountCollectionDataWithQuery(id_collection string, query string, toSelect string) (result int, err error) {
	var Url *url.URL
	Url, err = url.Parse("https://api.emploi-store.fr")
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	Url.Path += "/api/action/datastore_search_sql"
	parameters := url.Values{}
	parameters.Add("sql", "SELECT COUNT(" + toSelect + ") from \"" + id_collection + "\" " + query)
	Url.RawQuery = parameters.Encode()
	dataUrl := Url.String()
	data, err := get(dataUrl)
	if err != nil {
		return result, errors.New("Data doesn't exist")
	}
	var count JSONCount
	err = json.Unmarshal(data, &count)
	if err != nil {
		return result, errors.New("Service response is unidentified")
	}
	result, err = strconv.Atoi(count.Result.Records[0].Count)
	if err != nil {
			return 0, err
	}
	return
}
