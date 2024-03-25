func removeElement(nums []int, val int) int {
    if len(nums) == 0 {return 0}
    i := 0
    j := len(nums) - 1
    for (i <= j) {
        if nums[i] != val {
            i++
        } else {
            nums[i], nums[j] = nums[j], nums[i] 
            j--
        }
    }
    return i 
}