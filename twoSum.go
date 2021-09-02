// O(n^2) for time complexity 
func twoSum(nums []int, target int) []int {
    for i:=0; i < len(nums) - 1; i++ {
        for j:=i+1; j < len(nums); j++ {
            s := nums[i] + nums[j]
            if (s == target) {
                return []int{i, j}
            }
        }
    }
    return []int{}
}


// O(n) for time complexity 
// O(n) for space needed
func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    
    for i:=0; i < len(nums); i++ {
        remain := target - nums[i]
        _, found := m[remain]
        
        if (found) {
            return []int{m[remain], i}
        }  
        m[nums[i]] = i 
    }
    
    return []int{}
}
