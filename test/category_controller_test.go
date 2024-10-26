package test

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"rest-api-native/app"
	"rest-api-native/controller"
	"rest-api-native/helper"
	"rest-api-native/middleware"
	"rest-api-native/repository"
	"rest-api-native/service"
	"strings"
	"testing"
	"time"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:pw@tcp(localhost:3306)/golangtest")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}

func setupRouter() http.Handler {
	validate := validator.New()
	db := setupTestDB()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	return middleware.NewAuthMiddleware(app.NewRouter(categoryController))
}

func TestCreateCategorySuccess(t *testing.T) {
	router := setupRouter()

	requestBody := strings.NewReader(`{"name":"Electronic"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestCreateCategoryFailed(t *testing.T) {
	router := setupRouter()

	// Test case 1: Empty name
	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	// Test case 2: Invalid JSON
	requestBody = strings.NewReader(`{"name":123}`)
	request = httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response = recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
}

func TestUpdateCategorySuccess(t *testing.T) {
	router := setupRouter()

	requestBody := strings.NewReader(`{"name":"Updated Electronic"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/1", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestUpdateCategoryFailed(t *testing.T) {
	router := setupRouter()

	// Test case 1: Invalid ID
	requestBody := strings.NewReader(`{"name":"Updated Electronic"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/999", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	// Test case 2: Empty name
	requestBody = strings.NewReader(`{"name":""}`)
	request = httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/1", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response = recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
}

func TestFindAllCategoriesSuccess(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestFindAllCategoriesFailed(t *testing.T) {
	router := setupRouter()

	// Test without auth header
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)
}

func TestFindCategoryByIdSuccess(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/1", nil)
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestFindCategoryByIdFailed(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/999", nil)
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}

func TestDeleteCategorySuccess(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/1", nil)
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestDeleteCategoryFailed(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/999", nil)
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}

func TestUnauthorized(t *testing.T) {
	router := setupRouter()

	// Test with invalid API Key
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("X-API-Key", "INVALID-KEY")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	// Test without API Key
	request = httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response = recorder.Result()
	assert.Equal(t, 401, response.StatusCode)
}

func TestCleanupDB(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	// Cleanup test data
	_, err := db.Exec("DELETE FROM categories WHERE id > 0")
	assert.NoError(t, err)
}
