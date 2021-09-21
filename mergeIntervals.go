// https://leetcode.com/problems/merge-intervals

func merge(intervals [][]int) [][]int {
    i := 0
    ans := make([][]int, 0)
    
    if (len(intervals) == 1) {
        ans = append(ans, intervals[0])
    } else {
        // Sort the interval by the start value 
        
    }
    
    for i < len(intervals) - 1 {
        pair := intervals[i]
        next_pair := intervals[i+1]
        
        if (pair[1] >= next_pair[0]) {
            // Combine two pairs
            var min int 
            if (pair[0] < next_pair[0]) {
                min = pair[0]
            } else {
                min = next_pair[0]
            }
            
            var max int 
            if (pair[1] < next_pair[1]) {
                max = next_pair[1]    
            } else {
                max = pair[1]
            }
            
            newPair := []int{min, max}
            ans = append(ans, newPair)
        } else {
            ans = append(ans, pair)
            ans = append(ans, next_pair)
        }
        
        i++
        i++
    }
    return ans
}
