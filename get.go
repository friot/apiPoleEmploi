package apipe

import (
	"encoding/json"
	"errors"
    "bytes"
    "os"
	"fmt"
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

func GetCollectionDataWithQuery(accessToken string, id_collection string, query string, toSelect string) (result JSONDataWithQueryResult, err error) {
	//query = "WHERE \"CITY_NAME\" = 'METZ'LIMIT 50"
    //dataUrl := Conf["PE_API_GET_URI"] + Conf["PE_GET_URL"] + Conf["PE_COLLECTION_DATA_SQL"] + "sql=SELECT " + toSelect + " from \"" + id_collection + "\" " + query
	//dataUrl := "https://api.emploi-store.fr/api/action/datastore_search_sql?sql=SELECT * from \"" + id_collection + "\" WHERE DEPARTEMENT_CODE LIKE '57'"
	//dataUrl := "https://api.emploi-store.fr/api/action/datastore_search?resource_id=" + id_collection + "&DEPARTEMENT_CODE=57" // got an error
	dataUrl := "https://api.emploi-store.fr/api/action/datastore_search?resource_id=" + id_collection + "&q=57+MOSELLE"
	fmt.Println(dataUrl)
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
	fmt.Println(getAllData(accessToken, result.Result.Links.Start, ""))
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
