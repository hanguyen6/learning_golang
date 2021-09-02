// https://leetcode.com/problems/intersection-of-two-arrays-ii/
// brute-force solution use O(m*n) time complexity 
// For big data intersect, "broadcast" smaller array, load value occurence into memory and sequential process large array for interesection 
// When two arrays sorted, use two pointer method with linear time scan to compare values and load into result 


// O(m+n) time complexity 
// O(m) space needed 
func intersect(nums1 []int, nums2 []int) []int { 
    // Use a hash table to store map int -> occurence of the smaller array 
    if (len(nums1) > len(nums2)) {
      return intersect(nums2, nums1)
    } 
    
    m := make(map[int]int)
    
    for i:=0; i < len(nums1); i++ {
        _, found := m[nums1[i]]
        if (!found) {
            m[nums1[i]] = 1 
        } else {
            m[nums1[i]]++
        }
    }
    
    res := []int{}
    // Iterate through 2nd array and look for a match in hash table 
    for i:=0; i < len(nums2); i++ {
        _, found := m[nums2[i]]
        if (found && m[nums2[i]] >= 1 ) {
            m[nums2[i]]--
            res = append(res, nums2[i])
        }
    } 

    return res
}
