/*
* while traversing the list
* - get the square values of element and put into returning list
* Sort the returning list 
* O(nlogn) + O(n) for time complexity 
* O(n) for space need 
*/
func sortedSquares(nums []int) []int {
    res := []int{}
    for i:=0; i < len(nums); i++ {
        res = append(res, nums[i] * nums[i])
    }
    sort.Ints(res)
    return res
}



/*
* Use two pointer, one starts from beginning, one starts from the end 
* As traversing through the list from the right end
* - Compare square values of the pointer's poisitions 
* - Put larger values into the returning list at proper position 
* - Move right pointer to the left if its larger value was chosen 
* - Move left pointer to the right if its larger value was chosen 
* O(n) for time and space complexity
*/

func sortedSquares(nums []int) []int {
    n := len(nums)
    l := 0
    r := n - 1
    res := make([]int, n)
    
    for i:=n-1; i >=0 ; i-- {
        sq1 := nums[l] * nums[l]
        sq2 := nums[r] * nums[r]
        if (sq1 <= sq2) {
            res[i] = sq2
            r--
        } else {
            res[i] = sq1
            l++
        }
    }
    return res
}
