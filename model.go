package apipe

import (
)

type PEToken struct {
	Access_token string `json:"access_token"`
}

type JSONListOfServices struct {
	Result JSONDataResultOfServices         `json:"result"`
}

type JSONDataResultOfServices struct {
    Packages []JSONDataResultPackagesOfServices `json:"packages"`
}

type JSONDataResultPackagesOfServices struct {
    Id string `json:"id"`
    Title string `json:"title"`
    Revision_id string `json:"revision_id"`
}

type JSONListOfCollectionByServices struct {
	Result JSONDataResultOfCollection         `json:"result"`
}

type JSONDataResultOfCollection struct {
    Resources []JSONDataResultResourcesOfCollection `json:"resources"`
}

type JSONDataResultResourcesOfCollection struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Revision_id string `json:"revision_id"`
}

type JSONDataWithQueryResult struct {
	Result JSONResultDataWithQuery      `json:"result"`
}

type JSONResultDataWithQuery struct {
    Links JSONDataLink `json:"_links"`
}

type JSONDataLink struct {
    Start string `json:"start"`
    Next string `json:"next"`
}

type JSONCount struct {
    Result JSONResultCount `json:"result"`
}

type JSONResultCount struct {
    Records []JSONRecordCount `json:"records"`
}

type JSONRecordCount struct {
	Count string `json:"count"`
}

type JSONPoleEmploiToken struct {
    Scope           string  `json:"digidata"`
    Expires_in      int     `json:"expires_in"`
    Token_type      string  `json:"token_type"`
    Access_token    string  `json:"access_token"`
}

type QueryModel struct {
	Filter map[string]string `json:"filter"`
	Page   map[string]string `json:"page"`
}
