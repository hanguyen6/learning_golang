func max(x,y int) int {
    if (x > y) {
        return x 
    } else {
        return y 
    }
}

// Brute force solution 
func maxSubArray(nums []int) int {
    maxSubArray := nums[0]
    for i:=0; i < len(nums); i++ {
        currSubArray := 0
        for j:=i; j < len(nums); j++ {
            currSubArray += nums[j]
            maxSubArray = max(currSubArray, maxSubArray)
        }
    }
    return maxSubArray
}
