package main

import (
	"fmt"

	"github.com/labstack/echo"
)

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	// term := strings.ToLower(scapper.CleanString(c.FormValue("term")))
	fmt.Println(c.FormValue("term"))
	return nil
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	// e.POST("/scrape")
	e.Logger.Fatal(e.Start(":1323"))
}
