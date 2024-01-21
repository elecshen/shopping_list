package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)
	in := bufio.NewReader(os.Stdin)
	list := make([]string, n)
	var str string
	var founded bool
	for i := 0; i < n; i++ {
		str, _ = in.ReadString('\n')
		str = strings.TrimRight(str, "\r\n")
		list[i] = str
	}
	slices.Sort(list)
	for {
		str, _ = in.ReadString('\n')
		str = strings.TrimRight(str, "\r\n")
		if str == "" {
			break
		}
		founded = false
		for i := 0; i < n; i++ {
			if strings.Contains(list[i], str) {
				fmt.Println(list[i])
				founded = true
				break
			}
		}
		if !founded {
			fmt.Println("Не найдено")
		}
	}
}
