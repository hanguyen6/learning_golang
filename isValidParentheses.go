/*
* Use a stack to keep track of close parentheses 
* While traversing through each character in string 
* - If character is open parentheses, put its close parenthese  into stack 
* - If character is close parentheses, pop one element from stack 
* - return false the pop element is not the same with the close parentheses
* Return true after traversing 
* O(n) for time & space complexity 
*/

func closeParen(s byte) byte {
    if (s == '{') {
        return '}'
    } else if (s == '[') {
        return ']'
    } else if (s == '(') {
        return ')'
    } else {
        return s
    }
}

func isValid(s string) bool {
    stack := make([]byte, 0, len(s))
    
    for i:=0; i < len(s); i++ {
        c := s[i]
        if (c == '{' || c == '(' || c == '[') {
            stack = append(stack, closeParen(c))
        } else {
            if (len(stack) == 0) {
                return false
            } else {
                lastC := stack[len(stack)-1]
                if (lastC != c) {
                    return false
                }
                stack = stack[:len(stack)-1]
            }

        }
    }
    
    if (len(stack) == 0) {
        return true
    } else {
        return false
    }
}
