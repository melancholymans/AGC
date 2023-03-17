package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	s := strings.Split(sc.Text(), " ")
	a, _ := strconv.Atoi(s[0])
	b, _ := strconv.Atoi(s[1])
	k, _ := strconv.Atoi(s[3])
	if Abs(a-b) < 1000000000000000000 {
		if k%2 == 0 {
			fmt.Fprintln(writer, a-b)
		} else {
			fmt.Fprintln(writer, b-a)
		}
	} else {
		fmt.Fprintln(writer, "Unfair")
	}
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
