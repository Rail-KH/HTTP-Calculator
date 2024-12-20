package calculation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	sumbols = map[string]int{
		"+": 0,
		"-": 0,
		"*": 1,
		"/": 1,
		"(": 2,
		")": 3,
		"0": -1,
		"1": -1,
		"2": -1,
		"3": -1,
		"4": -1,
		"5": -1,
		"6": -1,
		"7": -1,
		"8": -1,
		"9": -1,
		".": -2}
)

func Сalculation(operations []string) ([]string, error) {
	priority_of_the_operation := 1
	for len(operations) != 1 {
		if priority_of_the_operation == 1 {
			for i := 0; i < len(operations); i++ {
				if operations[i] == "/" || operations[i] == "*" {
					x, err_x := strconv.ParseFloat(operations[i-1], 64)
					y, err_y := strconv.ParseFloat(operations[i+1], 64)
					if err_x != nil || err_y != nil {
						return []string{"error:"}, errors.New("unprocessable_entity")
					}
					switch operations[i] {
					case "*":
						elem := append(operations[:i-1], fmt.Sprintf("%f", x*y))
						if len(operations)-1 > i+2 {
							operations = append(elem, operations[i+2:]...)
						} else {
							operations = elem
						}
					case "/":
						if y == 0 {
							return []string{"error:"}, errors.New("internal_server_error")
						}
						elem := append(operations[:i-1], fmt.Sprintf("%f", x/y))
						if len(operations)-1 > i+2 {
							operations = append(elem, operations[i+2:]...)
						} else {
							operations = elem
						}

					}
					break
				}
				flag := false
				for i := 0; i < len(operations); i++ {
					if sumbols[operations[i]] == 1 {
						flag = true
					}
				}
				if !flag {
					priority_of_the_operation = 0
					break
				}
			}
		}
		if priority_of_the_operation == 0 {
			for i := 0; i < len(operations); i++ {
				if operations[i] == "-" || operations[i] == "+" {
					x, err_x := strconv.ParseFloat(operations[i-1], 64)
					y, err_y := strconv.ParseFloat(operations[i+1], 64)
					if err_x != nil || err_y != nil {
						return []string{"error:"}, errors.New("unprocessable_entity")
					}
					switch operations[i] {
					case "+":
						elem := append(operations[:i-1], fmt.Sprintf("%f", x+y))
						if len(operations)-1 > i+2 {
							operations = append(elem, operations[i+2:]...)
						} else {
							operations = elem
						}

					case "-":
						elem := append(operations[:i-1], fmt.Sprintf("%f", x-y))
						if len(operations)-1 > i+2 {
							operations = append(elem, operations[i+2:]...)
						} else {
							operations = elem
						}

					}
					break
				}
				if len(operations) == i+1 {
					priority_of_the_operation = 0
					break
				}
			}

		}
	}
	return operations, nil

}

func PunctumCounter(str string) int {
	counter := 0
	for _, char := range str {
		if string(char) == "." {
			counter += 1
		}
	}
	return counter
}

