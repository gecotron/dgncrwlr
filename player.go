package main

import (
	"os"

	toml "github.com/pelletier/go-toml/v2"
)

type player struct {
	Name  string    // Player name
	Level int       // Player level, determines abilities
	Class string    // What kind of adventurer the player is. Fighter, Mage, Cleric, Theif.
	Inv   inventory // Players total items, gold and experience
	X     int       // player x coordinate
	Y     int       // player y coordinate
}

func initPlayer(v rune, name string, x int, y int) player {
	p := player{}
	switch v {
	case 'l':
		loadUser(name)
	default:
		p.Name = name
		p.Class = "fighter"
		p.Level = 1
		p.X = x / 2
		p.Y = y / 2
	}
	return p
}

func saveUser(user player, fileName string) {
	fileName = getExtension(fileName)
	file, err := toml.Marshal(&user)
	errCheck(err)
	err = os.WriteFile(fileName, file, 0666)
	errCheck(err)
}

func loadUser(fileName string) player {
	fileName = getExtension(fileName)
	var user player
	file, err := os.ReadFile(fileName)
	errCheck(err)
	err = toml.Unmarshal(file, &user)
	errCheck(err)
	return user
}

func getExtension(fileName string) string {
	return "users/" + fileName + ".toml"
}
