package main

import (
	"testing"
)

func TestNewRover(t *testing.T) {
	rover := NewRovertito(3, 5, East)

	// Check if the rover's initial position and direction are set correctly
	if rover.X != 3 || rover.Y != 5 || rover.Direction != East {
		t.Errorf("Expected initial position (%d, %d) and direction %s, but got (%d, %d) and %s",
			3, 5, East, rover.X, rover.Y, rover.Direction)
	}
}

func TestRoverTurnLeft(t *testing.T) {
	rover := NewRovertito(0, 0, North)
	rover.TurnLeft()

	if rover.Direction != West {
		t.Errorf("Expected direction %s, but got %s", West, rover.Direction)
	}
}

func TestRoverTurnRight(t *testing.T) {
	rover := NewRovertito(0, 0, North)
	rover.TurnRight()

	if rover.Direction != East {
		t.Errorf("Expected direction %s, but got %s", East, rover.Direction)
	}
}

func TestRoverMoveForward(t *testing.T) {
	rover := NewRovertito(0, 0, North)
	rover.MoveForward()

	if rover.Y != 1 {
		t.Errorf("Expected Y coordinate to be 1, but got %d", rover.Y)
	}

	rover.TurnRight()
	rover.MoveForward()

	if rover.X != 1 {
		t.Errorf("Expected X coordinate to be 1, but got %d", rover.X)
	}

	rover.TurnRight()
	rover.MoveForward()

	if rover.Y != 0 {
		t.Errorf("Expected Y coordinate to be 0, but got %d", rover.Y)
	}

	rover.TurnRight()
	rover.MoveForward()

	if rover.X != 0 {
		t.Errorf("Expected X coordinate to be 0, but got %d", rover.X)
	}
}

func TestRoverExecuteCommands(t *testing.T) {
	tests := []struct {
		commands string
		finalX   int
		finalY   int
		finalDir Direction
		err      error
	}{
		{"M", 0, 1, North, nil},
		{"RMMM", 3, 0, East, nil},
		{"LM", -1, 0, West, nil},
		{"MMMMMMMMMM", 0, 10, North, nil},
		{"z", 0, 0, North, ErrInvalidCommand},
	}

	for _, test := range tests {
		rover := NewRovertito(0, 0, North)
		err := rover.ExecuteCommands(test.commands)

		if err != nil && test.err == nil {
			t.Errorf("Unexpected error: %v", err)
		} else if err == nil && test.err != nil {
			t.Errorf("Expected error: %v", test.err)
		} else if err != nil && test.err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected error: %v, but got: %v", test.err, err)
		}

		if rover.X != test.finalX || rover.Y != test.finalY || rover.Direction != test.finalDir {
			t.Errorf("For commands %s, expected position (%d, %d) and direction %s, but got (%d, %d) and %s",
				test.commands, test.finalX, test.finalY, test.finalDir, rover.X, rover.Y, rover.Direction)
		}
	}
}
