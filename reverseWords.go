/*
* While traversing through the string 
* - Detect word via white space 
* - Append word in reversed order into returning string 
*  O(n) for time and space complexity 
*/

func reverseWords(s string) string {
    rs := make([]byte, 0)
    
    p:= 0
    for i:=0; i < len(s); i++ {
        c := s[i]
        
        if (c == ' ') {
            for r:=i-1; r >= p; r-- {
                rs = append(rs, s[r])
            }
            rs = append(rs, c)
            p = i + 1
        }
    }
    
    for r:= len(s)-1; r >= p; r-- {
        rs = append(rs, s[r])
    }
    
    return string(rs)
    
}
