package apipe

import (
)

type CPACollectionModel struct {
	Nom                        string `json:"nom"`
	Description                string `json:"description"`
	Tableau_de_donnees         bool   `json:"tableau_de_donnees"`
	Jeton_fc_lecture_ecriture  bool   `json:"jeton_fc_lecture_ecriture"`
	Jeton_fc_lecture_seulement bool   `json:"jeton_fc_lecture_seulement"`
}

type CPAPostModel struct {
	Access_token string   `json:"access_token"`
	Data         JSONData `json:"data"`
}

type PEToken struct {
	Access_token string `json:"access_token"`
}

type JSONListOfCollection struct {
	Result JSONDataResult         `json:"result"`
}

type JSONDataResult struct {
    Packages []JSONDataResultPackages `json:"packages"`
}

type JSONDataResultPackages struct {
    Id string `json:"id"`
    Title string `json:"title"`
    Revision_id string `json:"revision_id"`
}

type JSONPoleEmploiToken struct {
    Scope           string  `json:"digidata"`
    Expires_in      int     `json:"expires_in"`
    Token_type      string  `json:"token_type"`
    Access_token    string  `json:"access_token"`
}

type JSONData struct {
	Id         string      `json:"id"`
	Type       string      `json:"type"`
	Attributes interface{} `json:"attributes"`
	Links      struct {
		Self  string `json:"self"`
		First string `json:"first"`
	} `json:"links"`
	Meta struct {
		Creation     string `json:"creation"`
		Modification string `json:"modification"`
		Version      int    `json:"version"`
	}
}

type JSONMetaPagination struct {
	Total       int `json:"total"`
	Total_pages int `json:"total_pages"`
	Offset      int `json:"offset"`
	Limit       int `json:"limit"`
	Count       int `json:"count"`
}

type QueryModel struct {
	Filter map[string]string `json:"filter"`
	Page   map[string]string `json:"page"`
}
