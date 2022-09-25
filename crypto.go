package main

type Crypto struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Network   string `json:"network"`
	UpVotes   int    `json:"upVotes"`
	DownVotes int    `json:"downVotes"`
}
