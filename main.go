package main

import (
	"log"
	"net/http"
	"os"
	"io"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)


func main() {
	//http.HandleFunc("/", handler)
	http.HandleFunc("/ok", handler_ok)

	port := os.Getenv("PORT")
	//log.Debug("port:%d", port)

	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		panic(err)
	}
}



func main_router() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", get_root)
	router.GET("/ok", get_ok)

	router.Run(":" + port)
}

func get_root(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)

}



func get_ok(c *gin.Context) {
	c.String(http.StatusOK, "Hello OK")
	
}

func handler_ok(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
        var w io.Writer = rw
	io.WriteString(w, "OK")
}