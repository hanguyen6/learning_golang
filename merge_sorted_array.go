// import "sort"

// O((n+m)*log(n+m)) for time complexity 


func merge(nums1 []int, m int, nums2 []int, n int)  {
    for i:=0; i < n; i++ {
        nums1[i+m] = nums2[i]
    }
    sort.Ints(nums1)
}



// Use three pointers
// O(m+n) for time complexity 
// O(m) for added space 
func merge(nums1 []int, m int, nums2 []int, n int)  {
    nums1_cp := make([]int, m, m)
    for i:=0; i< m; i++ {
        nums1_cp[i] = nums1[i]
    }
    
    p1 := 0 
    p2 := 0 
    
    // Move pointer through nums1  
    for i:=0; i < m +n ; i++ {
        if (p2 >= n || p1 < m && nums1_cp[p1] <= nums2[p2]) {
            nums1[i] = nums1_cp[p1]
            p1 +=1
        } else {
            nums1[i] = nums2[p2]
            p2 +=1
        }
    }
}


// Start from end of nums1
// O(1) space 
func merge(nums1 []int, m int, nums2 []int, n int)  {
    p1 := m -1 
    p2 := n - 1
    for i:= m + n - 1; i >=0; i-- {
        if (p2 < 0) {
            break
        }
        if (p1 >=0 && nums1[p1] > nums2[p2]) {
            nums1[i] = nums1[p1]
            p1--
        } else {
            nums1[i] = nums2[p2]
            p2--
        }
    }
}
