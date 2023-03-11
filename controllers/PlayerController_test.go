package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"

	ce "github.com/irahardianto/service-pattern-go/customerrors"
	"github.com/irahardianto/service-pattern-go/interfaces/mocks"
	"github.com/irahardianto/service-pattern-go/viewmodels"

	"testing"

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
	// create an instance of our test object
	playerService := new(mocks.IPlayerService)

	expectedScore := "Forty-Fifteen"
	expectedResult := viewmodels.ScoresVM{}
	expectedResult.Score = expectedScore

	player1Name := "Rafael"
	player2Name := "Serena"

	// setup expectations
	playerService.On(getScores, player1Name, player2Name).Return(expectedScore, nil)

	playerController := PlayerController{playerService}

	// call the code we are testing
	req := testRequest(player1Name, player2Name)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc(routerPattern, playerController.GetPlayerScore)
	r.ServeHTTP(w, req)

	actualResult := viewmodels.ScoresVM{}

	json.NewDecoder(w.Body).Decode(&actualResult)

	// assert that the expectations were met
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}

func TestPlayerScoreNoRecordPlayer1(t *testing.T) {
	// create an instance of our test object
	playerService := new(mocks.IPlayerService)

	player1Name := "fart"
	player2Name := "Rafael"

	expectedResult := ResponseError{}
	expectedResult.Message = recordNotFound

	// setup expectations
	playerService.On(getScores, player1Name, player2Name).Return("", ce.ErrRecordNotFound)

	playerController := PlayerController{playerService}

	// call the code we are testing
	req := testRequest(player1Name, player2Name)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc(routerPattern, playerController.GetPlayerScore)
	r.ServeHTTP(w, req)

	actualResult := ResponseError{}

	json.NewDecoder(w.Body).Decode(&actualResult)

	// assert that the expectations were met
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
}

func TestPlayerScoreNoRecordPlayer2(t *testing.T) {
	// create an instance of our test object
	playerService := new(mocks.IPlayerService)

	player1Name := "Rafael"
	player2Name := "fart"

	expectedResult := ResponseError{}
	expectedResult.Message = recordNotFound

	// setup expectations
	playerService.On(getScores, player1Name, player2Name).Return("", ce.ErrRecordNotFound)

	playerController := PlayerController{playerService}

	// call the code we are testing
	req := testRequest(player1Name, player2Name)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc(routerPattern, playerController.GetPlayerScore)
	r.ServeHTTP(w, req)

	actualResult := ResponseError{}

	json.NewDecoder(w.Body).Decode(&actualResult)

	// assert that the expectations were met
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
}

func TestPlayerScoreNoRecordBothPlayers(t *testing.T) {
	// create an instance of our test object
	playerService := new(mocks.IPlayerService)

	player1Name := "noway"
	player2Name := "fart"

	expectedResult := ResponseError{}
	expectedResult.Message = recordNotFound

	// setup expectations
	playerService.On(getScores, player1Name, player2Name).Return("", ce.ErrRecordNotFound)

	playerController := PlayerController{playerService}

	// call the code we are testing
	req := testRequest(player1Name, player2Name)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc(routerPattern, playerController.GetPlayerScore)
	r.ServeHTTP(w, req)

	actualResult := ResponseError{}

	json.NewDecoder(w.Body).Decode(&actualResult)

	// assert that the expectations were met
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
}

func TestPlayerScoreUnknownError(t *testing.T) {
	// create an instance of our test object
	playerService := new(mocks.IPlayerService)

	player1Name := "noway"
	player2Name := "fart"

	expectedResult := ResponseError{}
	expectedResult.Message = "Unexpected error."

	// setup expectations
	playerService.On(getScores, player1Name, player2Name).Return("", errors.New("Weird error"))

	playerController := PlayerController{playerService}

	// call the code we are testing
	req := testRequest(player1Name, player2Name)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc(routerPattern, playerController.GetPlayerScore)
	r.ServeHTTP(w, req)

	actualResult := ResponseError{}

	json.NewDecoder(w.Body).Decode(&actualResult)

	// assert that the expectations were met
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
}
