
/*
* Start at a middle row as pivot position. 
* Compare first row of middle vs target 
* - If value > target -> search rows before 
* - If value < target -> search rows starting from middle rows 
* - If value = target -> return true 
* Check 
* O(log(m*n)) for time complexity 
*/

func searchMatrix(matrix [][]int, target int) bool {
    
    l := 0
    r := len(matrix) - 1 
    
    for (l <= r) {
        pivot := l + (r - l) / 2
        firstValue := matrix[pivot][0]
        if (firstValue == target) {
            return true
        } else if (firstValue > target) {
            r = pivot - 1 
        } else {
            l = pivot + 1
        }  
    }
    
    if (r == -1) {
        return false 
    }
    
    // check r-th row 
    fmt.Println("Check row %d: %c", r, matrix[r])
    p1 := 0
    p2 := len(matrix[r])
    
    for i:= 0; i < len(matrix[r]); i++ {
        pivot :=  p1 + (p2-p1) / 2 
        
        if (matrix[r][i] == target) {
            return true
        } else if (matrix[r][i] > target) {
            p2 = pivot - 1 
        } else {
            p1 = pivot + 1
        }  
    }
    
    return false 
}
