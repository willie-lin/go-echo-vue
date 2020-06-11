package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

func GetIpInfo()  echo.HandlerFunc {
	return func(c echo.Context) error {

		ip := c.Param("ip")
		var baseUrl string = "http://ip-api.com/json/"
		var status string = "?fields=status,message,continent,continentCode,country,countryCode,region,regionName,city,district,zip,lat,lon,timezone,currency,isp,org,as,asname,reverse,mobile,proxy,hosting,query"

		url := baseUrl + ip + status
		//url := baseUrl + string(ip) + status
		response, _ := http.Get(url)
		defer response.Body.Close()
		body, _ :=  ioutil.ReadAll(response.Body)
		fmt.Println(body)
		//fmt.Println(string(body))
		//str := string(body)
		return c.JSONBlob(http.StatusOK, body)

	}
}
