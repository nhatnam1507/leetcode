func containsDuplicate(nums []int) bool {
    dict := make(map[int]int)
    for _, num := range nums {
        _, exist := dict[num]
        if exist {
            return true
        }
        dict[num] = 1
    }
    return false
}