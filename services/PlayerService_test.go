package services

import (
	"testing"

	ce "github.com/irahardianto/service-pattern-go/customerrors"
	"github.com/irahardianto/service-pattern-go/interfaces/mocks"
	"github.com/irahardianto/service-pattern-go/models"

	"github.com/stretchr/testify/assert"
)

const getPlayerByName = "GetPlayerByName"

func TestGetScore(t *testing.T) {
	playerRepository := new(mocks.IPlayerRepository)
	player1Name := "Rafael"
	player2Name := "Serena"
	expectedScore := "Forty-Fifteen"

	player1 := models.Player{}
	player1.Id = 101
	player1.Name = player1Name
	player1.Score = 3

	player2 := models.Player{}
	player2.Id = 103
	player2.Name = player2Name
	player2.Score = 1

	playerRepository.On(getPlayerByName, player1Name).Return(&player1, nil)
	playerRepository.On(getPlayerByName, player2Name).Return(&player2, nil)

	playerService := PlayerService{playerRepository}

	actualResult, err := playerService.GetScores(player1Name, player2Name)

	// Make sure we got the correct score
	assert.Equal(t, expectedScore, actualResult)
	// Make sure we there is no error
	assert.Nil(t, err)
}

func TestGetScoreNoRecordPlayer1(t *testing.T) {
	playerRepository := new(mocks.IPlayerRepository)
	player1Name := "fart"
	player2Name := "Rafael"
	expectedScore := ""

	var player1 models.Player

	player2 := models.Player{}
	player2.Id = 101
	player2.Name = player2Name
	player2.Score = 3

	playerRepository.On(getPlayerByName, player1Name).Return(&player1, nil)
	playerRepository.On(getPlayerByName, player2Name).Return(&player2, ce.ErrRecordNotFound)

	playerService := PlayerService{playerRepository}

	actualScore, err := playerService.GetScores(player1Name, player2Name)

	// Check that we got an empty player score
	assert.Equal(t, expectedScore, actualScore)

	// Check that we got an error
	assert.Equal(t, ce.ErrRecordNotFound, err)
}

func TestGetScoreNoRecordPlayer2(t *testing.T) {
	playerRepository := new(mocks.IPlayerRepository)
	player1Name := "Rafael"
	player2Name := "fart"
	expectedScore := ""

	player1 := models.Player{}
	player1.Id = 101
	player1.Name = player1Name
	player1.Score = 3

	var player2 models.Player

	playerRepository.On(getPlayerByName, player1Name).Return(&player1, nil)
	playerRepository.On(getPlayerByName, player2Name).Return(&player2, ce.ErrRecordNotFound)

	playerService := PlayerService{playerRepository}

	actualScore, err := playerService.GetScores(player1Name, player2Name)

	// Check that we got an empty player score
	assert.Equal(t, expectedScore, actualScore)

	// Check that we got a RecordNotFoundError error
	assert.Equal(t, ce.ErrRecordNotFound, err)
}

func TestGetScoreNoRecordBothPlayers(t *testing.T) {
	player1Name := "farty"
	player2Name := "fart"
	expectedScore := ""

	playerRepository := new(mocks.IPlayerRepository)

	// Zero value
	var player1 models.Player
	var player2 models.Player

	playerRepository.On(getPlayerByName, player1Name).Return(&player1, ce.ErrRecordNotFound)
	playerRepository.On(getPlayerByName, player2Name).Return(&player2, ce.ErrRecordNotFound)

	playerService := PlayerService{playerRepository}

	actualScore, err := playerService.GetScores(player1Name, player2Name)

	// Check that we got an empty player score
	assert.Equal(t, expectedScore, actualScore)

	// Check that we got a RecordNotFoundError error
	assert.Equal(t, ce.ErrRecordNotFound, err)
}
