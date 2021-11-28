package convert

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Dialogue() {
	fmt.Println("Hello, hell")
	fmt.Println("Введите температуру в цельсиях")
	temp, _ := readInt()

	fmt.Println("Куда переводим?\n 1: Фаренгейты\n 2: Кельвины")
	typeOfGrad, _ := readInt()

	if typeOfGrad == 1 {
		fmt.Println(temp, " по Цельсию = ", celsiusToFahrenheit(temp), " по Фаренгейту")
	} else if typeOfGrad == 2 {
		fmt.Println(temp, " по Цельсию = ", celsiusToKelvin(temp), " по Кельвину")
	}
}

func readInt() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	number, err := strconv.Atoi(text)
	return number, err
}

func celsiusToKelvin(temp int) int {
	return 273 + temp
}

func celsiusToFahrenheit(temp int) float32 {
	return float32(temp)*9/5 + 32
}
