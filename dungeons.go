package main

import "math/rand"

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
	// Create slice
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
	startX := rand.Intn(4 - 0 + 1)
	startY := rand.Intn(4 - 0 + 1)
	endX := rand.Intn(4 - 0 + 1)
	endY := rand.Intn(4 - 0 + 1)
	for endX == startX {
		endX = rand.Intn(4 - 0 + 1)
	}
	for endY == startY {
		endY = rand.Intn(4 - 0 + 1)
	}

	return newRooms
}
