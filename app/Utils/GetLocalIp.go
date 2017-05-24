package Utils

import (
	"net/http"
	"os"
	"io/ioutil"
	"github.com/tidwall/gjson"
)

type LocalIp struct {
	Ip string `json:"ip"`
}

func GetLocalIp () string {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	value := gjson.Get(string(b), "ip")



	return value.String()

}