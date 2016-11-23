package apipe

import (
	"encoding/json"
	"errors"
    "bytes"
    "os"
	"fmt"
	"net/url"
	"strconv"
)
//
// func GetCollectionDataID(collectionID string, accessToken string) (collectionDataID []string, err error) {
// 	collectionDataID = []string{}
// 	npage := 1
// 	total_pages := 0
//
// 	for ok := true; ok; ok = (npage <= total_pages) {
// 		collectionDataInfoUrl := Conf["CPA_API_URI"] + Conf["CPA_COLLECTION_URL"] + collectionID + "/" + Conf["CPA_COLLECTION_DATA_URL"] + Conf["POST_ACCESS_TOKEN"] + accessToken
// 		collectionDataInfoUrl += "&page[size]=200&page[number]=" + strconv.Itoa(npage)
// 		data, err := get(collectionDataInfoUrl)
// 		if err != nil {
// 			return nil, err
// 		}
// 		var content JSONContent
// 		err = json.Unmarshal(data, &content)
// 		if err != nil {
// 			return nil, errors.New("Service response is unidentified")
// 		}
// 		for _, value := range content.Data {
// 			collectionDataID = append(collectionDataID, value.Id)
// 		}
// 		npage++
// 		total_pages = content.Meta.Total_pages
// 	}
// 	return
// }
//
// func GetData(dataID string, collectionID string, accessToken string) (result JSONContentSingleData, err error) {
// 	dataUrl := Conf["CPA_API_URI"] + Conf["CPA_COLLECTION_URL"] + collectionID + "/" + Conf["CPA_COLLECTION_DATA_URL"] + dataID + "/" + Conf["POST_ACCESS_TOKEN"] + accessToken
// 	data, err := get(dataUrl)
// 	if err != nil {
// 		return result, errors.New("Data doesn't exist")
// 	}
// 	err = json.Unmarshal(data, &result)
// 	if err != nil {
// 		return result, errors.New("Service response is unidentified")
// 	}
// 	return
// }
//

func GetServicesList(accessToken string) (result JSONListOfServices, err error) {

    dataUrl := Conf["PE_API_GET_URI"] + Conf["PE_GET_URL"] + Conf["PE_LIST_COLLECTION"] + "id=digidata"
	data, err := get(dataUrl, accessToken)
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

func GetListOfCollectionsFromServicesId(accessToken string, id_services string) (result JSONListOfCollectionByServices, err error) {

    dataUrl := Conf["PE_API_GET_URI"] + Conf["PE_GET_URL"] + Conf["PE_INFO_COLLECTION"] + "id=" + id_services
	data, err := get(dataUrl, accessToken)
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

func GetMetaOfCollection(accessToken string, id_collection string) (result interface{}, err error) {

    dataUrl := Conf["PE_API_GET_URI"] + Conf["PE_GET_URL"] + Conf["PE_INFO_HOST_COLLECTION"] + "id=" + id_collection
	data, err := get(dataUrl, accessToken)
	if err != nil {
		return result, errors.New("Data doesn't exist")
	}
    var out bytes.Buffer
	json.Indent(&out, data, "=", "\t")
	out.WriteTo(os.Stdout)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, errors.New("Service response is unidentified")
	}
	return
}

func GetCollectionData(accessToken string, id_collection string) (result interface{}, err error) {

    dataUrl := Conf["PE_API_GET_URI"] + Conf["PE_GET_URL"] + Conf["PE_COLLECTION_DATA"] + "resource_id=" + id_collection
	data, err := get(dataUrl, accessToken)
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

func getAllData(accessToken string, start string, next string) (result interface{}, err error) {

    dataUrl := Conf["PE_API_GET_URI"] + start
	fmt.Println(start)
	data, err := get(dataUrl, accessToken)
	if err != nil {
		return result, errors.New("Data doesn't exist")
	}
    var out bytes.Buffer
	json.Indent(&out, data, "=", "\t")
	out.WriteTo(os.Stdout)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, errors.New("Service response is unidentified")
	}
	return
}

func GetCollectionDataWithQuery(accessToken string, id_collection string, query string, toSelect string) (result interface{}, err error) {
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
	data, err := get(dataUrl, accessToken)
	if err != nil {
		return result, errors.New("Data doesn't exist")
	}
    var out bytes.Buffer
	json.Indent(&out, data, "=", "\t")
	out.WriteTo(os.Stdout)
	var fromJSON interface{}
	err = json.Unmarshal(data, &fromJSON)
	if err != nil {
		return result, errors.New("Service response is unidentified")
	}
	result = fromJSON.(map[string]interface{})["result"].(map[string]interface{})["records"]
	return
}

func CountCollectionDataWithQuery(accessToken string, id_collection string, query string, toSelect string) (result int, err error) {
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
	data, err := get(dataUrl, accessToken)
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

// func GetDataWithQuery(query QueryModel, collectionID string, accessToken string) (result JSONContent, err error) {
// 	dataUrl := Conf["CPA_API_URI"] + Conf["CPA_COLLECTION_URL"] + collectionID + "/" + Conf["CPA_COLLECTION_DATA_URL"] + Conf["POST_ACCESS_TOKEN"] + accessToken
// 	for key, value := range query.Filter {
// 		dataUrl += "&filter[data." + key + "]=" + value
// 	}
// 	for key, value := range query.Page {
// 		dataUrl += "&page[" + key + "]=" + value
// 	}
// 	data, err := get(dataUrl)
// 	if err != nil || data == nil {
// 		return result, errors.New("Data doesn't exist:" + err.Error())
// 	}
// 	err = json.Unmarshal(data, &result)
// 	if err != nil {
// 		return result, errors.New("Parsing error")
// 	}
// 	return
// }
