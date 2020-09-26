package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() League {
	_, err := f.database.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
	}
	league, _ := NewLeague(f.database)
	return league
}
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	league := f.GetLeague()
	return league.GetPlayer(name).Wins
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.GetPlayer(name)
	if player == nil {
		league = append(league, Player{Name: name, Wins: 1})
	} else {
		player.Wins++
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}
