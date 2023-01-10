package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoRecords(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
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
