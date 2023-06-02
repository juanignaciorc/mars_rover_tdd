package main

import (
	"errors"
	"fmt"
)

var ErrInvalidCommand = errors.New("invalid command")

// Direction la dirección actual en la que está el rovertito
type Direction string

// Constants for the four cardinal directions
const (
	North Direction = "N"
	East  Direction = "E"
	South Direction = "S"
	West  Direction = "W"
)

// Rovertito es un Rovertito que puede moverse en un plano bidimensional
type Rovertito struct {
	X         int
	Y         int
	Direction Direction
}

// NewRovertito creates a new Rovertito with the given initial coordinates and direction
func NewRovertito(x, y int, direction Direction) *Rovertito {
	return &Rovertito{
		X:         x,
		Y:         y,
		Direction: direction,
	}
}

// TurnLeft rotates the rover 90 degrees to the left
func (r *Rovertito) TurnLeft() {
	switch r.Direction {
	case North:
		r.Direction = West
	case East:
		r.Direction = North
	case South:
		r.Direction = East
	case West:
		r.Direction = South
	}
}

// TurnRight rotates the rover 90 degrees to the right
func (r *Rovertito) TurnRight() {
	switch r.Direction {
	case North:
		r.Direction = East
	case East:
		r.Direction = South
	case South:
		r.Direction = West
	case West:
		r.Direction = North
	}
}

// MoveForward moves the rovertito one unit forward in its current direction
func (r *Rovertito) MoveForward() {
	switch r.Direction {
	case North:
		r.Y++
	case East:
		r.X++
	case South:
		r.Y--
	case West:
		r.X--
	}
}

// ExecuteCommands executes commands on the rovertito
func (r *Rovertito) ExecuteCommands(commands string) error {
	for _, command := range commands {
		switch command {
		case 'L':
			r.TurnLeft()
		case 'R':
			r.TurnRight()
		case 'M':
			r.MoveForward()
		default:
			return ErrInvalidCommand
		}
	}
	return nil
}

// main function for testing el Rovertito de Marte
func main() {
	rover := NewRovertito(0, 0, North)

	// Execute commands and check the final position and direction
	err := rover.ExecuteCommands("MMRMMRMRRM")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Final Position: (%d, %d), Direction: %s\n", rover.X, rover.Y, rover.Direction)
	}
}
