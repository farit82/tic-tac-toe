package main

import (
	"fmt"
	"strings"
)

const size = 3

// С начало создаем функцию входа в приложения.

func main() {
	fmt.Println("Tic tac toe start")
	startGame()
}

func gameBoard(board []string) {
	fmt.Println("-" + strings.Repeat("-", 4*size))
	for i := 0; i < size; i++ {
		fmt.Printf("| %s | %s | %s |\n", board[0+i*size], board[1+i*size], board[2+i*size])
		fmt.Println("-" + strings.Repeat("-", 4*size))
	}

}

func startGame() {
	board := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	currentStep := "x"
	player := ""
	step := 1

	gameBoard(board)

	for i := step; i < 10; i++ {
		var input int
		fmt.Println("Ходит игрок девелопер. Введите номер (0- выходи из игры):")
		_, err := fmt.Scan(&input)
		if err != nil {
			return
		}

		// Если игрок хочет закончить досрочно.
		if input == 0 {
			fmt.Println("Game end")
			return
		}
		// Если ошибочно ввел, не в ту клетку, которая занята.
		if board[input-1] == "x" || board[input-1] == "o" {
			fmt.Println("Это клетка занята, Попробуйте еще раз")
			i--
			continue
		}
		//  если все нормально, записываем игровое поле.
		if currentStep == "x" {
			board[input-1] = currentStep
			currentStep = "o"
			player = "x"
		} else {
			{
			}
			board[input-1] = "o"
			currentStep = "x"
			player = "o"
		}
		gameBoard(board)
		if CheckWin(board) {
			fmt.Printf("Игрок %s победил!\n", player)
			return
		}
	}
	// Если вышел из цикла, то ходы закончились, а победитель не выявлен.
	fmt.Println("Ничья")

}

func CheckWin(board []string) bool {
	// массив возможных комбинаций победы.
	winnerPositions := [8][3]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}}

	for _, p := range winnerPositions {
		if board[p[0]] == board[p[1]] && board[p[1]] == board[p[2]] {
			return true
		}
	}
	return false
}
