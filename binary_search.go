/**
* https://leetcode.com/problems/binary-search/
* Given sorted array, choose a pivot as middle element 
* if pivot smaller than the target, search for right half and reset pivot  
* if pivot larger than the target, search for left half and reset pivot  
* stop when found or no longer able to split into 2 parts 
* return -1 if not found
* O(log(n) for time complexity 
*/
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


/** 
 * https://leetcode.com/problems/first-bad-version/ 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

/**
* Start checking at pivot or middle position 
* If the checked version is good, all versions before it is good, then search for the right halve only 
* If the checked version is bad, ignore the right half and search only the left part 
* Keep track of bad version with the smallest number 
* Reset pivot position until unable to split 
*/
func min(x int, y  int) int {
    if (x < y) {
        return x 
    } else {
        return y 
    }
}

func firstBadVersion(n int) int {
    var p int 
    fbv := n
    l := 0
    r := n - 1 
    for (l <= r) {
        p = l + (r-l)/2
        if (isBadVersion(p)) {
            fbv = min(fbv, p)
            r = p - 1 
        } else {
            l = p + 1 
        }
    }
    return fbv
}

