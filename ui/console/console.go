package console

import (
	"fmt"
	"main/logic"
)

func readNum() int {
	num, err := fmt.Scanln()
	if err != nil {
		showMessage(err.Error())
	}

	return num
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
