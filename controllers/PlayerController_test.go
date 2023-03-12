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
	// mock the service
	playerService := new(mocks.IPlayerService)

	// names
	player1Name := "Rafael"
	player2Name := "Serena"

	// should get this score
	expectedScore := "Forty-Fifteen"
	expectedResult := viewmodels.ScoresVM{}
	expectedResult.Score = expectedScore

	// setup service return expectations
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
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}

func TestPlayerScoreNoRecord(t *testing.T) {
	// mock the service
	playerService := new(mocks.IPlayerService)

	// names
	player1Name := "fart"
	player2Name := "Rafael"

	// should get an error
	expectedResult := ResponseError{}
	expectedResult.Message = recordNotFound

	// setup service return expectations
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
	assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
}

func TestPlayerScoreUnknownError(t *testing.T) {
	// mock the service
	playerService := new(mocks.IPlayerService)

	// names
	player1Name := "noway"
	player2Name := "fart"

	// should get an error
	expectedResult := ResponseError{}
	expectedResult.Message = "Unexpected error."

	// setup service return expectations
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
	assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
}
