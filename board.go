/**
* An implementation of a program for playing go, written in go
* Ron Mackenzie
*/
package main

import "fmt"

type Board struct{
	board [][]int
}

func newBoard() *Board {
	f := new(Board)
	//top level slice of the board
	f.board = make([][]int, 19)
	//make the individual slices of the board
	for i := range f.board {
		f.board[i] = make([]int, 19)
	}
	return f
}

func (f Board) printBoard() {
	for i := range f.board {
		for j := range f.board[i] {
			fmt.Printf("[")
			fmt.Printf(string(f.board[i][j]))
			fmt.Printf("]")
		}
		fmt.Printf("\n")
	}
}

/**
*	Represents a move on the board
*	Check if the move is valid, add the move for appropriate player (p)
*	Then check to see if other positions affected (captures, ko)
*	TODO represent p as object of type Player later
*	Returns true if move is successful, false if invalid move
*/
func (f Board) move(x int, y int, p int) (bool) {
	if !f.validMove(x,y,p) {
		return false
	} 

	f.board[x-1][y-1] = p
	//check liberties
	//check for ko
	
	return true
}

/**
* Checks the liberties of a stone, including those of any group it is part of
*/
func (f Board) checkLibs(x int, y int, p int) int {
	var cnt = 0
	if f.board[x-1][y] == 0 {
		cnt++
	}
	//TODO: IMPLEMENTATION
	return cnt
}

/**
* Checks if a given move is valid
*/
func (f Board) validMove(x int, y int, p int) bool {
	//TODO: IMPLEMENTATION
	return false
}

func main() {
	board := newBoard()
	board.printBoard()
}