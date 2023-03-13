package controllers

import (
	"encoding/json"
	"net/http"

	ce "github.com/irahardianto/service-pattern-go/customerrors"
	"github.com/irahardianto/service-pattern-go/infrastructures/logutil"
	"github.com/irahardianto/service-pattern-go/interfaces"
	"github.com/irahardianto/service-pattern-go/viewmodels"

	"github.com/go-chi/chi"
)

type PlayerController struct {
	interfaces.IPlayerService
}

type ResponseError struct {
	Message string
}

func (controller *PlayerController) GetPlayerScore(res http.ResponseWriter, req *http.Request) {
	player1Name := chi.URLParam(req, "player1")
	player2Name := chi.URLParam(req, "player2")

	// Get the scores from the service
	scores, err := controller.GetScores(player1Name, player2Name)
	if err != nil {
		// Record not found
		if err == ce.ErrRecordNotFound {
			res.WriteHeader(http.StatusNotFound)
			json.NewEncoder(res).Encode(ResponseError{Message: "Record not found."})
			return
			// Everything else
		} else {
			res.WriteHeader(http.StatusInternalServerError)
			logutil.Error("Failed to get scores. Error was: %v", err)
			json.NewEncoder(res).Encode(ResponseError{Message: "Unexpected error."})
			return
		}
	}

	json.NewEncoder(res).Encode(viewmodels.ScoresVM{Score: scores})
}
