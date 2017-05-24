package controllers

import (
	"github.com/revel/revel"
	"time"
	"fmt"
	"projects/start/app/Models"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"os"
)

type App struct {
	*revel.Controller
}

type SuccessfulDownload struct {
	Ok Models.DataObject
}

type RetrievedData struct {
	Data []Models.DataObject
}

func (c App) Index() revel.Result {
	file, _ := os.OpenFile(revel.BasePath+"/app/index.html", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	//file := http.FileServer(rice.MustFindBox("http-files").HTTPBox())

	return c.RenderFile(file, revel.Attachment )
}

func (c App) Bundle() revel.Result {
	return c.Render()
}

func (c App) Downloaded() revel.Result {
	// IF IP HAS BEEN PASSED AS AN ARGUMENT USE THESE FUNCTIONS

	//ipRequest := c.Request.RemoteAddr
	// c.Request.ClientIp
	// c.Request.Body

	/*	localIp := Utils.GetLocalIp()
		// for local testing purposes
		location := Utils.GeoIp(localIp)
		*/

	// INSTEAD IF WE GET A POST REQUEST WITH THE LATITUDE AND LONGITUDE VALUE WE USE THESE FUNCTIONS
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	bodyParsed := gjson.GetMany(string(b), "longitude", "latitude", "Country", "Age", "Sex")

	download := Models.DataObject{
		Longitude:    bodyParsed[0].String(),
		Latitude:     bodyParsed[1].String(),
		DownloadedAt: time.Now().String(),
		Country:      bodyParsed[2].String(),
		Age:          bodyParsed[3].String(),
		Sex:          bodyParsed[4].String(),
	}

	Models.DbSession.NewRecord(&download)

	Models.DbSession.Create(&download)

	Models.DbSession.Save(&download)

	response := SuccessfulDownload{Ok: download}

	return c.RenderJSON(response)
}

func (c App) GetDownloadData() revel.Result {
	result := []Models.DataObject{}
	Models.DbSession.Find(&result)
	fmt.Print(result)
	return c.RenderJSON(result)
}
