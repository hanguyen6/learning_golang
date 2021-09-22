// https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/
// Use a stack to keep track of index of '(' 
// While traversing through characters in string 
// If seeing ')' while the stack is empty, add its index to a look up map  
// or else, pop from stack the index that match with ')'

func minRemoveToMakeValid(s string) string {
    
    st := make([]int, 0)
    removeIndexes := make(map[int]bool, 0)
    
    for i:=0; i < len(s); i++ {
        c := s[i]
        if (c == '(') {
            st = append(st, i)
        }
        if (c == ')') {
            if (len(st) == 0) {
                removeIndexes[i] = true 
            } else {
                st = st[1:]
            }
        }
    }
    
    //
    for (len(st) > 0) {
        removeIndexes[st[0]] = true
        st = st[1:]
    }
    
    // build a string 
    ans := make([]byte, 0)
    
    for i:=0; i < len(s); i++ {
        _, found := removeIndexes[i]
        if (!found) {
            ans = append(ans, s[i])
        }
    }
    
    return string(ans)
}
