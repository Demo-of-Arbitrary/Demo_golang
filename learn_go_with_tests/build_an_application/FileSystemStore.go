package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
	league   League
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, 0)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{database, league}
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	league := f.GetLeague()
	fmt.Println("get league: ", league)
	player := league.GetPlayer(name)
	if player == nil {
		return 0
	}
	return player.Wins
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.GetPlayer(name)
	fmt.Println("output: ", player)
	if player == nil {
		f.league = append(f.league, Player{Name: name, Wins: 1})
	} else {
		player.Wins++
	}
	fmt.Println("league: ", f.league)

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(f.league)
}
