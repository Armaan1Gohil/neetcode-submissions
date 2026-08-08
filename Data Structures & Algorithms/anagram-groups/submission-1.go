func groupAnagrams(strs []string) [][]string {
	wordTrack := make(map[[26]int][]string)

	for _, word := range strs {
		var count [26]int
		for _, char := range word {
			count[char-'a']++
		}
		wordTrack[count] = append(wordTrack[count], word)
	}

	var res [][]string
	for _, words := range wordTrack {
		res = append(res, words)
	}
	return res
}
