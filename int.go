package main

import "math/rand"

// Выдаёт случайное число
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
