package services

import (
	"github.com/irahardianto/service-pattern-go/interfaces"
)

type PlayerService struct {
	interfaces.IPlayerRepository
}

// Computes and returns score as a string, from the repository player's number values
func (ps *PlayerService) GetScores(player1Name string, player2Name string) (string, error) {
	baseScore := [4]string{"Love", "Fifteen", "Thirty", "Forty"}
	var result string

	// Get players
	player1, err := ps.GetPlayerByName(player1Name)
	if err != nil {
		return "", err
	}

	player2, err := ps.GetPlayerByName(player2Name)
	if err != nil {
		return "", err
	}

	// Computer string score from players number scores
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
