import "strings"

/*
* Use two pointers starting from both end of the string 
* Move pointers inward 
*.- if their character is non alpha numeric character 
* - or their lower characters are matched 
* Return false if their lower character are not match 
* O(n) for time complexity 
* O(1) for space need 
*/

func isAlNum(s string) bool {
    if ((s <= "Z" && s >= "A") || (s <= "z" && s >= "a") || (s <= "9" && s >= "0")) {
        return true
    }  else {
        return false 
    }
}

func isPalindrome(s string) bool {
    l := 0 
    r := len(s) - 1

    for (l <=r) {
        
        for (l < r && !isAlNum(s[l:l+1])) {
            l++ 
        }
        
        for (l < r && !isAlNum(s[r:r+1])) {
            r-- 
        }
        
        ll := strings.ToLower(s[l:l+1])
        lr := strings.ToLower(s[r:r+1])
                
        if (ll == lr) {
            l++
            r--
        } else {
            return false
        }
    }
    
    return true
}
