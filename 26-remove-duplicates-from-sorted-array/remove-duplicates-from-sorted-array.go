func removeDuplicates(nums []int) int {
    if len(nums) ==0 {
        return 0
    }
    k := 1
    pivot := nums[0]
    for _, num := range nums {
        if num != pivot {
            nums[k] = num
            pivot = num
            k++
        }
    }
    return k;
}