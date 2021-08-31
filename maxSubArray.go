func max(x,y int) int {
    if (x > y) {
        return x 
    } else {
        return y 
    }
}

// Brute force solution - O(n^2) time complexity 
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

// Dynamic programming - O(n) time complexity 
func maxSubArray(nums []int) int {
    maxSubArray := nums[0]
    curSubArray := nums[0]
    for i:=1; i < len(nums); i++ {
        // if current subArray is negative, throw it away. Otherwise, keep adding to it 
        curSubArray = max(nums[i], nums[i] + curSubArray)
        // update maxSubArray with the best current subArray 
        maxSubArray = max(maxSubArray, curSubArray)
    }
    return maxSubArray
}
