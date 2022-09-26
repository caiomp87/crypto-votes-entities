package models

import "time"

type Crypto struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Network   string    `json:"network"`
	UpVotes   int       `json:"upVotes"`
	DownVotes int       `json:"downVotes"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
