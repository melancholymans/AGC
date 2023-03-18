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
	c, _ := strconv.Atoi(s[2])
	if (a + b + 1) >= c {
		fmt.Fprintln(writer, b+c)
	} else {
		fmt.Fprintln(writer, b+a+b+1)
	}
}
