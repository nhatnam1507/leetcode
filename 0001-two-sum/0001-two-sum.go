func twoSum(nums []int, target int) []int {
  dict := make(map[int]int)
    for i, num := range nums {
      if idx, exist := dict[target - num]; exist {
        return []int{idx, i}
      } else {
        dict[num] = i
      }
    }
    return []int{}
}