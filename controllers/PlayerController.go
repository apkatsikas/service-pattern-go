package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	ce "github.com/irahardianto/service-pattern-go/customerrors"
	"github.com/irahardianto/service-pattern-go/interfaces"

	"github.com/go-chi/chi"
	"github.com/irahardianto/service-pattern-go/viewmodels"
)

type PlayerController struct {
	interfaces.IPlayerService
}

type ResponseError struct {
	Message string
}

// TODO - how should we handle and test negative case?
func (controller *PlayerController) GetPlayerScore(res http.ResponseWriter, req *http.Request) {

	player1Name := chi.URLParam(req, "player1")
	player2Name := chi.URLParam(req, "player2")

	scores, err := controller.GetScores(player1Name, player2Name)
	if err != nil {
		if err == ce.RecordNotFoundError {
			res.WriteHeader(http.StatusNotFound)
			json.NewEncoder(res).Encode(ResponseError{Message: "Record not found."})
			return
		} else {
			res.WriteHeader(http.StatusInternalServerError)
			// TODO - implement logging here
			errMsg := fmt.Errorf("Got an unexpected error: %v", err)
			fmt.Println(errMsg.Error())
			json.NewEncoder(res).Encode(ResponseError{Message: "Unexpected error."})
		}
	}

	// maybe marshal instead... TODO check that
	json.NewEncoder(res).Encode(viewmodels.ScoresVM{Score: scores})
}
