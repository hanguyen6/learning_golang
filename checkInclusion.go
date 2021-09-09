/* https://leetcode.com/problems/permutation-in-string
* Brute-force solution: 
Get all permuation of s1 and check if the generated permuation is a substring of s2 
* Generate all permuation: 
- Take an index (starting from 0) and swap the current element with other other element in the string, toward its right 
- Advance the index to the next element and do the swaping again until the index reach to the end 
*/

/*
*  Swapping elements between i1 and i2 index of string s
* split s = s[0:i1] + s[i1] + s[i1+1:i2] + s[i2] + s[i2:n]
* make news swapped string = s[0:i1] + s[i2] + s[i1+1:i2] + s[i1] + s[i2+1:n]
*/
func swap(s string, i1 int, i2 int) string {
    if (i1 == i2) {
        return s
    }  else {
        return s[0:i1] + s[i2:i2+1] + s[i1+1:i2] + s[i1:i1+1] + s[i2+1:len(s)]
    }
}

/**
* Starting from index l, swap the current element with other element in the string, toward its right 
* Until found a permuation of s1 which is substring of s2 then return true 
* Otherwise, keep getting permutaion of s1 until l = len(s1)
* Backtrack to original string s1 and swaping from l + 1 
*/
func permute(s1 string, s2 string, l int) bool {
    found := false
    if (strings.Index(s2, s1) >= 0) {
            found = true
        } else {
            for i:= l; i < len(s1); i++ {
                s1 = swap(s1, l, i)
                found = permute(s1, s2, l+1)
                if (found) {
                    return true
                }
                s1 = swap(s1, l, i)
            }
        }
    return found
}

func checkInclusion(s1 string, s2 string) bool {    
    return permute(s1, s2, 0)
}
