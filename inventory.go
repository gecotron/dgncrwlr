package main

type inventory struct {
	gold  int             // How much gold the player has. Used to purchase items
	exp   int             // How much experience the player has gained
	items map[string]item // Items the player has on hand
}

type item struct {
	name  string // Item name. Used when adding to inventory item map
	form  rune   // Type of item. Wand, Weapon, Ring, Staff
	value int    // Value to buy/sell
}
