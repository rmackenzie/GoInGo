/**
* An implementation of a program for playing go, written in go
* Ron Mackenzie
*/
package main

import "fmt"

/**
* ///////////////////////////////////////////////////// START BOARD METHODS /////////////////////////////////////////
*/
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
			fmt.Print(f.board[i][j])
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
func (f Board) move(pair Pair, p int) (bool) {
	if !f.validMove(pair,p) {
		return false
	} 

	f.board[pair.x-1][pair.y-1] = p
	//check liberties
	//check for ko
	
	return true
}

/**
* Checks the liberties of a stone, including those of any group it is part of
*/
func (f Board) checkLibs(pair Pair, p int) int {

	stack := newStack()
	checkedPairs := newArrayList()
	cnt := 0
	//TODO: IMPLEMENTATION
	// first push every terminal stone into a stack
	// then sum the liberties of all those stones

	stack.push(newPair(pair.x,pair.y))
	for !stack.isEmpty() {
		checkPair := stack.pop()
		if !checkedPairs.member(*(checkPair)) {
			if f.board[checkPair.x-1][checkPair.y] == p {
				stack.push(newPair(checkPair.x-1, checkPair.y))
			} else if f.board[checkPair.x-1][checkPair.y] == 0 {
				cnt++
			}
			if f.board[checkPair.x][checkPair.y-1] == p {
				stack.push(newPair(checkPair.x, checkPair.y-1))
			} else if f.board[checkPair.x][checkPair.y-1] == 0 {
				cnt++
			}
			if f.board[checkPair.x+1][checkPair.y] == p {
				stack.push(newPair(checkPair.x+1, checkPair.y))
			} else if f.board[checkPair.x+1][checkPair.y] == 0 {
				cnt++
			}
			if f.board[checkPair.x][checkPair.y+1] == p {
				stack.push(newPair(checkPair.x, checkPair.y+1))
			} else if f.board[checkPair.x][checkPair.y+1] == 0 {
				cnt++
			}
		}
	}
	return cnt
}

/**
* Checks if a given move is valid
*/
func (f Board) validMove(pair Pair, p int) bool {
	//TODO: IMPLEMENTATION
	return false
}

/**
* Will use stacks for counting liberties
* /////////////////////////////////////////////////////// START STACK METHODS //////////////////////////////////////////////////
*/
type Stack struct{
	array []*Pair
	last int
}

func newStack() *Stack {
	f := new(Stack)
	f.array = make([]*Pair,5)
	f.last = 0
	return f
}

func (f Stack) isEmpty() bool {
	return f.last==0
}

func (f *Stack) push(pair *Pair) bool {
	if len(f.array) == cap(f.array) { //need to expand the stack
		var newArray = make([]*Pair, len(f.array), (cap(f.array)+1)*2)
		copy(newArray, f.array)
		f.array = newArray
	}
	f.last++
	f.array[f.last] = pair
	return true
}

func (f *Stack) pop() (*Pair) {
	if f.last == 0 {
		return nil
	}
	tmp := f.array[f.last]
	f.array[f.last] = nil
	f.last--
	return tmp
}

/**
* Represent pairs on the board, used instead of passing x,y references to every method
* //////////////////////////////////////////////////////// START PAIR METHODS /////////////////////////////////////////////////
*/
type Pair struct{
	x int
	y int
}

func newPair(x int, y int) *Pair {
	f := new(Pair)
	f.x = x
	f.y = y
	return f
}

func (f *Pair) Print() {
	fmt.Printf("%d %d\n", f.x, f.y)
}

func (f Pair) equals(pair Pair) bool {
	return (f.x == pair.x && f.y == pair.y)
}

/**
* Use lists for keeping track of things counted
* /////////////////////////////////////////////////////// START ARRAYLIST METHODS ///////////////////////////////////////////////
*/

type ArrayList struct{
	array []Pair
	last int
}

func newArrayList() *ArrayList {
	f := new(ArrayList)
	f.array = make([]Pair, 10)
	f.last = 0
	return f
}

func (f ArrayList) member(pair Pair) bool {
	for i := range f.array {
		if pair.equals(f.array[i]) {
			return true
		}
	}
	return false
}

func (f *ArrayList) add(pair Pair) {
	if f.last == cap(f.array) {
		newArray := make([]Pair, len(f.array), (cap(f.array)+1)*2)
		copy(newArray, f.array)
		f.array = newArray
	}
	f.last++
	f.array[f.last] = pair
}

func main() {
	board := newBoard()
	board.board[2][2] = 1
	board.printBoard()
	stone := newPair(2,2)	
	fmt.Println(board.checkLibs(*stone, 1))
}