package api

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetCustomerList(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gdb, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		t.Errorf("Failed to open gorm v2 db, got error: %v", err)
	}

	if gdb == nil {
		t.Error("gorm db is null")
	}

	t.Run("Should return empty array", func(t *testing.T) {
		expected := "{\"data\": []}"

		rows := sqlmock.NewRows(nil)
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "customers"  WHERE "customers"."deleted_at" IS NULL`)).
			WillReturnRows(rows)

		r := gin.Default()
		h := NewHandler(nil, gdb)
		h.SetupRoutes(&r.RouterGroup)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/customer", nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.JSONEq(t, expected, w.Body.String())

		err = mock.ExpectationsWereMet()
		if err != nil {
			t.Errorf("Failed to meet expectations, got error: %s", err.Error())
		}
	})
}
