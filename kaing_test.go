package kaing

import "testing"

const boardSize int = 30

func TestGenerateBoard(t *testing.T) {
	board := GenerateBoard(boardSize)

	if len(board) != boardSize || len(board[0]) != boardSize {
		t.Error("Size error")
	}

	last := board[0]
	for i := 1; i < boardSize; i++ {

		if len(board[i]) != len(last) {
			t.Error("All rows are not the same size")
		}

		last = board[i]
	}

}

func TestPlaceDot(t *testing.T) {
	board := GenerateBoard(30)

	board = PlaceDot(board, 1, 3, 3)

	if board[3][3] != 1 {
		t.Error("Dot not placed at right place")
	}

}

func Test_checkCols(t *testing.T) {
	board := GenerateBoard(boardSize)

	// place a kaing in col
	i := 3
	j := 18
	for i < 7 {
		board = PlaceDot(board, 1, i, j)
		i++
	}

	if checkCols(board, 1) {
		t.Error("Win when no kaing in col")
	}

	board = PlaceDot(board, 1, i, j)

	if !checkCols(board, 1) {
		t.Error("Kaing in column, but no win")
	}
}

func Test_checkRows(t *testing.T) {
	board := GenerateBoard(boardSize)

	// place a kaing in row
	i := 18
	j := 3
	for j < 7 {
		board = PlaceDot(board, 1, i, j)
		j++
	}

	if checkRows(board, 1) {
		t.Error("Win when no kaing in col")
	}

	board = PlaceDot(board, 1, i, j)

	if !checkRows(board, 1) {
		t.Error("Kaing in row, but no win")
	}
}

func Test_checkDiags(t *testing.T) {
	board := GenerateBoard(boardSize)

	// place a kaing in diagonale
	i := 0
	j := 3
	for i < 4 && j < 7 {
		board = PlaceDot(board, 2, i, j)
		i++
		j++
	}

	if checkDiags(board, 2) {
		t.Error("Win when no kaing in diagonales")
	}

	board = PlaceDot(board, 2, i, j)

	if !checkDiags(board, 2) {
		t.Error("Kaing in diagonale, but no win")
	}
}

func Test_checkAntiDiags(t *testing.T) {
	board := GenerateBoard(boardSize)

	// place a kaing in anti-diagonale
	i := 20
	j := 16
	for i < 24 && j < 20 {
		board = PlaceDot(board, 2, i, j)
		i--
		j++
	}

	if checkAntiDiags(board, 2) {
		t.Error("Win when no kaing in anti-diagonales")
	}

	board = PlaceDot(board, 2, i, j)

	if !checkAntiDiags(board, 2) {
		t.Error("Kaing in anti-diagonale, but no win")
	}
}

func TestWin(t *testing.T) {
	board := GenerateBoard(boardSize)

	if Win(board, 1) {
		t.Error("Board win when empty")
	}

}
