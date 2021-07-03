package models

import (
	"errors"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

// Represent a kill phrase between N-1 fighter to another fighter
// is structure should be
// "$1 here goes the announcement of the kill to player $2"
type Phrase struct {
	gorm.Model
	// Phrase contains the special character $X that will be mapped to a fighter
	Phrase string
	// Number of differents $X in the phrase
	N int
}

// Maps a phrase to two players
// return error in case the map has not been correct (not all players could be
// mapped or there is room for more players
func (p *Phrase) MapToPlayers(fighters ...*FighterApi) (string, error) {
	if p.N != len(fighters) {
		return "", errors.New("phrase.N and the number of arguments passed does not match")
	}

	finalPhrase := p.Phrase

	for i := 1; i <= len(fighters); i++ {
		finalPhrase = strings.ReplaceAll(finalPhrase, "$"+strconv.Itoa(i), fighters[i-1].Name)
	}

	return finalPhrase, nil

}

//Creates a new phrase
//Return error in case of N is less than 2
func NewPhrase(phrase string) (*Phrase, error) {
	//Need to filter how many different $X are
	numberOfDollars := strings.Count(phrase, "$")

	countedDollars := 0
	i := 1
	for ; ; i++ {
		dollarsAti := strings.Count(phrase, "$"+strconv.Itoa(i))
		if dollarsAti == 0 {
			break
		}

		countedDollars += dollarsAti
	}

	if countedDollars != numberOfDollars {
		return &Phrase{}, errors.New("$X does not start at 1 or it skipped a natural number")
	}

	if i <= 2 {
		return &Phrase{}, errors.New("not enough different $X, at least 2")
	}

	return &Phrase{Phrase: phrase, N: i - 1}, nil
}
