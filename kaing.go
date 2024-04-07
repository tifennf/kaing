package kaing

import (
	"context"
	"sync"
)

func GenerateBoard(n int) [][]int {
	// generate a n x n board
	// pre: n >= 5

	res := make([][]int, n)

	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	return res
}

func PlaceDot(board [][]int, p int, x int, y int) [][]int {
	// place p dot in (x,y) board's case, no board mutation
	// pre: (x,y) is a valid case coordinate and empty

	boardCopy := make([][]int, len(board))
	copy(boardCopy, board)

	boardCopy[x][y] = p

	return boardCopy
}

func checkCols(board [][]int, p int) bool {
	// check for a Kaing in every columns
	// pre: p > 0

	n := len(board)

	j := 0
	for j < n {
		count := 0
		i := 0

		for i < n {
			if board[i][j] == p {
				count++
			} else {
				count = 0
			}

			if count == 5 {
				return true
			}

			i++
		}

		j++
	}

	return false
}

func checkRows(board [][]int, p int) bool {
	// check for a for a Kaing in every rows
	// pre: p > 0

	n := len(board)

	i := 0
	for i < n {
		count := 0
		j := 0

		for j < n {

			if board[i][j] == p {
				count++
			} else {
				count = 0
			}

			if count == 5 {
				return true
			}

			j++
		}

		i++
	}

	return false

}

func checkDiags(board [][]int, p int) bool {
	// check for a Kaing in every diagonales
	// pre: p > 0

	n := len(board)

	kSup := 1 + 2*(n-1)

	k := 0
	for k < kSup {
		count := 0

		var i int
		var j int

		if k < n {
			i = n - k - 1
			j = 0
		} else {
			i = 0
			j = k%n + 1
		}

		for i < n && j < n {
			if board[i][j] == p {
				count++
			} else {
				count = 0
			}

			if count == 5 {
				return true
			}

			i++
			j++
		}

		k++
	}

	return false
}

func checkAntiDiags(board [][]int, p int) bool {
	// check a for a Kaing in every anti-diagonales
	// pre: p > 0 && board is a valid board with size >= 5

	n := len(board)

	iBasis := 4
	jBasis := 0

	// check the anti-diagonales, no need for the last 5 columns or first 5 rows
	for iBasis < n && jBasis < n-5 {
		count := 0

		// initial point of the anti-diagonale
		i := iBasis
		j := jBasis

		// iterate over the anti-diagonale to check for a Kaing
		for i >= 0 && j < n {
			if board[i][j] == p {
				count++
			} else {
				count = 0
			}

			if count == 5 {
				return true
			}

			i--
			j++
		}

		// if we are not at the last anti-diagonale, we move to the next one
		if iBasis < n-1 {
			iBasis++
		} else {
			jBasis++
		}
	}

	return false
}

func Win(board [][]int, p int) bool {
	// check if player p did a Kaing
	// pre: p > 0

	nb_checks := 4

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	out := make(chan bool, nb_checks)

	checks := []func(board [][]int, p int) bool{
		checkCols,
		checkDiags,
		checkAntiDiags,
		checkRows,
	}

	wg.Add(nb_checks)
	for _, check := range checks {
		go func(check func(board [][]int, p int) bool) {
			defer wg.Done()

			if check(board, p) {
				select {
				case out <- true:
				case <-ctx.Done():
				}
				cancel()
			}
		}(check)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for result := range out {
		if result {
			return true
		}
	}

	return false
}

func PrintBoard(board [][]int) {
	// print the board
	// pre: board is a valid board

	n := len(board)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			print(board[i][j], " ")
		}
		println()
	}
}
