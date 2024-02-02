package gateway

import "server/zinsearch/models"

type ZincsearchClientGateway interface {
	Search(searchStruct *models.Search) (*models.SearchResponse, error)
}

type ZincsearchGateway struct {
	ZincsearchClientGateway
}

func NewZincsearchClientGateway(client *models.ZincsearchClient) *ZincsearchGateway {
	return &ZincsearchGateway{&SearchRequest{client}}
}
