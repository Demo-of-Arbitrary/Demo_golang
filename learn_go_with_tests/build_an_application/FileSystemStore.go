package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	file.Seek(0, 0)
	info, err := file.Stat()

	if err != nil {
		return nil, fmt.Errorf("problem getting file info form file: %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{json.NewEncoder(&tape{file}), league}, nil
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	league := f.GetLeague()
	player := league.GetPlayer(name)
	if player == nil {
		return 0
	}
	return player.Wins
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.GetPlayer(name)
	if player == nil {
		f.league = append(f.league, Player{Name: name, Wins: 1})
	} else {
		player.Wins++
	}

	f.database.Encode(f.league)
}
