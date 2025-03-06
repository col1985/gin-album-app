package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	// swagger embed files
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


type Album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

var albums = []Album{
    {ID: "1", Title: "2001", Artist: "Dr. Dre", Price: 56.99},
    {ID: "2", Title: "Rumours", Artist: "Fleetwood Mac", Price: 67.99},
    {ID: "3", Title: "Dark side of the Moon", Artist: "Pink Floyd", Price: 39.99},
    {ID: "4", Title: "Songs from the Dead", Artist: "Queens of the Stone Age", Price: 56.99},
    {ID: "5", Title: "Foo Fighters", Artist: "Foo Fighters", Price: 30.99},
}

//	@title			Gin Album Demo API
//	@version		1.0
//	@description	This is a basic Gin API app for demos.

//	@contact.name	Colum Bennett
//	@contact.email	col1985@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

func main() {
    router := gin.Default()

    sys := router.Group("/sys")
    {
        sys.GET("ping", GetStatus)
        sys.GET("metrics", gin.WrapH(promhttp.Handler()))
    }

    v1 := router.Group("/api/v1")
    {
        albums := v1.Group("/albums")
        {
            albums.GET("", GetAlbums)
            albums.GET(":id", GetAlbumByID)
            albums.POST("", PostAlbums)
        }
    }


    router.GET("/api-docs", ginSwagger.WrapHandler(swaggerFiles.Handler))

    err := router.Run(":8080")
    if err != nil {
        println("Unable to start router.")
        return
    }
}

//	@Summary		System Health Endpoint
//	@Description	Get status response of app
//	@Tags			System Ping
//	@Success		200	{string}	pong
//	@Router			/sys/ping [get]
func GetStatus(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, "pong")
}

//	@Summary		Get List of Albums Endpoint
//	@Description	Get List of Albums
//	@Tags			Albums List
//	@Success		200	{array}	Album
//	@Router			/albums [get] Album
func GetAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

//	@Summary		Create a new album
//	@Description	Add a new album to list of Albums
//	@Tags			Albums List
//	@Success		200	{object}	Album
//	@Router			/albums [post] Album
func PostAlbums(c *gin.Context) {
    var newAlbum Album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

//	@Summary		Get Album by ID
//	@Description	Get an album from the list using ID property
//	@Tags			Albums
//	@Success		200	{object}	Album
//	@Router			/albums/:id [get] Album
func GetAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}