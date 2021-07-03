package repository

import "github/mariogmarq/WarTwitterBot/internal/models"

// Implements all database operations
// Is a singleton
type IRepository interface {
	AddFighter(names ...string)
	AddPhrase(sentences ...string)
	GetFighterById(id uint) (*models.Fighter, error)
	GetPhrasesByN(N int) ([]models.Phrase, error)
	GetPhraseByN(N int) (models.Phrase, error)
	AliveFightersIDs() []uint
	KillPlayerByID(id uint)
	AddKillToPlayerByID(id uint)
	GetFighterAPIById(id uint) (*models.FighterApi, error)
}
