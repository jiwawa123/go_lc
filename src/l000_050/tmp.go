package main

import (
	"strconv"
	"strings"
)

func main() {
	//fmt.Print("hello world")
	lengthOfLongestSubstring("abcabcbb")
	//arr := []int{1, 3}
	//s := "B0B6G0R6R0R6G9"
	numberOfMatches(7)
	//println(s[0:1])
	//println(s[0:3])
	//println(s[0:4])
	//println(balancedStringSplit("RLRRLLRLRL"))
	//leftRightDifference(arr)
	isValid("()")
}

func defangIPaddr(address string) string {
	res := ""
	for _, char := range address {
		if char == '.' {
			res = res + "[.]"
		} else {
			res = res + string(char)
		}
	}
	return res
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	res := 0
	index := -1
	for i := 0; i < len(s); i++ {
		v, ok := m[s[i]]
		if ok && index <= v {
			index = v
		}
		m[s[i]] = i

		if res < (i - index) {
			res = i - index
		}
	}
	return res
}

func leftRightDifference(nums []int) []int {
	ans := make([]int, len(nums))
	sum := 0 // 计算总和
	for _, v := range nums {
		sum += v
	}
	sumLeft, sumRight := 0, sum
	for i := 0; i < len(nums); i++ {
		sumRight -= nums[i]
		diff := sumRight - sumLeft
		if diff < 0 { // 取绝对值
			diff = -diff
		}
		ans[i] = diff
		sumLeft += nums[i]
	}
	return ans
}

// 2942
func findWordsContaining(words []string, x byte) []int {
	res := []int{}
	s := string(x)
	for index, word := range words {
		if strings.Contains(word, s) {
			res = append(res, index)
		}
	}

	return res
}

// 2656
func maximizeSum(nums []int, k int) int {
	max_var := -1
	for _, n := range nums {
		if max_var < n {
			max_var = n
		}
	}

	right := max_var + k - 1
	return (right + max_var) * k / 2
}

// 2828
func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}

	for index, word := range words {
		if word[0] != s[index] {
			return true
		}
	}
	return true
}

// 2114
func mostWordsFound(sentences []string) int {
	res := 0
	for i := range sentences {
		tmp := len(strings.Split(sentences[i], " "))
		if res < tmp {
			res = tmp
		}
	}
	return res
}

// 83
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}

// 2697
func makeSmallestPalindrome(s string) string {
	left, right := 0, len(s)-1
	res := make([]byte, len(s))
	for left < right {
		if s[left] > s[right] {
			res[left] = s[right]
		} else {
			res[left] = s[left]
		}
		res[right] = res[left]
		left++
		right--
	}
	return string(res)
}

// 1773
func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	res := 0
	for i := 0; i < len(items); i++ {
		if ruleKey == "type" && items[i][0] == ruleValue {
			res++
			continue
		}

		if ruleKey == "color" && items[i][1] == ruleValue {
			res++
			continue
		}

		if ruleKey == "name" && items[i][2] == ruleValue {
			res++
			continue
		}
	}

	return res
}

// 1678
func interpret(command string) string {
	cn := strings.ReplaceAll(command, "(al)", "al")
	return strings.ReplaceAll(cn, "()", "o")
}

// 771
func numJewelsInStones(jewels string, stones string) int {
	res := 0
	j_map := make(map[rune]int)
	for _, j := range jewels {
		j_map[j] = 1
	}

	for _, s := range stones {
		_, ok := j_map[s]
		if ok {
			res++
		}
	}
	return res
}

// 1221
func balancedStringSplit(s string) int {
	res := 0
	l, r := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			l++
		} else {
			r++
		}
		if l == r {
			res++
			l = 0
			r = 0
		}
	}
	return res
}

// LCR 182
func dynamicPassword(password string, target int) string {
	return password[target:] + password[0:target]
}

// 1684
func countConsistentStrings(allowed string, words []string) int {
	m := make(map[byte]int)
	for i := 0; i < len(allowed); i++ {
		m[allowed[i]] = 1
	}
	res := 0
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			_, ok := m[words[i][j]]
			if !ok {
				break
			}
			if j == len(words[i])-1 {
				res++
			}
		}
	}
	return res
}

// 2325
func decodeMessage(key string, message string) string {
	m := make(map[rune]rune)
	res := ""
	a := 'a'
	for _, s := range key {
		if s == ' ' {
			continue
		}
		_, ok := m[s]
		if !ok {
			m[s] = a
			a = a + 1
		}
		if len(m) >= 26 {
			break
		}
	}

	for _, r := range message {
		if r == ' ' {
			res = res + string(r)
			continue
		}
		v, _ := m[r]
		res = res + string(v)
	}
	return res
}

// 2744
func maximumNumberOfStringPairs(words []string) int {
	res := 0
	map_word := make(map[string]int)
	for i := 0; i < len(words); i++ {
		_, ok := map_word[words[i]]
		if ok {
			res++
			delete(map_word, words[i])
		} else {
			tmp := ""
			for j := len(words[i]) - 1; j >= 0; j-- {
				tmp = tmp + string(words[i][j])
			}
			map_word[tmp] = 1
		}
	}

	return res
}

// 804
func uniqueMorseRepresentations(words []string) int {
	dic := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..",
		".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	dic_res := make(map[string]int)
	for i := 0; i < len(words); i++ {
		tmp := ""
		for j := 0; j < len(words[i]); j++ {
			tmp = tmp + dic[words[i][j]-'a']
		}
		dic_res[tmp] = 1
	}
	return len(dic_res)
}

