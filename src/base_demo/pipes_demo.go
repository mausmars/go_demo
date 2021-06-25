package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// WordCount 获取一个io.Reader并返回一个map，每个单词作为一个键，它的出现次数为对应的值
func WordCount(f io.Reader) map[string]int {
	result := make(map[string]int)

	// 建立scanner处理文件io.Reader接口
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		result[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	return result
}

func main() {
	fmt.Printf("string: number_of_occurrences\n\n")
	for key, value := range WordCount(os.Stdin) {
		fmt.Printf("%s: %d\n", key, value)
	}
}