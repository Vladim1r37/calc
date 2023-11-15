package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var romanNums1 = [11]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
var romanNums2 = [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var arabicNums = [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
var ops = [4]string{"+", "-", "*", "/"}
var t = ""

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите строку в формате \"a operator b\"")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		regex := regexp.MustCompile("\\s+")
		stringNums := regex.Split(text, -1)
		err := checkInput(stringNums)
		if err != nil {
			fmt.Println(err)
			break
		}
		err = compute(stringNums[0], stringNums[1], stringNums[2])
		if err != nil {
			fmt.Println(err)
		}
	}
}

func checkInput(inputString []string) error {
	if len(inputString) != 3 {
		return errors.New("Неправильное количество чисел во вводе")
	}
	if !(isNumCorrect(inputString[0]) && isNumCorrect(inputString[2])) {
		return errors.New("Некорректные числа")
	}
	if !isSameArgs(inputString[0]) {
		return errors.New("Разные типы чисел")
	}
	if !isOpCorrect(inputString[1]) {
		return errors.New("Некорректная операция")
	}

	return nil
}

func isNumCorrect(s string) bool {
	for i := 0; i < len(romanNums1); i++ {
		if romanNums1[i] == s {
			t = "r"
			return true
		}
	}
	for i := 0; i < len(arabicNums); i++ {
		if arabicNums[i] == s {
			t = "a"
			return true
		}
	}
	return false
}

func isSameArgs(s1 string) bool {
	var nums []string
	if t == "r" {
		nums = romanNums1[:]
	} else {
		nums = arabicNums[:]
	}
	for _, num := range nums {
		if num == s1 {
			return true
		}
	}
	return false
}

func isOpCorrect(s string) bool {
	for _, op := range ops {
		if op == s {
			return true
		}
	}
	return false
}

func compute(s1 string, op string, s2 string) error {
	var d1, d2, res int

	if t == "r" {
		d1 = convRToNum(s1)
		d2 = convRToNum(s2)
	} else {
		d1, _ = strconv.Atoi(s1)
		d2, _ = strconv.Atoi(s2)
	}

	switch op {
	case "+":
		res = d1 + d2

	case "-":
		res = d1 - d2

	case "*":
		res = d1 * d2

	case "/":
		res = d1 / d2
	}
	strRes := strconv.Itoa(res)
	if t == "r" {
		if res < 1 {
			return errors.New("Ошибка вычисления, результат не может быть меньше 1")
		}
		strRoman := ""
		for i := len(strRes); i > 0; i-- {
			if i == 3 {
				strRoman += "C"
				break
			}
			if i == 2 {
				index, _ := strconv.Atoi(string(strRes[0]))
				strRoman += romanNums2[index]
			}
			if i == 1 {
				index, _ := strconv.Atoi(string(strRes[1]))
				strRoman += romanNums2[index]
			}
		}
		fmt.Println(strRoman)
	} else {
		fmt.Println(strRes)
	}
	t = ""
	return nil
}

func convRToNum(s string) int {
	for i, num := range romanNums1 {
		if num == s {
			return i
		}
	}
	return -1
}
