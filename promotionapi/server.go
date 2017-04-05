package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	fmt.Println("hi Go")

	router := echo.New()

	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "wellcome to promition info service :)!")
	})

	router.POST("/promotion/", getPromotion)

	router.GET("/version", getVersion)

	router.GET("/healthz", getHealth)

	router.Logger.Fatal(router.Start(":8082"))

}

func getPromotion(c echo.Context) error {
	promotions := make([]promotion, 4)

	promotions[0] = promotion{"1222", "2654", "Percentage", "10"}
	promotions[1] = promotion{"23455", "2455", "Percentage", "15"}
	promotions[2] = promotion{"5523455", "2564", "Buytow", "GetOneFree"}
	promotions[3] = promotion{"213441", "6542", "MonetoryDiscount", "2"}

	er := c.JSON(200, promotions)
	return er
}

func getVersion(c echo.Context) error {
	return c.String(http.StatusOK, "v1")
}
func getHealth(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

type promotion struct {
	ProductId     string `json:"pid"`
	Storeid       string `json:"sid"`
	PromotionType string
	Value         string
}
