package harvest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var userJSON = `{ "plantType": "Lettuce" }`


func TestHarvest(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/harvest/harvest", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, Harvest(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		fmt.Println(rec.Body.String())
	}

}
