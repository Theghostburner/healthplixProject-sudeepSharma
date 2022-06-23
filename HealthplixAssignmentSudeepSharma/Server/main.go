package main

import (
	"net/http"
	//"strings"
	"time"

	"github.com/gin-gonic/gin"

	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	db, err := sql.Open("mysql", "root:SoumiliDas1#@tcp(127.0.0.1:3306)/healthplixAssignmentDB")
	if err != nil {
        panic(err.Error())
    }
    defer db.Close()


	router := gin.Default()
	router.GET("/medicine", getMedicines)
	router.POST("/medicine", postMedicine)
	router.GET("/medicine/:name", getMedicineByName)
	router.DELETE("/medicine/:name",deleteMedicineByName)
	router.PATCH("/medicine/:",updateMedicineByName)
	router.Run("localhost:8080")
}
type Timestamp struct {
	time.Time
}
type medicine struct {
    created_on Timestamp `json:"created_on"`
	updated_on Timestamp `json:"updated_on"`
	updated_by string `json:"updated_by"`
	created_by string `json:"created_by"`
    medicine_name string `json:"medicine_name"`
	manufacturer string `json:"manufacturer"`
	medicine_price int `json:"medicine_price"`
	medicine_id int` json:"medicine_id"`
}
type user struct {
    created_on Timestamp `json:"created_on"`
	updated_on Timestamp `json:"updated_on"`
	updated_by string `json:"updated_by"`
	created_by string `json:"created_by"`
    user_name string `json:"user_name"`
	access_level int `json:"access_level"`
	user_email_id string `json:"user_email_id"`
	user_password string ` json:"user_password"`
	user_id int ` json:"user_id"`
}

func getMedicines(c *gin.Context) {
	//c.IndentedJSON(http.StatusOK, albums)
	db, err := sql.Open("mysql", "root:SoumiliDas1#@tcp(127.0.0.1:3306)/healthplixAssignmentDB")
	if err != nil {
        panic(err.Error())
    }
    defer db.Close()
	results, err := db.Query("SELECT * FROM medicines")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
	for results.Next() {
        var tag Tag
        // for each row, scan the result into our tag composite object
        err = results.Scan(&tag.ID, &tag.Name)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
                // and then print out the tag's Name attribute
        log.Printf(tag.Name)
    }

}

}
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
