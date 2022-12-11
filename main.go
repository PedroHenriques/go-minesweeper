package main

import (
	_ "embed"
	"encoding/json"
	"log"

	"github.com/pedrohenriques/go-minesweeper/internal/configs"
	"github.com/pedrohenriques/go-minesweeper/internal/gui"
)

//go:embed configs/main.json
var configFileData []byte

/*
main is the entry point into the application.
*/
func main() {
	config := &configs.Configs{}
	err := json.Unmarshal(configFileData, config)
	if err != nil {
		log.Fatal(err)
	}

	gui.Run(config)
}
