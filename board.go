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
	maxX int
	maxY int
	moveList *ArrayList
}

func newBoard() *Board {
	f := new(Board)
	//top level slice of the board
	f.board = make([][]int, 19)
	//make the individual slices of the board
	for i := range f.board {
		f.board[i] = make([]int, 19)
	}
	f.moveList = newArrayList()
	f.maxX = 18
	f.maxY = 18
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

	f.board[pair.x][pair.y] = p
	
	//TODO: IMPLEMENT ANY EFFECTS ON CAPTURES
	return true
}

/**
* Checks the liberties of a stone, including those of any group it is part of
* Starting with a stone, push into the stack and then check all adjacent points
* For every connected stone, push those stones into the stack
* Until the stack is empty, check the adjacent spaces for every stone, adding each liberty to a list
* Return the length of the list
*/
func (f Board) checkLibs(pair Pair, p int) int {

	//if the pair to check is out of the bounds of the board
	if pair.x < 0 || pair.x > f.maxX || pair.y < 0 || pair.y > f.maxY {
		return -1
	}

	stack := newStack()
	checkedPairs := newArrayList()
	liberties := newArrayList()

	stack.push(newPair(pair.x,pair.y))
	for !stack.isEmpty() {
		checkPair := stack.pop()
		if !checkedPairs.member(*(checkPair)) {

			//check to the left
			if checkPair.x-1 >= 0 {
				if f.board[checkPair.x-1][checkPair.y] == p {
					stack.push(newPair(checkPair.x-1, checkPair.y))
				} else if f.board[checkPair.x-1][checkPair.y] == 0 {
					lib := newPair(checkPair.x-1,checkPair.y)
					if !liberties.member(*lib) {
						liberties.add(*lib)
					}
				}
			}

			//check above
			if checkPair.y-1 >= 0 {
				if f.board[checkPair.x][checkPair.y-1] == p {
					stack.push(newPair(checkPair.x, checkPair.y-1))
				} else if f.board[checkPair.x][checkPair.y-1] == 0 {
					lib := newPair(checkPair.x,checkPair.y-1)
					if !liberties.member(*lib) {
						liberties.add(*lib)
					}				
				}
			}

			//check to the right
			if checkPair.x+1 <= f.maxX {
				if f.board[checkPair.x+1][checkPair.y] == p {
					stack.push(newPair(checkPair.x+1, checkPair.y))
				} else if f.board[checkPair.x+1][checkPair.y] == 0 {
					lib := newPair(checkPair.x+1,checkPair.y)
					if !liberties.member(*lib) {
						liberties.add(*lib)
					}				
				}
			}

			if checkPair.y+1 <= f.maxY {
				if f.board[checkPair.x][checkPair.y+1] == p {
					stack.push(newPair(checkPair.x, checkPair.y+1))
				} else if f.board[checkPair.x][checkPair.y+1] == 0 {
					lib := newPair(checkPair.x,checkPair.y+1)
					if !liberties.member(*lib) {
						liberties.add(*lib)
					}			
				}
			}
		}
		checkedPairs.add(*checkPair)
	}
	return liberties.length()
}

/**
* Checks if a given move is valid
*/
func (f Board) validMove(pair Pair, p int) bool {
	fmt.Println("in valid move")
	//In go, a move is valid unless:
	//You place a stone where it would have no liberties (be immediately captured) without creating a liberty
	//You place a stone that returns the board exactly to its previous position

	//within board boundaries
	if pair.x < 0 || pair.x > f.maxX || pair.y < 0 || pair.y > f.maxY {
		fmt.Println("out of boundaries")
		return false
	}	

	//on top of another stone
	if f.board[pair.x][pair.y] != 0 {
		fmt.Println("on top of stone")
		return false
	}		

	//check to see if would have 0 liberties
	testBoard := f.copy()
	testBoard.board[pair.x][pair.y] = p
	if !(testBoard.checkLibs(pair, p) > 0) {
		//check to see if it will create a liberty
		//check if any of the adjacent stones have only 1 liberty
		//since we already checked that this lib is open, they necessarily have only this liberty
		//if it does make a liberty, we still must check to see if is a ko situation
		makesLib := 0

		//left
		if pair.x-1 > 0 {
			fmt.Println("check left")
			checkLeft := newPair(pair.x-1,pair.y)
			if f.checkLibs(*checkLeft, f.board[checkLeft.x][checkLeft.y]) == 1 {
				makesLib++
			}
		}
		//above
		if pair.y-1 > 0 {
			fmt.Println("check above")
			checkAbove := newPair(pair.x,pair.y-1)
			if f.checkLibs(*checkAbove, f.board[checkAbove.x][checkAbove.y]) == 1 {
				makesLib++
			}
		}	

		//right
		if pair.x+1 < f.maxX {
			fmt.Println("check right")
			checkRight := newPair(pair.x+1,pair.y)
			if f.checkLibs(*checkRight, f.board[checkRight.x][checkRight.y]) == 1 {
				makesLib++
			}
		}

		//below
		if pair.y+1 < f.maxY {
			fmt.Println("check below")
			checkBelow := newPair(pair.x,pair.y+1)
			if f.checkLibs(*checkBelow, f.board[checkBelow.x][checkBelow.y]) == 1 {
				makesLib++
			}
		}

		//do we need to check if ko or not?
		//just check to see if 
		if makesLib > 0 {

			//can never have ko if you make more than one lib
			if makesLib > 1 {
				fmt.Println("makes more than one lib")
				return true
			}

			//cannot play a move that brings the board back to the previous position
			testBoard = f.copy()

			//what will the board look like after playing there
			testBoard.board[pair.x][pair.y] = p

			//does match the board's position after your last turn
			f.board[f.moveList.Last().x][f.moveList.Last().y] = 0
			if f.equals(testBoard) {
				fmt.Println("breaks ko rule")
				return false
			}
		}
		fmt.Println("doesn't make any liberties, wouldn't have any")
		return false
	}
	fmt.Println("has liberties, valid move")
	return true 
}

