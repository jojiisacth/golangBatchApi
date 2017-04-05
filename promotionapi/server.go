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

	router.POST("/products", getProducts2)

	router.GET("/products", getProducts)
	router.GET("/price", getPrice)
	// router.GET("/promotion", getPrice)

	router.GET("/version", getVersion)

	router.GET("/healthz", getHealth)

	router.Logger.Fatal(router.Start(":8082"))

}

func getProducts2(c echo.Context) error {
	products := make([]product, 6)
	x, _ := c.MultipartForm()
	y := c.ParamNames()

	if x == nil {

	}

	if y == nil {

	}

	products[0] = product{"Post Data sweet mangoes", "1222", "mangoes of the time"}
	products[1] = product{"greate grapes", "23455", "grapes of the time"}
	products[2] = product{"greate grapes", "5523455", "grapes of the time"}
	products[3] = product{"Tesco Gala Apple Minimum 5 Pack", "213441", "Tesco Gala Apple Minimum 5 Pack"}
	products[4] = product{"Tesco Braeburn Apple Minimum 5 Pack 670G", "2344333", ""}
	products[5] = product{"Tesco Apple Juice 1 Litre", "0987654444", ""}

	er := c.JSON(200, products)
	return er

}

func getProducts(c echo.Context) error {
	products := make([]product, 6)

	x := c.Request

	if x == nil {
		fmt.Println("empty")
	}
	products[0] = product{"sweet mangoes", "1222", "mangoes of the time"}
	products[1] = product{"greate grapes", "23455", "grapes of the time"}
	products[2] = product{"greate grapes", "5523455", "grapes of the time"}
	products[3] = product{"Tesco Gala Apple Minimum 5 Pack", "213441", "Tesco Gala Apple Minimum 5 Pack"}
	products[4] = product{"Tesco Braeburn Apple Minimum 5 Pack 670G", "2344333", ""}
	products[5] = product{"Tesco Apple Juice 1 Litre", "0987654444", ""}

	er := c.JSON(200, products)
	return er

}

func getPrice(c echo.Context) error {
	prices := make([]price, 6)

	prices[0] = price{"1222", "2654", "76", "$"}
	prices[1] = price{"23455", "23455", "76", "$"}
	prices[2] = price{"5523455", "2564", "76", "$"}
	prices[3] = price{"213441", "6542", "76", "$"}
	prices[4] = price{"2344333", "2344", "76", "$"}
	prices[5] = price{"0987654444", "5256", "76", "$"}
	er := c.JSON(200, prices)
	return er
}

func getVersion(c echo.Context) error {
	return c.String(http.StatusOK, "v1")
}
func getHealth(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

type product struct {
	Name        string `json:"pname"`
	Id          string `json:"id"`
	Description string
}

type price struct {
	ProductId string `json:"pid"`
	Storeid   string `json:"sid"`
	Value     string
	Currency  string
}
