package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

func plant(c echo.Context) (err error) {
	fmt.Println("This is server: request received")
	u := new(User)
	if err = c.Bind(u); err != nil {
		return
	}
	ret := c.JSON(http.StatusOK, u)
	fmt.Println(u)
	return ret
}
func main() {
	// Echo instance
	e := echo.New()

	e.Use(middleware.CORS())

	// Routes
	//e.File("/", "index.html")
	e.POST("/plant/greet", plant)
	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
