package services

import (
	ce "github.com/irahardianto/service-pattern-go/customerrors"
	"github.com/irahardianto/service-pattern-go/interfaces"
)

type PlayerService struct {
	interfaces.IPlayerRepository
}

func (service *PlayerService) GetScores(player1Name string, player2Name string) (string, error) {

	baseScore := [4]string{"Love", "Fifteen", "Thirty", "Forty"}
	var result string

	player1, err := service.GetPlayerByName(player1Name)
	if err != nil {
		return "", err
	}

	player2, err := service.GetPlayerByName(player2Name)
	if err != nil {
		return "", err
	}

	// TODO - maybe we should raise an error here as opposed to in repository
	if player1.Name == "" || player2.Name == "" {
		return "", ce.RecordNotFoundError
	}

	if player1.Score < 4 && player2.Score < 4 && !(player1.Score+player2.Score == 6) {

		s := baseScore[player1.Score]

		if player1.Score == player2.Score {
			result = s + "-All"
		} else {
			result = s + "-" + baseScore[player2.Score]
		}
	}

	if player1.Score == player2.Score {
		result = "Deuce"
	}

	return result, nil
}
