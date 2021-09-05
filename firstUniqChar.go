/**
* HashTable to store char -> occurence 
* As traversing through the list of character, update char with its occurence
* traversing through the list again  
* track the min index 
* O(n) for time complexity 
* O(m) for space complexity 
*/
func firstUniqChar(s string) int {
    m := map[string]int{}
    
    // store 
    for i:=0; i< len(s); i++ {
        c := s[i:i+1]
        _, found := m[c]
        if (found) {
            m[c]++
        } else {
            m[c] = 1
        }
    }
    
    fmt.Println("unique character: ", m)
    
    ans := len(s)
    for i:=0; i < len(s); i++ {
        c := s[i:i+1]
        occurence := m[c]
        if (occurence == 1 && i < ans) {
            ans = i
        }
    }
    
    if (ans == len(s)) {
        return -1 
    } else {
        return ans
    }
}
