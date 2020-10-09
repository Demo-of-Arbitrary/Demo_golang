package poker

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]
			`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.GetPlayerScore("Chris")
		want := 33
		assertScoreEqual(t, got, want)
	})

	t.Run("store wins for existing player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]
			`)
		defer cleanDatabase()

		store, _ := NewFileSystemPlayerStore(database)

		store.RecordWin("Chris")

		got := store.GetPlayerScore(("Chris"))
		want := 34

		assertScoreEqual(t, got, want)
	})

	t.Run("store wins for new player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]
			`)
		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		store.RecordWin("French")

		got := store.GetPlayerScore(("French"))
		want := 1

		assertScoreEqual(t, got, want)
	})

	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
		got := store.GetLeague()
		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}
		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})
}

func assertScoreEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect and error but got one, %v", err)
	}
}

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()
	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file: %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}
	return tmpfile, removeFile
}
