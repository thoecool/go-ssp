package main

import (
	"net/http"
	"usecases"
)

func main() {
	campaignInteractor := new(usecases.CampaignInteractor)
	http.HandleFunc("/api", func(res http.ResponseWriter, req *http.Request) {
		//webserviceHandler.ShowOrder(res, req)
	})
	http.ListenAndServe(":8080", nil)
}
