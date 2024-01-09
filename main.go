package main
import "github.com/gin-gonic/gin"
import "net/http"

type album struct {
	ID string `json "id"`
	Title string `json "title"`
	Artist string `json "artist"`
	Year uint16 `json "year"`
}

var albums = []album{
	{ID: "1", Title: "Familia", Artist: "Camila Cabello", Year: 2022},
	{ID: "2", Title: "21", Artist: "Adele", Year: 2022},
	{ID: "3", Title: "The Eminem Show", Artist: "Eminem", Year: 2022},
	{ID: "4", Title: "Sin Nombre Vol1", Artist: "Rookie", Year: 2022},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album
	c.BindJSON(&newAlbum)
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Id no encontrado"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/album", postAlbum)
	router.GET("/album/:id", getAlbumsById)
	router.Run("localhost:8080")
}