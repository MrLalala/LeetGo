package main

import "fmt"

const (
	Empty = '.'
	Queen = 'Q'
)

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	chess := make([][]byte, n)
	for i := 0; i < n; i++ {
		chess[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			chess[i][j] = Empty
		}
	}

	dfs(chess, 0, &res)
	return res

}

// row行 之前的数据都已经正确的放置了皇后
func dfs(chess [][]byte, row int, res *[][]string) {
	l := len(chess)
	if row == l {
		r := make([]string, 0)
		for i := 0; i < l; i++ {
			r = append(r, string(chess[i]))
		}
		*res = append(*res, r)
		return
	}
	// 一行只能存在一个皇后
	for col := 0; col < l; col++ {
		if !isValid(chess, row, col) {
			continue
		}
		// 交换后换回去
		chess[row][col] = Queen
		dfs(chess, row+1, res)
		chess[row][col] = Empty
	}
}

func isValid(chess [][]byte, row, col int) bool {
	// 检查列
	for i := 0; i <= row-1; i++ {
		if chess[i][col] == Queen {
			return false
		}
	}
	// 检查左上
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chess[i][j] == Queen {
			return false
		}
	}
	// 检查右上
	for i, j := row-1, col+1; i >= 0 && j < len(chess); i, j = i-1, j+1 {
		if chess[i][j] == Queen {
			return false
		}
	}

	return true
}

func solveNQueens2(n int) [][]string {
	chess := make([][]byte, n)
	for i := 0; i < n; i++ {
		chess[i] = make([]byte, n)
	}

	var dfs func(i int)
	var res [][]string

	cols := make([]int, n)   // 同一行, 最多n个值
	ml := make([]int, n*2)   // 主对角线 主对角线的 row-col = const, 最大为(n-1, 0), 最小为(0, n-1), 需要通过+n 来保证都是>0的
	sl := make([]int, 2*n+1) // 次对角线 次对角线的 row+col = const, 最大为(n-1, n-1)

	dfs = func(i int) {
		if i == n {
			// 增加一个解
			vs := make([]string, n)
			vb := make([]byte, n)
			for j := 0; j < n; j++ {
				for k := 0; k < n; k++ {
					if chess[j][k] == 1 {
						vb[k] = Queen
					} else {
						vb[k] = Empty
					}
				}
				vs[j] = string(vb)
			}
			res = append(res, vs)
		} else {
			for j := 0; j < n; j++ {
				// 尝试每列添加
				if cols[j] == 0 && ml[i-j+n] == 0 && sl[i+j] == 0 {
					// 更新列, 主次对角线, 棋盘对应位置的值
					cols[j], ml[i-j+n], sl[i+j], chess[i][j] = 1, 1, 1, 1
					dfs(i + 1)
					// 复原
					cols[j], ml[i-j+n], sl[i+j], chess[i][j] = 0, 0, 0, 0
				}

			}
		}
	}

	dfs(0)
	return res
}

func main() {
	res := solveNQueens2(4)
	for i := 0; i < len(res); i++ {
		for _, str := range res[i] {
			fmt.Println(str)
		}
		fmt.Println("============")
	}
}
