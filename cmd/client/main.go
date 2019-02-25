package main

import (
	"log"

	"github.com/larwef/cert-monitor/pkg/cert"
	"github.com/larwef/cert-monitor/pkg/config"
)

func main() {
	conf := config.New("configs/config.json")

	client := cert.NewClient(conf)

	req := cert.BuypassTestRequest("993884871")

	res, err := client.Search(req)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, elem := range res.Certs {
		log.Printf("%+v\n", elem)
	}

}
