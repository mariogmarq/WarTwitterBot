package models

import (
	"testing"
)

func TestNewFighter(t *testing.T) {
	name := "Mario"
	fighter := NewFighter(name)
	if fighter.Name != name || fighter.Alive != true || fighter.TeammateId != 0 || fighter.Killcount != 0 {
		t.Fatal(fighter)
	}
}
