func groupAnagrams(strs []string) [][]string {
    dict := make(map[string][]string)
    for _, str := range strs {
        key := makeKey(str)
        if _, exist := dict[key]; exist != true {
            dict[key] = make([]string, 0)
        }
        if isAnagram(key, str) {
            dict[key] = append(dict[key], str)
        }
    }
    results := [][]string{}
    for _, group := range dict {
        results = append(results, group)
    }
    return results
}

func isAnagram(key, target string) bool {
    lettersMark := [26]int{}
    for idx, char := range key {
        lettersMark[char - 'a']++
        lettersMark[target[idx] - 'a']--
    }
    for _, value := range lettersMark {
        if value != 0 {
            return false
        }
    }
    return true
}

func makeKey(input string) string {
    chars := strings.Split(input, "") //split the string into characters
    sort.Strings(chars) //sort the characters
    return strings.Join(chars, "")
}