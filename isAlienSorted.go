// https://leetcode.com/problems/verifying-an-alien-dictionary

func isAlienSorted(words []string, order string) bool {
    indexes := make([]int, 26)
    
    // Store the order of the characters in a array 
    for i:=0; i < len(order); i++ {
        indexes[order[i] - 'a'] = i
    }
    
    // compare a pair of words at once 
    for i:=0; i < len(words) - 1 ; i++ {
        
        // matching each character in two words 
        for j:=0; j < len(words[i]); j ++ {
            
            // for words are like ("apple", "app") just return false 
            if ( j >= len(words[i + 1])) {
                return false 
            }
            
            // when characters are different, check its order in the indexes array
            if (words[i][j] != words[i+1][j]) {
                c1_index := words[i][j] - 'a'
                c2_index := words[i+1][j] - 'a'
                if (indexes[c1_index] > indexes[c2_index]) {
                    return false 
                }
                
                // No need to check the remaining characters, 
                break 
            }
        }
    }
    return true 
    
    
}
