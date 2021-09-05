/**
* Construct a hashTable of char -> occurence from the magazine 
* While checking chars in ransomNote
*  - return false if char in ransomNote not found in hashTable or occurence = 0 
*  - reduce occurence whenever found a char in ransomNote 
* Return true 
* O(m) for space complexity 
* O(m + n) for time complexity 
*/
func canConstruct(ransomNote string, magazine string) bool {
    m  := make(map[string]int, len(magazine))
    
    for i:=0; i < len(magazine); i++ {
        c := magazine[i:i+1]
        _, found := m[c]
        if (found) {
            m[c]++
        } else {
            m[c] = 1
        }
    }
    
    for i:=0; i < len(ransomNote); i++ {
        c := ransomNote[i:i+1]
        _, found := m[c]
        
        if (!found || m[c] == 0) {
            return false
        } else {
            m[c]--
        }
    }
    
    return true
}
