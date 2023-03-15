package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	ce "github.com/irahardianto/service-pattern-go/customerrors"
	"github.com/irahardianto/service-pattern-go/interfaces/mocks"
	"github.com/irahardianto/service-pattern-go/viewmodels"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

const getScores = "GetScores"
const routerPattern = "/getScore/{player1}/vs/{player2}"
const recordNotFound = "Record not found."

func testRequest(player1 string, player2 string) *http.Request {
	return httptest.NewRequest("GET",
		fmt.Sprintf(
			"http://localhost:8080/getScore/%v/vs/%v", player1, player2), nil)
}

func TestPlayerScore(t *testing.T) {
	// names
	player1Name := "Rafael"
	player2Name := "Serena"

	// expectations
	expectedScore := "Forty-Fifteen"
	expectedResult := viewmodels.ScoresVM{}
	expectedResult.Score = expectedScore
	expectedStatus := http.StatusOK

	// setup mock service
	playerService := new(mocks.IPlayerService)
	playerService.On(getScores, player1Name, player2Name).Return(expectedScore, nil)
	playerController := PlayerController{playerService}

	// make the request
	req := testRequest(player1Name, player2Name)
	w := httptest.NewRecorder()
	r := chi.NewRouter()
	r.HandleFunc(routerPattern, playerController.GetPlayerScore)
	r.ServeHTTP(w, req)

	// decode result
	actualResult := viewmodels.ScoresVM{}
	json.NewDecoder(w.Body).Decode(&actualResult)

	// check the value
	assert.Equal(t, expectedResult, actualResult)
	// check the status code
	assert.Equal(t, expectedStatus, w.Result().StatusCode)
}

func TestPlayerScoreNoRecord(t *testing.T) {
	// names
	player1Name := "fake"
	player2Name := "Rafael"

	// expectations
	expectedResult := ResponseError{}
	expectedResult.Message = recordNotFound
	expectedStatus := http.StatusNotFound

	// setup mock service
	playerService := new(mocks.IPlayerService)
	playerService.On(getScores, player1Name, player2Name).Return("", ce.ErrRecordNotFound)
	playerController := PlayerController{playerService}

	// make the request
	req := testRequest(player1Name, player2Name)
	w := httptest.NewRecorder()
	r := chi.NewRouter()
	r.HandleFunc(routerPattern, playerController.GetPlayerScore)
	r.ServeHTTP(w, req)

	// decode the result
	actualResult := ResponseError{}
	json.NewDecoder(w.Body).Decode(&actualResult)

	// check the value
	assert.Equal(t, expectedResult, actualResult)
	// check the status code
	assert.Equal(t, expectedStatus, w.Result().StatusCode)
}

func TestPlayerScoreUnknownError(t *testing.T) {
	// names
	player1Name := "Rafael"
	player2Name := "fake"

	// expectations
	expectedResult := ResponseError{}
	expectedResult.Message = "Unexpected error."
	expectedStatus := http.StatusInternalServerError

	// setup mock service
	playerService := new(mocks.IPlayerService)
	playerService.On(getScores, player1Name, player2Name).Return("", errors.New("Weird error"))
	playerController := PlayerController{playerService}

	// make the request
	req := testRequest(player1Name, player2Name)
	w := httptest.NewRecorder()
	r := chi.NewRouter()
	r.HandleFunc(routerPattern, playerController.GetPlayerScore)
	r.ServeHTTP(w, req)

	// decode the result
	actualResult := ResponseError{}
	json.NewDecoder(w.Body).Decode(&actualResult)

	// check the value
	assert.Equal(t, expectedResult, actualResult)
	// check the status code
	assert.Equal(t, expectedStatus, w.Result().StatusCode)
}
