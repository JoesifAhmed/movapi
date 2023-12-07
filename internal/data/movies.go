package data

import (
	"time"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	RunTime   int32     `json:"runtime,omitempty,string"`
	Genres    []string  `json:"genres"`
	Version   int32     `json:"version"`
}
