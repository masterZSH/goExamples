package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCase(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/v1/area/1", nil)
	GetArea(c)
	assert.Equal(t, http.StatusOK, w.Code)

	var got gin.H
	json.Unmarshal(w.Body.Bytes(), &got)
	errcode, find := got["errcode"]
	assert.True(t, find)
	assert.Equal(t, float64(0), errcode)
	rs := got["result"].([]interface{})
	assert.NotEmpty(t, len(rs))

}
