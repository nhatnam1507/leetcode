func topKFrequent(nums []int, k int) []int {
    frequency := make(map[int]int)
    for _, num := range nums {
        frequency[num]++
    }
    type pair struct {
        num int
        count int
    }
    pairList := []pair{}
    for k,v := range frequency {
        pairList = append(pairList, pair{num:k, count:v})
    }
    sort.Slice(pairList, func(i, j int) bool { return pairList[i].count > pairList[j].count })
    res := []int{}
    for idx := 0; idx < k; idx++ {
        res = append(res, pairList[idx].num)
    }
    return res
}