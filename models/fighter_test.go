package models

import (
	"testing"
)

func TestNewFighter(t *testing.T) {
	name := "Mario"
	fighter, _ := NewFighter("mario.png")
	if fighter.Name != name || fighter.Alive != true || fighter.TeammateId != 0 || fighter.Killcount != 0 {
		t.Fatal(fighter)
	}
}
