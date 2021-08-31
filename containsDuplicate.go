import "sort" 

// Use HashTable to store unique values. 
// O(n) space and O(n) of time complexity 

func containsDuplicate(nums []int) bool {
    m := make(map[int]int)
    for i:=0; i < len(nums); i++ {
        _, exists := m[nums[i]]
        if (exists) {
            return true
        } else {
            m[nums[i]] = 1
        }
    }
    return false 
}

// Sort an array and traverse through the sorted array, comparing consecutive elements 
// O(nlogn) for sorting array, O(1) for space complexity 
func containsDuplicate(nums []int) bool {
    sort.Ints(nums)
    for i:=0; i < len(nums) - 1 ; i++ {
        if (nums[i] == nums[i+1]) {
            return true
        }
    }
    return false
}
