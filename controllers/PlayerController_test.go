package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	ce "github.com/irahardianto/service-pattern-go/customerrors"
	"github.com/irahardianto/service-pattern-go/interfaces/mocks"
	"github.com/irahardianto/service-pattern-go/viewmodels"

	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

/*
  Actual test functions
*/

// TestSomething is an example of how to use our test object to
// make assertions about some target code we are testing.
func TestPlayerScore(t *testing.T) {

	// create an instance of our test object
	playerService := new(mocks.IPlayerService)

	// setup expectations
	playerService.On("GetScores", "Rafael", "Serena").Return("Forty-Fifteen", nil)

	playerController := PlayerController{playerService}

	// call the code we are testing
	req := httptest.NewRequest("GET", "http://localhost:8080/getScore/Rafael/vs/Serena", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/getScore/{player1}/vs/{player2}", playerController.GetPlayerScore)

	r.ServeHTTP(w, req)

	expectedResult := viewmodels.ScoresVM{}
	expectedResult.Score = "Forty-Fifteen"

	actualResult := viewmodels.ScoresVM{}

	json.NewDecoder(w.Body).Decode(&actualResult)

	// assert that the expectations were met
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}

func TestPlayerScoreNoRecordPlayer1(t *testing.T) {

	// create an instance of our test object
	playerService := new(mocks.IPlayerService)

	// setup expectations
	playerService.On("GetScores", "fart", "Rafael").Return("", ce.RecordNotFoundError)

	playerController := PlayerController{playerService}

	// call the code we are testing
	req := httptest.NewRequest("GET", "http://localhost:8080/getScore/fart/vs/Rafael", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/getScore/{player1}/vs/{player2}", playerController.GetPlayerScore)

	r.ServeHTTP(w, req)

	expectedResult := ResponseError{}
	expectedResult.Message = "Record not found."

	actualResult := ResponseError{}

	json.NewDecoder(w.Body).Decode(&actualResult)

	// assert that the expectations were met
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
}

func TestPlayerScoreNoRecordPlayer2(t *testing.T) {

	// create an instance of our test object
	playerService := new(mocks.IPlayerService)

	// setup expectations
	playerService.On("GetScores", "Rafael", "fart").Return("", ce.RecordNotFoundError)

	playerController := PlayerController{playerService}

	// call the code we are testing
	req := httptest.NewRequest("GET", "http://localhost:8080/getScore/Rafael/vs/fart", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/getScore/{player1}/vs/{player2}", playerController.GetPlayerScore)

	r.ServeHTTP(w, req)

	expectedResult := ResponseError{}
	expectedResult.Message = "Record not found."

	actualResult := ResponseError{}

	json.NewDecoder(w.Body).Decode(&actualResult)

	// assert that the expectations were met
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
}

func TestPlayerScoreUnknownError(t *testing.T) {

	// create an instance of our test object
	playerService := new(mocks.IPlayerService)

	// setup expectations
	playerService.On("GetScores", "Rafael", "fart").Return("", errors.New("Weird error"))

	playerController := PlayerController{playerService}

	// call the code we are testing
	req := httptest.NewRequest("GET", "http://localhost:8080/getScore/Rafael/vs/fart", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/getScore/{player1}/vs/{player2}", playerController.GetPlayerScore)

	r.ServeHTTP(w, req)

	expectedResult := ResponseError{}
	expectedResult.Message = "Unexpected error."

	actualResult := ResponseError{}

	json.NewDecoder(w.Body).Decode(&actualResult)

	// assert that the expectations were met
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
}
