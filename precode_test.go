package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, "/cafe?city=moscow&count="+strconv.Itoa(totalCount+1), nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	res := responseRecorder.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)

	// здесь нужно добавить необходимые проверки
	assert.NoError(t, err)
	resCafeList := strings.Split(string(data), ",")
	assert.Equal(t, totalCount, len(resCafeList), "Response body should return "+strconv.Itoa(totalCount)+" string value")
	assert.Equal(t, cafeList["moscow"], resCafeList, "Response body should contain all cafe from the given city")
}

func TestMainHandlerWhenCityNotSupported(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, "/cafe?city=city&count="+strconv.Itoa(totalCount), nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	res := responseRecorder.Result()

	// здесь нужно добавить необходимые проверки
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "Response status code should be 200")
	data, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "Response should not produce error")
	assert.Equal(t, "wrong city value", string(data), "Response body should return error string")
}

func TestMainHandlerWhenRequestIsCorrect(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, "/cafe?city=moscow&count="+strconv.Itoa(totalCount), nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Result().StatusCode, "Response status code should be 200")
	assert.NotEmpty(t, responseRecorder.Result().Body, "Response body should not be empty")
}
