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

// Use sliding window to expand / contract the substring 
// Use a hashset to store characters in current window.
// slide window to the right 
// if character in hashset, contract the windows from the left 
// O(2*n) for time complexity 
// O(min(m,n)) for space needed 
func lengthOfLongestSubstring(s string) int {
    
    var chars [128]int
    
    res := 0 
    left := 0
    right := 0
    
    for right < len(s) {
        r := s[right]
        chars[r]++
        
        for chars[r] > 1 {
            l := s[left]
            chars[l]--
            left++
        }
        
        res = max(res, right - left + 1)
        right++
    }
    
    return res
        
}

// Optimized sliding window use map of string -> index 
func lengthOfLongestSubstring(s string) int {
    n := len(s)
    ans := 0 
    m := make(map[string]int)
    i := 0
    for j:=0; j < n; j++ {
        c := s[j:j+1]
        _, found := m[c]
        if (found) {
            i = max(i, m[c])
        }
        ans = max(ans, j - i + 1)
        m[c] = j + 1
    }
    return ans
}
