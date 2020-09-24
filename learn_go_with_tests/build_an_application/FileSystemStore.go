package main

import (
	"fmt"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	_, err := f.database.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
	}
	league, _ := NewLeague(f.database)
	return league
}
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	players := f.GetLeague()
	for _, player := range players {
		if player.Name == name {
			return player.Wins
		}
	}
	return 0
}
