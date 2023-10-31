func isValidSudoku(board [][]byte) bool {
    for i := 0; i < len(board); i++ {
        markRow := make(map[byte]bool)
        markCol := make(map[byte]bool)
        markBox := make(map[byte]bool)
        for j := 0; j < len(board[i]); j++ {
            if _, ok:= markRow[board[i][j]]; ok == true {
                return false
            }
            if board[i][j] != byte('.'){
                markRow[board[i][j]] = true
            }
            if _, ok:= markCol[board[j][i]]; ok == true {
                return false
            } 
            if board[j][i] != byte('.'){
                markCol[board[j][i]] = true
            }
            if _, ok:= markBox[board[(i/3)*3+j/3][(i%3)*3+j%3]]; ok == true {
                return false
            }
            if board[(i/3)*3+j/3][(i%3)*3+j%3] != byte('.') {
                markBox[board[(i/3)*3+j/3][(i%3)*3+j%3]] = true
            }
        }
    }
    return true
}