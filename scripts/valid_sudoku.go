package scripts

func IsValidSudoku(board [][]byte) bool {
	hashMap := make(map[string]int)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if string(board[i][j]) == "." {
				continue
			}
			rk := "r" + "_" + string(i) + string(board[i][j])
			ck := "c" + "_" + string(j) + string(board[i][j])
			bk := "b" + "_" + string(i/3) + "_" + string(j/3) + string(board[i][j])
			if _, ok := hashMap[rk]; ok {
				return false
			}
			if _, ok := hashMap[ck]; ok {
				return false
			}
			if _, ok := hashMap[bk]; ok {
				return false
			}
			hashMap[rk] = 1
			hashMap[ck] = 1
			hashMap[bk] = 1
		}
	}
	return true
}