// 2194
func cellsInRange(s string) []string {

	res := []string{}
	for i := s[0]; i <= s[3]; i++ {
		for j := s[1]; j <= s[4]; j++ {
			res = append(res, string(i)+string(j))
		}
	}
	return res
}

// 1614
func maxDepth(s string) int {
	l := 0
	res := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			l++
			break
		case ')':
			l--
			break
		default:
			break
		}
		if res < l {
			res = l
		}
	}
	return res

}

// 2810
func finalString(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] == 'i' {
			t := ""
			for j := len(res) - 1; j >= 0; j-- {
				t = t + string(res[j])
			}
			res = t
		} else {
			res = res + string(s[i])
		}
	}

	return res
}

// 2103
func countPoints(rings string) int {
	m := make(map[byte]map[byte]int)
	res := make(map[byte]int)
	ind := rings[len(rings)-1]
	for i := len(rings) - 1; i >= 0; i-- {
		if rings[i] == 'R' || rings[i] == 'B' || rings[i] == 'G' {
			val, ok := m[ind]
			if !ok {
				val = make(map[byte]int)
			}
			val[rings[i]] = 1
			if len(val) > 2 {
				res[ind] = 1
			}
			m[ind] = val
		} else {
			ind = rings[i]
		}

	}
	return len(res)
}

// 1313
func decompressRLElist(nums []int) []int {
	res := []int{}
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			res = append(res, nums[i+1])
		}
	}

	return res
}

// 1688
func numberOfMatches(n int) int {
	res := 0
	for n > 1 {
		res += n / 2
		if n%2 == 0 {
			n /= 2
		} else {
			n = n/2 + 1
		}
	}
	return res
}

// 2341
func numberOfPairs(nums []int) []int {
	mapN := make(map[int]int)
	for _, n := range nums {
		v, ok := mapN[n]
		if ok {
			mapN[n] = v + 1
		} else {
			mapN[n] = 1
		}
	}

	res, last := 0, 0
	for _, value := range mapN {
		res += value / 2
		if value%2 == 1 {
			last++
		}
	}
	return []int{res, last}
}

// 1662
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	str1, str2 := "", ""
	for i := 0; i < len(word1); i++ {
		str1 += word1[i]
	}

	for i := 0; i < len(word2); i++ {
		str2 += word2[i]
	}
	return str1 == str2
}

// 2956
func findIntersectionValues(nums1 []int, nums2 []int) []int {
	m1, m2 := make(map[int]int), make(map[int]int)
	res := []int{0, 0}
	for i := 0; i < len(nums1); i++ {
		m1[nums1[i]] = 1
	}

	for i := 0; i < len(nums2); i++ {
		m2[nums2[i]] = 1
	}

	for i := 0; i < len(nums1); i++ {
		_, ok := m2[nums1[i]]
		if ok {
			res[0]++
		}
	}

	for i := 0; i < len(nums2); i++ {
		_, ok := m1[nums2[i]]
		if ok {
			res[1]++
		}
	}

	return res
}

// 20
func isValid(s string) bool {
	sli := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(',
			'{', '[':
			sli = append(sli, s[i])
			break
		case ')':
			if sli[len(sli)-1] != '(' {
				return false
			}
			sli = append(sli[:len(sli)-1], sli[len(sli)-1+1:]...)
			break
		case ']':
			if sli[len(sli)-1] != '[' {
				return false
			}
			sli = append(sli[:len(sli)-1], sli[len(sli)-1+1:]...)
			break
		case '}':
			if sli[len(sli)-1] != '{' {
				return false
			}
			sli = append(sli[:len(sli)-1], sli[len(sli)-1+1:]...)
			break
		}
	}

	return len(sli) == 0

}

// 14
func longestCommonPrefix(strs []string) string {
	res := ""
	for i := 0; i < len(strs[0]); i++ {
		l := 0
		for j := 0; j < len(strs) && i < len(strs[j]); j++ {
			if strs[j][i] != strs[0][i] {
				break
			}
			l++
		}
		if l == len(strs) {
			res = strs[0][:i+1]
		} else {
			break
		}
	}

	return res
}

// 21
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	l := &ListNode{
		Val:  0,
		Next: nil,
	}
	l1 := l
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			next := &ListNode{
				Val:  list1.Val,
				Next: nil,
			}
			l1.Next = next
			list1 = list1.Next
			l1 = next
		} else {
			next := &ListNode{
				Val:  list2.Val,
				Next: nil,
			}
			l1.Next = next
			l1 = next
			list2 = list2.Next
		}
	}

	for list1 != nil {
		next := &ListNode{
			Val:  list1.Val,
			Next: nil,
		}
		l1.Next = next
		list1 = list1.Next
		l1 = next
	}

	for list2 != nil {
		next := &ListNode{
			Val:  list2.Val,
			Next: nil,
		}
		l1.Next = next
		l1 = next
		list2 = list2.Next
	}
	return l.Next
}

// 58
func lengthOfLastWord(s string) int {
	arr := strings.Split(strings.Trim(s, " "), " ")
	return len(arr[len(arr)-1])
}

// 38
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	res := "1"
	for i := 1; i <= n; i++ {
		len_ := 1
		tmp := ""
		last := res[0]
		for j := 1; j < len(res); j++ {
			if last == res[j] {
				len_++
			} else {
				tmp += strconv.Itoa(len_) + string(last)
				len_ = 1
				last = res[j]
			}
		}
		tmp += strconv.Itoa(len_) + string(last)
		res = tmp

	}
	return res
}
