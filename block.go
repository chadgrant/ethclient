package main

//DTO data structure to serialize for json
type Block struct {
	Block        string   `json:Block`
	Transactions []string `json:Transactions`
}
