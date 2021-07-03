// Implements the whole logic of the game
package warlogic

import (
	"github/mariogmarq/WarTwitterBot/internal/models"
	"github/mariogmarq/WarTwitterBot/internal/repository"
	"github/mariogmarq/WarTwitterBot/utils"
	"log"
	"math/rand"
	"sync"
)

var (
	repo     repository.IRepository
	messages chan []*models.FighterApi
	winner   chan *models.FighterApi
	playing  chan byte
	once     = sync.Once{} //used to initialize the rest of vars
)

// Return an array of all players name
func readPlayersName() []string {
	return utils.ReadNamesFromImagesFolder()
}

// Starts the game, returns two channels of arrays of fighterApis, the first one shows the fighters involved in a kill, the second one the winners of the game
// you only can read from the seccond one once the first one has been closed
func StartGame() (chan []*models.FighterApi, chan *models.FighterApi) {

	once.Do(func() {
		messages = make(chan []*models.FighterApi, 1)
		winner = make(chan *models.FighterApi, 1)
		playing = make(chan byte, 1) // If its opened then the game is not finished
		repo = repository.GetInstance()

		go func() {
			repo.AddFighter(readPlayersName()...)
			repo.AddPhrase("$1 kills $2")

			turn()
			for range playing {
				turn()
			}

		}()

	})

	return messages, winner
}

// Make a player to kill other advancing a turn
func turn() {
	log.Println("Running turn")
	ids := repo.AliveFightersIDs()

	// Checks if the game has ended
	if haveAWinner(ids) {
		winnerMessage(ids[0])
		close(playing)
		return
	}

	// fighters that interact in the kill
	// the last one is the one who dies
	var fighters []*models.FighterApi
	rand.Shuffle(len(ids), func(i int, j int) {
		ids[i], ids[j] = ids[j], ids[i]
	})

	enemyID := ids[1]

	// Append the one who kills
	fighter, _ := repo.GetFighterAPIById(ids[0])
	fighters = append(fighters, fighter)

	// Append the enemy
	fighter, _ = repo.GetFighterAPIById(enemyID)
	fighters = append(fighters, fighter)

	messages <- fighters

	// Kill the player and add 1 to the killcount
	repo.AddKillToPlayerByID(ids[0])
	repo.KillPlayerByID(ids[1])

	playing <- '0'
}

// Checks if a fighter has won the game
func haveAWinner(AlivePlayersIDs []uint) bool {
	return len(AlivePlayersIDs) == 1
}

// Return the message of victory
func winnerApi(id uint) *models.FighterApi {
	rv, _ := repository.GetInstance().GetFighterAPIById(id)
	return rv
}

// Send the winner message and return the error to stop the turn
func winnerMessage(id uint) {
	close(messages)
	winner <- winnerApi(id)
	close(winner)
}
