package console

import (
	"bufio"
	"fmt"
	"main/logic"
	"os"
	"strconv"
	"strings"
)

func readNum() int {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		showMessage("error: " + err.Error())
	}

	num, err := strconv.Atoi(formatInput(str))
	if err != nil {
		showMessage("error: " + err.Error())
	}

	return num
}

func formatInput(input string) string {
	return strings.TrimSpace(strings.Trim(input, "\n"))
}

func showMessage(message any) {
	fmt.Println(message)
}

func Run() {
	showMessage("Введите число:")
	num := readNum()
	if logic.IsFibonacci(num) {
		showMessage(logic.FindNext(num))
	} else {
		showMessage(logic.FindNearest(num))
	}
}
