package services

import (
	"testing"

	ce "github.com/irahardianto/service-pattern-go/customerrors"
	"github.com/irahardianto/service-pattern-go/interfaces/mocks"
	"github.com/irahardianto/service-pattern-go/models"

	"github.com/stretchr/testify/assert"
)

func TestGetScore(t *testing.T) {

	playerRepository := new(mocks.IPlayerRepository)

	player1 := models.PlayerModel{}
	player1.Id = 101
	player1.Name = "Rafael"
	player1.Score = 3

	player2 := models.PlayerModel{}
	player2.Id = 103
	player2.Name = "Serena"
	player2.Score = 1

	playerRepository.On("GetPlayerByName", "Rafael").Return(player1, nil)
	playerRepository.On("GetPlayerByName", "Serena").Return(player2, nil)

	playerService := PlayerService{playerRepository}

	expectedResult := "Forty-Fifteen"

	actualResult, _ := playerService.GetScores("Rafael", "Serena")

	assert.Equal(t, expectedResult, actualResult)

	// TODO - add tests for cant find player
}

func TestGetScoreNoRecord(t *testing.T) {
	playerRepository := new(mocks.IPlayerRepository)

	player1 := models.PlayerModel{}
	player1.Id = 101
	player1.Name = "Rafael"
	player1.Score = 3

	player2 := models.PlayerModel{}

	playerRepository.On("GetPlayerByName", "Rafael").Return(player1, nil)
	playerRepository.On("GetPlayerByName", "fart").Return(player2, ce.RecordNotFoundError)

	playerService := PlayerService{playerRepository}

	expectedResult := ""

	actualResult, err := playerService.GetScores("Rafael", "fart")

	assert.Equal(t, expectedResult, actualResult)

	assert.NotNil(t, err)
}
