package kaing

func GenerateBoard(n int) [][]int {
	// generate a n x n board
	// 0: empty case; 1: p1 dot; 2: p2 dot

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

	n := len(board)

	kSup := 1 + 2*(n-1)

	k := 0
	for k < kSup {
		count := 0

		var i int
		var j int

		if k < n {
			i = k
			j = 0
		} else {
			i = n - 1
			j = k%n + 1
		}

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

		k++
	}

	return false
}

func Win(board [][]int, p int) bool {
	// check is player p did a Kaing

	return checkCols(board, p) || checkRows(board, p) || checkDiags(board, p) || checkAntiDiags(board, p)
}
