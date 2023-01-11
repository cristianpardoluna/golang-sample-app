package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var router = setupRouter()
var w = httptest.NewRecorder()

func TestGetRecords(t *testing.T) {
	req, _ := http.NewRequest("GET", "/records", nil)
	router.ServeHTTP(w, req)

	// assert if status code is 200
	assert.Equal(t, 200, w.Code)

	// assert if body parsed is an array of
	// objects w/ string-based mapping
	arr := []interface{}{map[string]interface{}{}}
	bodyString := w.Body.String()
	var bodyParsed []interface{}
	if err := json.Unmarshal([]byte(bodyString), &bodyParsed); err != nil {
		fmt.Println(err)
	}
	assert.IsTypef(t, arr, bodyParsed,
		"Error message %s", "formatted")

}

func TestGetRecordByID(t *testing.T) {
	recordID := 1
	req, _ := http.NewRequest("GET",
		fmt.Sprintf("/records/%v", recordID), nil)
	router.ServeHTTP(w, req)

	// assert if status code is 200
	assert.Equalf(t, 200, w.Code,
		"Error message %s", "formatted")

}

func TestPostRecord(t *testing.T) {
	type Record struct {
		id   string
		name string
		val  float64
	}
	newRecord := Record{id: "666",
		name: "test record", val: 0.6}

	jsonData, err := json.Marshal(newRecord)
	if err != nil {
		fmt.Printf("Json error: %v", err)
	}
	req, _ := http.NewRequest("POST", "/records",
		bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)

	// assert if status code is 200
	assert.Equalf(t, 200, w.Code,
		"Error message %s", "formatted")

}
