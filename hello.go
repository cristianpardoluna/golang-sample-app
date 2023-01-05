// package main is always when is not a library ('standalone' app)
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// data structure for record object
type record struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Val  float64 `json:"val"`
}

// records slice to seed record data
var records = []record{
	{ID: "2", Name: "SSS", Val: 0.0},
	{ID: "1", Name: "ASASS", Val: 0.001},
	{ID: "5", Name: "hoalaa", Val: 3.31},
}

// program
func main() {
	// := -> short declaration form
	// infer the type object declared
	router := gin.Default()

	// function: GET RECORDS
	router.GET("/records", getRecords)

	// function: POST RECORD
	router.POST("/records", postRecord)

	// function: GET RECORD BY ID
	router.GET("/records/:id", getRecordByID)

	// TODO: create test cases for gin app
	router.Run("localhost:8000")
}

// handler GET RECORDS: responds w/ list of records as JSON
func getRecords(c *gin.Context) {
	// passes all records
	c.IndentedJSON(http.StatusOK, records)
}

// handler POST RECORD: creates new record
func postRecord(c *gin.Context) {
	var newRecord record

	// binds JSON passed w/ structure
	// is this JSON validation before inserting in DB?
	if err := c.BindJSON(&newRecord); err != nil {
		fmt.Print(err)
		return
	}

	// Appends new record to slice
	records = append(records, newRecord)
	c.IndentedJSON(http.StatusCreated, newRecord)
}

// handler GET RECORD: gets new record by ID
func getRecordByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range records {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}
