func productExceptSelf(nums []int) []int {
    pre := 1
    post := 1
    out := make([]int, len(nums))
    for idx, _ := range out {
        out[idx] = pre
        pre = pre * nums[idx]
    }
    for idx := len(nums) - 1; idx >= 0; idx-- {
        out[idx] = out[idx] * post
        post = post * nums[idx]
    }
    return out
}