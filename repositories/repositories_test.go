// Database wrapper
package repositories

import (
	"testing"
)

//TODO: Use SqlMock or other mocking library for this test

// To be run in a docker container
func TestDBIntegration(t *testing.T) {
	repo := openRepositories()
	repo.AddFighter("David", "Jose")
	repo.AddPhrase("$1 kills $2", "$1 kills $1")

	phrases, err := repo.GetPhrasesByN(2)
	if err != nil {
		t.Fatalf("Error %v not expected on GetPhrasesByN(2)", err)
	}

	if len(phrases) != 1 {
		t.Fatalf("Expected 1 phrases, got %d", len(phrases))
	}

	fighter, err := repo.GetFighterById(1)
	if err != nil {
		t.Fatalf("Error %v not expected on GetFighterById(1)", err)
	}

	if fighter.Name != "David" {
		t.Fatalf("Expected name David, got %s", fighter.Name)
	}

	_, err = repo.GetFighterById(0)
	if err == nil {
		t.Fatalf("Expected error got none on GetFighterById(0)")
	}

	ids := repo.AliveFightersIDs()
	if len(ids) != 2 || ids[0] != 1 || ids[1] != 2 {
		t.Fatalf("AliveFightersIDs not working")
	}

	repo.KillPlayerByID(2)
	ids = repo.AliveFightersIDs()
	if len(ids) != 1 || ids[0] != 1 {
		t.Fatalf("KillPlayerByID not working")
	}

	repo.KillPlayerByID(10) //it does not crash

}
