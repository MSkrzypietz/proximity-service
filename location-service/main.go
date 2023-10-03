package main

import (
	"encoding/json"
	"fmt"
	"github.com/MSkrzypietz/proximity-service/location-service/geohash"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type NearbySearchRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Radius    float64 `json:"radius,omitempty"`
}

func main() {
	router := httprouter.New()
	router.POST("/v1/search/nearby", func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		var search NearbySearchRequest
		decoder := json.NewDecoder(req.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&search)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		res.Write([]byte(geohash.CalcGeohash(search.Latitude, search.Longitude)))
	})
	log.Fatal(http.ListenAndServe(":3001", router))
}
