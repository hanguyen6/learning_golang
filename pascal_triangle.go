
/**
* O(numRows^2) time & space complexity 
*/
func generate(numRows int) [][]int {
    res := make([][]int, numRows)
    
    for i := 1; i <= numRows; i++ {
        res[i-1] = make([]int, i)
        
        for j:=0; j < i; j++ {
            if (j == 0) {
                res[i-1][j] = 1
            } else if (j == i - 1) {
                res[i-1][j] = 1
            } else {
                res[i-1][j] = res[i-2][j-1] + res[i-2][j]
            }
        }
    }
    
    return res
}