func SettingPriorities(expression string) ([]string, error) {
	current_simbol := ""
	first_op := false
	var operations []string
	flag := false
	for _, char := range expression {
		if string(char) == " " {
			continue
		}
		if _, ok := sumbols[string(char)]; ok {
			flag = true
		}
		if !flag {
			return []string{"error:"}, errors.New("unprocessable_entity")
		}
		if !first_op {
			if sumbols[string(char)] == -1 || sumbols[string(char)] == 2 {
				current_simbol = string(char)

			} else {
				return []string{"error:"}, errors.New("unprocessable_entity")
			}
			first_op = true
		} else {
			if sumbols[string(current_simbol[len(current_simbol)-1])] == -1 {
				if sumbols[string(char)] == -1 {
					current_simbol += string(char)
				} else if sumbols[string(char)] == 0 || sumbols[string(char)] == 1 || sumbols[string(char)] == 3 {
					operations = append(operations, current_simbol)
					current_simbol = string(char)
				} else if sumbols[string(char)] == 2 {
					return []string{"error:"}, errors.New("unprocessable_entity")
				} else if sumbols[string(char)] == -2 && PunctumCounter(current_simbol) == 0 {
					current_simbol += string(char)
				}

			} else if 0 <= sumbols[string(current_simbol[len(current_simbol)-1])] && sumbols[string(current_simbol[len(current_simbol)-1])] <= 2 {
				if sumbols[string(char)] == -1 {
					operations = append(operations, current_simbol)
					current_simbol = string(char)
				} else if sumbols[string(char)] == 0 || sumbols[string(char)] == 1 || sumbols[string(char)] == 3 {
					return []string{"error:"}, errors.New("unprocessable_entity")
				} else if sumbols[string(char)] == 2 {
					operations = append(operations, current_simbol)
					current_simbol = string(char)
				} else if sumbols[string(char)] == -2 {
					return []string{"error:"}, errors.New("unprocessable_entity")
				}
			} else if sumbols[string(current_simbol[len(current_simbol)-1])] == 3 {
				if sumbols[string(char)] == -1 {
					return []string{"error:"}, errors.New("unprocessable_entity")
				} else if sumbols[string(char)] == 0 || sumbols[string(char)] == 1 || sumbols[string(char)] == 3 {
					operations = append(operations, current_simbol)
					current_simbol = string(char)
				} else if sumbols[string(char)] == 2 {
					return []string{"error:"}, errors.New("unprocessable_entity")
				} else if sumbols[string(char)] == -2 {
					return []string{"error:"}, errors.New("unprocessable_entity")
				}
			} else if sumbols[string(current_simbol[len(current_simbol)-1])] == -2 {
				if sumbols[string(char)] == -1 && PunctumCounter(current_simbol) == 1 {
					current_simbol += string(char)
				} else if sumbols[string(char)] == 0 || sumbols[string(char)] == 1 || sumbols[string(char)] == 3 {
					return []string{"error:"}, errors.New("unprocessable_entity")
				} else if sumbols[string(char)] == 2 {
					return []string{"error:"}, errors.New("unprocessable_entity")
				} else if sumbols[string(char)] == -2 {
					return []string{"error:"}, errors.New("unprocessable_entity")
				}
			}
		}
	}
	if sumbols[string(current_simbol[len(current_simbol)-1])] == -1 || sumbols[string(current_simbol[len(current_simbol)-1])] == 3 {
		operations = append(operations, current_simbol)
	} else {
		return []string{"error:"}, errors.New("unprocessable_entity")
	}

	return operations, nil
}

func CorrectBrackets(operations []string) bool {
	var stack []string
	for _, char := range operations {
		if string(char) == ")" || string(char) == "(" {
			stack = append(stack, string(char))
		}
	}
	counter := 0
	for _, char := range stack {
		if string(char) == "(" {
			counter++
		} else if string(char) == ")" {
			counter--
		}
		if counter < 0 {
			return false
		}
	}
	return counter == 0
}

func BracketsCounter(operations []string) int {
	counter := 0
	for _, char := range operations {
		if string(char) == ")" || string(char) == "(" {
			counter += 1
		}
	}
	return counter
}

func Calcul(operations []string) (string, error) {
	for BracketsCounter(operations) != 0 {
		counter := 0
		first_found := 0
		last_found := 0
		if CorrectBrackets(operations) {
			for i, char := range operations {
				if string(char) == "(" {
					counter += 1
					first_found = i
				} else if string(char) == ")" {
					var elem []string
					counter -= 1
					last_found = i
					answer, err := Сalculation(operations[first_found+1 : last_found])
					if err != nil {
						return "", err
					}
					elem = append(operations[:first_found], answer[0])
					operations = append(elem, operations[last_found+1:]...)
					break
				}

			}
		} else {
			return "", errors.New("unprocessable_entity")
		}
	}
	answer, err := Сalculation(operations)
	if err != nil {
		return strings.Join(answer, ""), err
	}
	return strings.Join(answer, ""), nil
}

func Calc(expression string) (float64, error) {
	if strings.Replace(expression, " ", "", -1) == "" {
		return 0, errors.New("unprocessable_entity")
	}
	operations, err1 := SettingPriorities(expression)
	if err1 != nil {
		return 0, err1
	} else {
		answer, err2 := Calcul(operations)
		if err2 != nil {
			return 0, err2
		} else {
			res, err := strconv.ParseFloat(answer, 64)
			if err != nil {
				return 0, err
			}
			return res, nil
		}
	}

}
