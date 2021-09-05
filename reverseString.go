/** O(n) time complexity and O(1) space 
* Use two pointer, one starts from beginning of array, one starts from the end 
* Switch values at two pointer's position 
* Move left pointer to the right and right pointer to the left until 
* left pointer > right pointer 
*/
func reverseString(s []byte)  {
    l := 0
    r := len(s) - 1
    
    for (l <= r) {
        tmp := s[l] 
        s[l] = s[r]
        s[r] = tmp
        l++
        r--
    }
}
