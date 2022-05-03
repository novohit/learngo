package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	lastOccurred := make(map[byte]int) // 记录每个字符最后出现的位置
	l := 0
	res := 0
	for r, b := range []byte(s) {
		// 如果当前字符在此前出现过且位置在窗口中，则将窗口左边界移至不在出现该字符
		if lastIndex, ok := lastOccurred[b]; ok && lastIndex >= l {
			l = lastIndex + 1
		}
		// 更新窗口长度
		if r-l+1 > res {
			res = r - l + 1
		}
		// 记录当前字符最后出现的位置
		lastOccurred[b] = r
	}
	return res
}

// 支持中文
func lengthOfLongestSubstring2(s string) int {
	lastOccurred := make(map[rune]int) // 记录每个字符最后出现的位置
	l := 0
	res := 0
	for r, b := range []rune(s) {
		// 如果当前字符在此前出现过且位置在窗口中，则将窗口左边界移至不在出现该字符
		if lastIndex, ok := lastOccurred[b]; ok && lastIndex >= l {
			l = lastIndex + 1
		}
		// 更新窗口长度
		if r-l+1 > res {
			res = r - l + 1
		}
		// 记录当前字符最后出现的位置
		lastOccurred[b] = r
	}
	return res
}
func main() {
	fmt.Println(lengthOfLongestSubstring("fadhehell"))
	fmt.Println(lengthOfLongestSubstring("aaabc"))
	fmt.Println(lengthOfLongestSubstring("你好好世界"))
	fmt.Println(lengthOfLongestSubstring2("你好好世界"))
	fmt.Println(lengthOfLongestSubstring2("一二三二一"))
}
