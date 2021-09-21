// https://leetcode.com/problems/sort-colors/submissions/
/* Move element around pivot 
func partition(pivot int, nums []int)  {
    
    fmt.Println("pivot: ", nums[pivot])
    lo := 0 
    // Group all elements smaller than pivot 
    for i:=0; i < len(nums); i++ {  
        if (nums[i] < pivot) {
            tmp := nums[i]
            nums[i] = nums[lo]
            nums[lo] = tmp
            lo++
        }  
    }
    fmt.Println("First sort: ", nums)
    
    // Second pass: group elements larger than pivot 
    hi := len(nums) - 1 
    for i:= len(nums) - 1; i >=0; i-- {
        if (nums[i] > pivot) {
            tmp := nums[i]
            nums[i] = nums[hi]
            nums[hi] = tmp
            hi--
            
        }
    }
    
    fmt.Println("Second sort: ", nums)
        
}

*/

/**
* Image three pointers to track rightmost boundary of zeros, leftmost boundary of two and current element under considratio 
* Move curr pointer along the array 
* - if nums[curr] = 0, swap it with nums[po]
* - if nums[curr] = 2, swap it with nums[p2]
*/

func swap(nums []int, i int, j int) {
    tmp := nums[i]
    nums[i] = nums[j]
    nums[j] = tmp 
}

func sortColors(nums []int)  {
    // Initialize pointers 
    p0 := 0
    p2 := len(nums) - 1
    
    curr := 0
    
    for (curr <= p2) {
        if (nums[curr] == 0) {
            // Swap curr vs p0 and move p0 to the left 
            swap(nums, curr, p0)
            p0++
            curr++
        } else if (nums[curr] == 2) {
            // Swap curr vs p2 and move p2 to the right
            swap(nums, curr, p2)
            p2--
        } else {
            curr++
        }
    }
}
