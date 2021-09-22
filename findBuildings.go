// https://leetcode.com/problems/buildings-with-an-ocean-view/

func findBuildings(heights []int) []int {
    // Use a stack to keep track of height or buildings 
    s := make([]int, 0)
    
    // move from left to right 
    for i:=0; i < len(heights); i++ {
        
        // if seeing taller building than the last one in stack
        // Keep removing shorter building in the stack 
        for len(s) > 0 && heights[s[len(s)-1]] <= heights[i] {
            s = s[:(len(s)-1)]
        }
        s = append(s, i)
    }
    
    return s
}


