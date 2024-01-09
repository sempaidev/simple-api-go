package main
import "github.com/gin-gonic/gin"
import "net/http"

// Creamos una estructura para los albunes
type album struct {
	ID string `json "id"`
	Title string `json "title"`
	Artist string `json "artist"`
	Year uint16 `json "year"`
}

// Creamos diferentes albunes del typo de la estructura
var albums = []album{
	{ID: "1", Title: "Familia", Artist: "Camila Cabello", Year: 2022},
	{ID: "2", Title: "21", Artist: "Adele", Year: 2022},
	{ID: "3", Title: "The Eminem Show", Artist: "Eminem", Year: 2022},
	{ID: "4", Title: "Sin Nombre Vol1", Artist: "Rookie", Year: 2022},
}

// Creamos una funcion para obtener todos los albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// Creamos funcion para guardar un nuevo album
func postAlbum(c *gin.Context) {
	var newAlbum album
	c.BindJSON(&newAlbum)
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// Funcion para obtener un album por id
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
	// Definimos las rutas para acceder a las diferentes funciones
	router.GET("/albums", getAlbums)
	router.POST("/album", postAlbum)
	router.GET("/album/:id", getAlbumsById)
	// Ruta principal del api y su puerto
	router.Run("localhost:8080")
}