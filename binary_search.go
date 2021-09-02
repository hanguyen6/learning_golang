// https://leetcode.com/problems/binary-search/
// Given sorted array, choose a pivot as middle element 
// if pivot smaller than the target, search for right half and reset pivot  
// if pivot larger than the target, search for left half and reset pivot  
// stop when found or no longer able to split into 2 parts 
// return -1 if not found 
func search(nums []int, target int) int {
    var pivot int 
    l := 0
    r := len(nums) - 1 
    for (l <= r)  {
        pivot = l + (r - l) / 2 
        if (nums[pivot] == target) {
            return pivot 
        }
        if (nums[pivot] < target) {
            l = pivot + 1 
        } else {
            r = pivot - 1 
        }
    }
    return -1 
}
