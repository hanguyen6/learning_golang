/** 
* O(m*n) for time and space complexity 
*/
func matrixReshape(mat [][]int, r int, c int) [][]int {
    
    // When to return original matrix 
    if (len(mat) == 0 || r * c != len(mat) * len(mat[0])) {
        return mat
    }
    
    // Extract all elements into a queue 
    queue := []int{}
    for i := 0; i < len(mat); i++ {
        for j:=0; j < len(mat[i]); j++ {
            queue = append(queue, mat[i][j])
        }
    }
    
    // Rebuild new matrix from the queue 
    m := make([][]int, r)
    for i := 0; i < r; i++ {
        m[i] = make([]int, c)    
    }
    
    for i:=0; i < r; i++ {
        for j:=0; j < c; j++ {
            head := queue[0]
            queue = queue[1:]
            m[i][j] = head     
        }
    }
    return m
}

// Refactor without using queue 
func matrixReshape(mat [][]int, r int, c int) [][]int {
    
    // When to return original matrix 
    if (len(mat) == 0 || r * c != len(mat) * len(mat[0])) {
        return mat
    }
    
    
    // Rebuild new matrix
    m := make([][]int, r)
    for i := 0; i < r; i++ {
        m[i] = make([]int, c)    
    }
    
    
    count := 0 
    for i:=0; i < len(mat); i++ {
        for j:=0; j < len(mat[0]); j++ {
            m[count / c][ count % c] = mat[i][j]
            count++
        }
    }
    return m
}
