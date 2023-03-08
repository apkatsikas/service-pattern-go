package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/irahardianto/service-pattern-go/interfaces"

	"github.com/go-chi/chi"
	"github.com/irahardianto/service-pattern-go/viewmodels"
)

type PlayerController struct {
	interfaces.IPlayerService
}

// TODO - how should we handle and test negative case?
func (controller *PlayerController) GetPlayerScore(res http.ResponseWriter, req *http.Request) {

	player1Name := chi.URLParam(req, "player1")
	player2Name := chi.URLParam(req, "player2")

	scores, err := controller.GetScores(player1Name, player2Name)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(res).Encode(viewmodels.ScoresVM{Score: scores})
}
