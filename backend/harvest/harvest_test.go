package harvest

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"../utils"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var userJSON = `{ "plantType": "Lettuce" }`

func TestHarvest(t *testing.T) {
	database, err := sql.Open("sqlite3", "../rowa.db")
	if err != nil {
		fmt.Println(err)
	}
	db := &utils.DbObject{database}
	db.Init()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/harvest/harvest", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, Harvest(c, db)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		fmt.Println(rec.Body.String())
	}

}
