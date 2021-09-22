// https://leetcode.com/problems/valid-palindrome-ii/
func validPalindrome(s string) bool {
    lo := 0 
    hi := len(s) - 1 
    return isValidPalin(s, lo, hi, false)
}

func isValidPalin(s string, lo int, hi int, changed bool) bool {
        
    for (lo < hi) {
        if (s[lo] != s[hi]) {
                if (changed) {
                    return false 
                } else {
                    return isValidPalin(s, lo + 1, hi, true) || isValidPalin(s, lo, hi - 1, true)
                }
        } else {
            lo++
            hi--
        }
    }

    
        return true
            
    
}