func (f Board) equals(b Board) bool {
	for i := range f.board {
		for j := range f.board[i] {
			if !(f.board[i][j] == b.board[i][j]) {
				return false	
			}
		}
	}
	return true
}

func (f Board) copy() Board {
	b := newBoard()
	for i := range f.board {
		for j := range f.board[i] {
			b.board[i][j] = f.board[i][j]
		}
	}
	b.moveList = f.moveList
	b.maxX = f.maxX
	b.maxY = f.maxY	
	return *b
}

/**
* Move struct is a pair + a player
* ///////////////////////////////////////////////////////// START MOVE METHODS /////////////////////////////////////////////////
*/

type Move struct{
	pair Pair
	p int
}

func (f Move) equals(move Move) bool {
	return (f.pair.equals(move.pair) && f.p == move.p)
}

func (f Move) Print() {
	fmt.Print(f.pair)
	fmt.Print(" , Player: ")
	fmt.Println(f.p)
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
	f.last = -1
	return f
}

func (f Stack) isEmpty() bool {
	return f.last== -1
}

func (f *Stack) push(pair *Pair) bool {
	if f.last+1 >= cap(f.array) { //need to expand the stack
		var newArray = make([]*Pair, len(f.array), (cap(f.array)+1)*2)
		copy(newArray, f.array)
		f.array = newArray
	}
	f.last++
	f.array[f.last] = pair
	return true
}

func (f *Stack) pop() (*Pair) {
	if f.last == -1 {
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
* Would be nice to use generics but I don't really get how to implement the same functionality in go
* /////////////////////////////////////////////////////// START ARRAYLIST METHODS ///////////////////////////////////////////////
*/

type ArrayList struct{
	array []Pair
	last int
}

func newArrayList() *ArrayList {
	f := new(ArrayList)
	f.array = make([]Pair, 0, 10)
	f.last = -1
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
	if f.last+1 >= cap(f.array) {
		newArray := make([]Pair, len(f.array), (cap(f.array)+1)*2)
		copy(newArray, f.array)
		f.array = newArray
	}
	f.last++
	f.array = f.array[:f.last+1]
	f.array[f.last] = pair
}

func (f ArrayList) length() int {
	return len(f.array)
}

func (f ArrayList) Print() {
	for i := range f.array {
		f.array[i].Print()
	}
}

func (f ArrayList) Last() Pair {
	return f.array[f.last]
}

/**
* A move ArrayList, only because I don't know how generics work in go
*/

type MoveArrayList struct{
	array []Move
	last int
}

func newMoveArrayList() *MoveArrayList {
	f := new(MoveArrayList)
	f.array = make([]Move, 0, 10)
	f.last = -1
	return f
}

func (f *MoveArrayList) add(move Move) {
	if f.last+1 >= cap(f.array) {
		newArray := make([]Move, len(f.array), (cap(f.array)+1)*2)
		copy(newArray, f.array)
		f.array = newArray
	}
	f.last++
	f.array = f.array[:f.last+1]
	f.array[f.last] = move
}

func (f MoveArrayList) member(move Move) bool {
	for i := range f.array {
		if move.equals(f.array[i]) {
			return true
		}
	}
	return false
}

func (f MoveArrayList) length() int {
	return len(f.array)
}

func (f MoveArrayList) Print() {
	for i := range f.array {
		f.array[i].Print()
	}
}

func (f MoveArrayList) Last() Move {
	return f.array[f.last]
}




func main() {
	board := newBoard()
	board.printBoard()
}