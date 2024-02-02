package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"server/zinsearch/models"
)

type SearchRequest struct {
	*models.ZincsearchClient
}

func CreateUrl(searchStruct *SearchRequest) string {
	formattedUrl := fmt.Sprintf("%s:%d/api/mails/_search", searchStruct.ServerHost, searchStruct.ServerPort)
	return formattedUrl
}

func (searchRequest *SearchRequest) Search(searchStruct *models.Search) (*models.SearchResponse, error) {
	jsonData, err := json.Marshal(searchStruct)
	if err != nil {
		return nil, err
	}
	url := CreateUrl(searchRequest)
	fmt.Println(url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", searchRequest.Authorization)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	bodyResponse := resp.Body
	defer bodyResponse.Close()
	var searchResponse models.SearchResponse
	_ = json.NewDecoder(bodyResponse).Decode(&searchResponse)
	fmt.Println(resp.StatusCode)
	return &searchResponse, nil
}
