/**
* Use hashset/map to track unique number for each row, each column, each box 
* For a board of 9 number -> need a list of 9 hashsets -> 3 lists 
* Iterate over each position in the board and 
* - Update maps if there is no value in the maps 
* - Given current row r, column c, the index of the box map list is: (r/3) * 3 + (c/3) 
* - Return false if found value in either map 
* O(N^2) for time & space complexity where N = board length 
*/
func isValidSudoku(board [][]byte) bool {
  
    rows := make([]map[byte]bool, 9)    
    cols := make([]map[byte]bool, 9)
    boxes := make([]map[byte]bool, 9)
    
    for i:=0; i < 9; i++ {
        rows[i] = make(map[byte]bool)
        cols[i] = make(map[byte]bool)
        boxes[i] = make(map[byte]bool)
    }
    
    var dotChar byte = '.'
    
    for r:=0; r < 9; r++ {
        for c:=0; c < 9; c++ {
            
            val := board[r][c]
            
            if (val == dotChar) {
                continue
            }
            
            _, foundr := rows[r][val]
            if (foundr) {
                return false 
            } else {
                rows[r][val] = true
            }
            
            _, foundc := cols[c][val]
            if (foundc) {
                return false
            } else {
                cols[c][val] = true 
            }
            
            idx := (r / 3) * 3 + c / 3
            _, foundb := boxes[idx][val]
            if (foundb) {
                return false 
            } else {
                boxes[idx][val] = true 
            }
        }
    }
    
    return true
}
