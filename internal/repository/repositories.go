// Database wrapper
package repository

import (
	"github/mariogmarq/WarTwitterBot/internal/models"
	"os"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Implements IRepository
type Repository struct {
	db *gorm.DB
}

//Repository is a singleton
var singleton *Repository
var once = sync.Once{}

func GetInstance() *Repository {
	once.Do(func() {
		singleton = openRepositories()
	})

	return singleton
}

// Open the database, makes migrations and return the database
func openRepositories() *Repository {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_URL")))

	if err != nil {
		panic(err.Error())
	}

	if err = migrations(db); err != nil {
		panic(err.Error())
	}

	return &Repository{db: db}
}

//Migrate all models to database
func migrations(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Fighter{}, &models.Phrase{}); err != nil {
		return err
	}

	return nil
}

// Create a fighters from his names and introduces them into the database
func (r *Repository) AddFighter(names ...string) {
	var fighters []*models.Fighter
	for _, name := range names {
		fighter, err := models.NewFighter(name)
		if err == nil {
			fighters = append(fighters, fighter)
		}
	}

	r.db.Create(fighters)
}

// Create phrases from his sentences and introduces them into the database
// if a phrase is not valid it will be skipped
func (r *Repository) AddPhrase(sentences ...string) {
	var phrases []*models.Phrase
	for _, sentence := range sentences {
		phrase, err := models.NewPhrase(sentence)
		if err != nil {
			continue
		}

		phrases = append(phrases, phrase)
	}

	r.db.Create(phrases)
}

// Retrieves a fighter by his id, if there is no user it returns error
func (r *Repository) GetFighterById(id uint) (*models.Fighter, error) {
	fighter := new(models.Fighter)
	result := r.db.First(fighter, id)

	return fighter, result.Error
}

// Returns all the phrases with a given size
func (r *Repository) GetPhrasesByN(N int) ([]models.Phrase, error) {
	var phrases []models.Phrase

	result := r.db.Where(&models.Phrase{N: N}, "N").Find(&phrases)

	return phrases, result.Error
}

// Return a random phrase of the given size, only one
func (r *Repository) GetPhraseByN(N int) (models.Phrase, error) {
	var phrase models.Phrase

	result := r.db.Where(&models.Phrase{N: N}, "N").Take(&phrase)

	return phrase, result.Error
}

// Return an array of the fighters that are still alive
func (r *Repository) AliveFightersIDs() []uint {
	var ids []uint
	r.db.Model(&models.Fighter{}).Where(&models.Fighter{Alive: true}, "Alive").Select("id").Find(&ids)

	return ids
}

// Update a player with a given id to kill him
func (r *Repository) KillPlayerByID(id uint) {
	r.db.Model(&models.Fighter{}).Where("id = ?", id).Update("Alive", false)
	r.db.Model(&models.Fighter{}).Where("teammate_id = ?", id).Update("teammate_id", 0)
}

//Adds one to player killcount
func (r *Repository) AddKillToPlayerByID(id uint) {
	//retrieve previous killcount
	var kills int
	result := r.db.Model(&models.Fighter{}).Where("id = ?", id).Select("killcount").First(&kills)
	if result.Error != nil {
		return
	}

	r.db.Model(&models.Fighter{}).Where("id = ?", id).Update("killcount", kills+1)
}
