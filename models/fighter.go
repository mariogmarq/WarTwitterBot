package models

import (
	"github/mariogmarq/WarTwitterBot/utils"

	"gorm.io/gorm"
)

// This model represent a fighter in the game, each fighter has his name, a picture to themselves, an status if he is still
// alive, a pointer that may be null to another user indicating that they are a team, and a kill
// count
type Fighter struct {
	gorm.Model
	Name       string
	Picture    string
	Alive      bool
	TeammateId uint
	Killcount  uint
}

type FighterApi struct {
	Name    string
	Picture string
}

// NewFighter creates a fighter with the default values that are expected and follows the
// conventions to the pictures
func NewFighter(picture string) (*Fighter, error) {
	name, err := utils.ParseImageName(picture)
	if err != nil {
		return nil, err
	}
	return &Fighter{
		Name:       name,
		Picture:    picture,
		Alive:      true,
		TeammateId: 0, //Gorm first index is 1
		Killcount:  0,
	}, nil
}

func (f *Fighter) GetApi() FighterApi {
	return FighterApi{
		Name:    f.Name,
		Picture: f.Picture,
	}
}
