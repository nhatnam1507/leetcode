func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    dict := make(map[rune]int)
    for i, char := range t {
        dict[char]++
        dict[rune(s[i])]--
    }
    for _, v := range dict {
        if v != 0 {
            return false
        }
    }
    return true
}