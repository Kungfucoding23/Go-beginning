package main

import (
	"github.com/labstack/echo"
	"net/http"
)

//Response ...
type Response struct {
	Message string `json:"message" xml:"message"`
	Status  string `json:"status" xml:"status"`
}

// HelloHandlerPrettyJSON ...
func (r Response) HelloHandlerPrettyJSON(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, r, "  ")
}

// HelloHandlerPrettyXML ...
func (r Response) HelloHandlerPrettyXML(c echo.Context) error {
	return c.XMLPretty(http.StatusOK, r, "  ")
}

func main() {
	e := echo.New()
	r := Response{
		Message: "Hello from Echo!",
		Status:  "OK",
	}
	//make routes
	e.GET("/", r.HelloHandlerPrettyJSON)
	e.GET("/xml", r.HelloHandlerPrettyXML)
	e.Logger.Fatal(e.Start(":1323"))
}
