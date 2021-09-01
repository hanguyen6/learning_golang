import "strings"
// https://leetcode.com/problems/longest-substring-without-repeating-characters


func max(x, y int) int {
    if (x > y) {
        return x 
    }  else {
        return y
    }
}

func havingDuplicate(s string, start int, end int) bool {
    // use direct access table to record the occurence of each character 
    var chars [128]int
    
    for i:= start; i <= end; i++ {
        c := s[i]
        chars[c]++
        if (chars[c] > 1) {
            return false
        }
    }
    return true 
}

// Brute-Force solution
// O(n^3) for time complexity
// O(m) for space complexity 
func lengthOfLongestSubstring(s string) int {
    
    res := 0 
    for i:= 0; i < len(s); i++ {
        for j:=i; j < len(s); j++ {
            // check substring from i to j having 
            if (havingDuplicate(s, i, j)) {
                res = max(res, j - i + 1)
            }
        }
    }
    return res
}