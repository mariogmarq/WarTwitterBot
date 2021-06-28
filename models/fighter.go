package models

import "gorm.io/gorm"

// This model represent a fighter in the game, each fighter has his name, a picture to themselves, an status if he is still
// alive, a pointer that may be null to another user indicating that they are a team, and a kill
// count
type Fighter struct {
	gorm.Model
	Name       string
	Alive      bool
	TeammateId uint
	Killcount  uint
}

// NewFighter creates a fighter with the default values that are expected and follows the
// conventions to the pictures(disabled for now)
func NewFighter(name string) *Fighter {
	return &Fighter{
		Name:       name,
		Alive:      true,
		TeammateId: 0, //Gorm first index is 1
		Killcount:  0,
	}
}
