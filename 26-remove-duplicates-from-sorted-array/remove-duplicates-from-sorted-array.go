func removeDuplicates(nums []int) int {
    if len(nums) ==0 {
        return 0
    }
    k := 1
    for _, num := range nums {
        if num != nums[k - 1] {
            nums[k] = num
            k++
        }
    }
    return k;
}