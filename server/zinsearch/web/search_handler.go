package web

import (
	"encoding/json"
	"net/http"
	"server/zinsearch/gateway"
	"server/zinsearch/models"
)

type SearchHandler struct {
	gateway.ZincsearchClientGateway
}

func (searchHandler *SearchHandler) SearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	searchRequest, err := parseSearchRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	searchResult, err := searchHandler.Search(searchRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(searchResult)
}

func NewSearchHandler(client *models.ZincsearchClient) *SearchHandler {
	return &SearchHandler{gateway.NewZincsearchClientGateway(client)}
}

func parseSearchRequest(r *http.Request) (*models.Search, error) {
	body := r.Body
	defer body.Close()
	var cmd models.Search
	_ = json.NewDecoder(body).Decode(&cmd)
	return &cmd, nil
}
