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
