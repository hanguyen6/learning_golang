func moveZeroes(nums []int)  {
    // Keep track of non-zero index in the array  
    // update non-zero index with non-zero number   
    lastNonZeroIndex := 0 
    n := len(nums)
    for i:=0; i < n; i++ {
        if (nums[i] != 0) {
            nums[lastNonZeroIndex] = nums[i]
            lastNonZeroIndex += 1 
        }
    }
    for i:= lastNonZeroIndex; i < n; i++ {
        nums[i] = 0
    }
}
