package main

import (
	"os"
	"server/zinsearch/models"
	"server/zinsearch/web"
	"strconv"
)

func main() {
	zincsearchBaseUrl := os.Getenv("ZINCSEARCH_API_URL")
	zincsearchPort, err := strconv.Atoi(os.Getenv("ZINCSEARCH_API_PORT"))
	zincsearchAuthorization := os.Getenv("ZINCSEARCH_API_AUTHORIZATION")
	if err != nil {
		panic(err)
	}

	searchHandler := web.NewSearchHandler(&models.ZincsearchClient{
		ServerHost:    zincsearchBaseUrl,
		ServerPort:    zincsearchPort,
		Authorization: zincsearchAuthorization,
	})
	mux := Routes(searchHandler)
	server := NewServer(mux)
	server.Start()
}
