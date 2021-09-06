/**
* Populate hashTable of chars => occurence from one of the strings 
* Use hash table to look up when searching through the other string
* O(n) for time complexity 
*/
func isAnagram(s string, t string) bool {
    if (len(s) != len(t)) {
        return false
    } else {
        m := map[string]int{}
        
        for i:=0; i < len(s); i++ {
            c := s[i:i+1]
            _, found := m[c]
            if (found) {
                m[c]++
            } else {
                m[c] = 1
            }
        }
        
        for i:=0; i < len(t); i++ {
            c:= t[i:i+1]
            _, found := m[c]
            if (found && m[c] > 0) {
                m[c]--
            } else {
                return false
            }
        }
        
        return true
    }
}
