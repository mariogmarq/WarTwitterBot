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
	messages chan []models.FighterApi
	winner   chan []models.FighterApi
	playing  chan byte
	once     = sync.Once{} //used to initialize the rest of vars
)

// Return an array of all players name
func readPlayersName() []string {
	return utils.ReadNamesFromImagesFolder()
}

// Starts the game, returns two channels of arrays of fighterApis, the first one shows the fighters involved in a kill, the second one the winners of the game
// you only can read from the seccond one once the first one has been closed
func StartGame() (chan []models.FighterApi, chan []models.FighterApi) {

	once.Do(func() {
		messages = make(chan []models.FighterApi, 1)
		winner = make(chan []models.FighterApi, 1)
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
	}

	// fighters that interact in the kill
	// the last one is the one who dies
	var fighters []*models.Fighter
	var apis []models.FighterApi
	rand.Shuffle(len(ids), func(i int, j int) {
		ids[i], ids[j] = ids[j], ids[i]
	})

	enemyID := ids[1]

	// Append the one who kills
	fighter, _ := repo.GetFighterById(ids[0])
	fighters = append(fighters, fighter)

	// Append the teammate if he exists
	if fighter.TeammateId != 0 {
		fighter, _ = repo.GetFighterById(fighter.TeammateId)
		fighters = append(fighters, fighter)
		if fighter.ID == enemyID {
			enemyID = ids[2]
		}
	}

	// Append the enemy
	fighter, _ = repo.GetFighterById(enemyID)
	fighters = append(fighters, fighter)

	for _, fighter := range fighters {
		apis = append(apis, fighter.GetApi())
	}

	// TODO: Move to view
	// // All players added, now we get the phrase
	// phrase, _ := repo.GetPhraseByN(len(fighters))
	// message, _ := phrase.MapToPlayers(fighters...)
	// message = message + strconv.Itoa(len(fighters)-1) + " luchadores vivos."
	messages <- apis

	// Kill the player and add 1 to the killcount
	for i := 0; i < len(fighters)-1; i++ {
		repo.AddKillToPlayerByID(fighters[i].ID)
	}
	repo.KillPlayerByID(fighters[len(fighters)-1].ID)

	playing <- '0'
}

// Checks if a fighter has won the game
func haveAWinner(AlivePlayersIDs []uint) bool {
	// Only can win up to 2 players
	if len(AlivePlayersIDs) > 2 {
		return false
	}

	// If there is only one player alive, he/she wins
	if len(AlivePlayersIDs) == 1 {
		return true
	}

	// If there is two players alive, they wins only if they are teammates
	if len(AlivePlayersIDs) == 2 {
		fighter, _ := repo.GetFighterById(AlivePlayersIDs[0])
		if fighter.TeammateId == AlivePlayersIDs[1] {
			return true
		}
	}

	return false
}

// Return the message of victory
func winnerApi(id uint) []models.FighterApi {
	var apis []models.FighterApi

	fighter, err := repo.GetFighterById(id)
	utils.Must(err)
	apis = append(apis, fighter.GetApi())

	if fighter.TeammateId != 0 {
		fighter2, err := repo.GetFighterById(fighter.ID)
		utils.Must(err)
		apis = append(apis, fighter2.GetApi())
	}

	return apis
}

// Send the winner message and return the error to stop the turn
func winnerMessage(id uint) {
	close(messages)
	winner <- winnerApi(id)
	close(winner)
}
