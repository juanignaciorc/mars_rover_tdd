# Mars Rover TDD Practice

This project is a TDD (Test-Driven Development) practice implementation of the Mars Rover problem. The objective of this project is to demonstrate the TDD approach by implementing the Mars Rover functionality using the Go programming language.

## Problem Description

The Mars Rover is a fictional vehicle that moves on a rectangular grid. The grid is divided into cells, and each cell represents a coordinate (x, y) on the grid. The Rover can receive a sequence of commands to move and change its direction on the grid.

The commands the Rover can receive are as follows:
- **M**: Move forward in the current direction.
- **L**: Turn 90 degrees to the left (without moving).
- **R**: Turn 90 degrees to the right (without moving).

The Rover is always initialized with an initial position (x, y) and a direction it is facing (North, South, East, or West).

## Project Structure

The project consists of the following files:

- `main.go`: Contains the main function to execute and test the Mars Rover functionality.
- `rover.go`: Defines the `Rover` struct and its methods to execute commands and move on the grid.
- `rover_test.go`: Contains the unit tests for the `Rover` methods using the TDD approach.

## Running the Tests

To run the tests for this project, follow these steps:

1. Install Go on your machine if you haven't already.
2. Open a terminal and navigate to the project directory.
3. Run the command `go test` to execute all the unit tests.

The tests will validate the correctness of the `Rover` methods by checking the final position and direction after executing a sequence of commands.

## TDD Approach

Test-Driven Development (TDD) is a software development technique where tests are written before writing the implementation code. The TDD approach follows these steps:

1. Write a failing test: Start by writing a test that defines the expected behavior or functionality.
2. Run the test: Execute the test to ensure it fails, indicating that the functionality is not yet implemented.
3. Write the implementation code: Write the minimal code required to make the failing test pass.
4. Run the test again: Execute the test to verify that it passes after the implementation code is written.
5. Refactor the code (if necessary): Improve the code structure without changing the functionality.
6. Repeat the process: Continue writing new tests and implementing the corresponding functionality, one test at a time.

By following the TDD approach, we can ensure that the code is thoroughly tested, maintainable, and satisfies the specified requirements.

## Conclusion

This project serves as a TDD practice implementation of the Mars Rover problem. By following the TDD approach, the code is developed incrementally, with each functionality being thoroughly tested. This project demonstrates the benefits of TDD in terms of code quality, maintainability, and test coverage.
