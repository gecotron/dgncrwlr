package main

import (
	"math/rand"
)

// Simple dungeon structure. Contains player position and layout of rooms + walls
type dungeon struct {
	columns int
	rows    int
	x       int
	y       int
	rooms   [][]rune
}

// Initialise dungeon structure
func initDungeon(c int, r int) dungeon {
	d := dungeon{
		columns: c,
		rows:    r,
		x:       0,
		y:       0,
		rooms:   genRooms(c, r),
	}
	return d
}

/* Dungeon rules
* 1. The Dungeon begins as all walls to avoid issues from random generation
* 2. All empty rooms must have at least one room next to them
* 3. The dungeon has a "start" and "end" point. These must connect in some way
 */

func genRooms(rows int, cols int) [][]rune {
	// Create dungeon representation (slice)

	newRooms := make([][]rune, rows)
	for c := range newRooms {
		newRooms[c] = make([]rune, cols)
	}
	// Fill with walls
	for i := 0; i < len(newRooms); i++ {
		for x := 0; x < len(newRooms[i]); x++ {
			newRooms[i][x] = '#'
		}
	}
	// Carve out rooms

	// Choose start and end point

	var startX int = rand.Intn((rows - 1) - 0 + 1)
	var startY int = rand.Intn((cols - 1) - 0 + 1)
	var endX int = rand.Intn((rows - 1) - 0 + 1)
	var endY int = rand.Intn((cols - 1) - 0 + 1)
	// Make sure start and end are not the same

	for endX == startX {
		endX = rand.Intn((rows - 1) - 0 + 1)
	}
	for endY == startY {
		endY = rand.Intn((cols - 1) - 0 + 1)
	}
	// Set start and end points

	newRooms[startX][startY] = '.'
	newRooms[endX][endY] = '.'
	var X int = startX
	var Y int = startY
	// Begin carving

	switch X > endX {
	case true:
		for X > endX {
			newRooms[X][Y] = '.'
			X--
		}
	case false:
		for X < endX {
			newRooms[X][Y] = '.'
			X++
		}
	}
	switch Y > endY {
	case true:
		for Y > endY {
			newRooms[X][Y] = '.'
			Y--
		}
	case false:
		for Y < endY {
			newRooms[X][Y] = '.'
			Y++
		}
	}
	// Carve Full dungeon
	for i := 0; i < cols; i++ {
		X = rand.Intn((rows - 1) - 0 + 1)
		Y = rand.Intn((cols - 1) - 0 + 1)
		for newRooms[X][Y] == '.' {
			X = rand.Intn((rows - 1) - 0 + 1)
			Y = rand.Intn((cols - 1) - 0 + 1)
		}
		endX = rand.Intn((rows - 1) - 0 + 1)
		endY = rand.Intn((cols - 1) - 0 + 1)
		for newRooms[endX][endY] == '#' {
			endX = rand.Intn((rows - 1) - 0 + 1)
			endY = rand.Intn((cols - 1) - 0 + 1)
		}

		switch X > endX {

		case true:
			for X > endX {
				newRooms[X][Y] = '.'
				X--
			}

		case false:
			for X < endX {
				newRooms[X][Y] = '.'
				X++
			}
		}

		switch Y > endY {

		case true:
			for Y > endY {
				newRooms[X][Y] = '.'
				Y--
			}

		case false:
			for Y < endY {
				newRooms[X][Y] = '.'
				Y++
			}
		}
	}

	// Return created dungeon

	return newRooms
}
