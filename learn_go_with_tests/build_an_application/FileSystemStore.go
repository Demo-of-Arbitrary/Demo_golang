package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type FileSystemPlayerStore struct {
	database io.Reader
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	var league []Player
	err := json.NewDecoder(f.database).Decode(&league)
	if err != nil {
		fmt.Printf("panic: got error when decode json --> %v \n", err)
	}
	return league
}
