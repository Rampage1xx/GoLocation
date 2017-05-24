package Utils

import (
	"net/http"
	"log"
	"io/ioutil"
	"github.com/tidwall/gjson"
)


type Location struct {
	Longitude string
	Latitude string
}

func GeoIp (ip string)  Location  {
	resp, err := http.Get("http://freegeoip.net/json/" + ip)
	if err != nil {
		log.Fatalf("Got error while fetching the ip location, the error is '%v'", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	value := gjson.GetMany(string(b), "latitude", "longitude")

	var location = Location{Latitude:value[0].String(), Longitude:value[0].String()}

	return location
}