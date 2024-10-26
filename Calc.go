package main

import (
	"errors"
	"fmt"
	"strconv"
)

func Calc(expression string) (float64, error) {
	number, egor := podschet(expression)
	num, _ := strconv.ParseFloat(number, 64)
	return num, egor
}
func skobki(str string, ind int) (string, int) {
	run := []rune(str)
	sl := ""
	col := 0
	j := 0
	for i := 0; i < len(run); i++ {
		if ind > 0 {
			ind--
			continue
		}
		if run[i] == 40 {
			s, y := skobki(str, i+1)
			sl += s
			ind += y
			j+=y
			continue
		} else if run[i] == 41 {
			sq, _ := podschet(sl)
			col++
			j+=col
			return sq, j
		}
		sl += string(run[i])
		col++
	}
	sq, _ := podschet(sl)
	return sq, col
}
func podschet(str string) (string, error) {
	run := []rune(str)
	q := 0
	slice := make([]string, 1)
	col := 0
	h := 0
	g := false
	openskob := 0
	closeckobe := 0
	colznak := 0
	colsimv := 0
	// для особых случаев
	for i := 0; i < len(run); i++{
		if run[i] == 43 || run[i] == 45 || run[i] == 42 || run[i] == 47 {
		colznak++
		}
		if run[i] == 40{
			h++
			openskob++
		}
		if run[i] == 41{
			closeckobe++
		}
		if run[i] >= 48 && run[i] <= 57 {
			colsimv++
		}
	}
	if openskob!=closeckobe || colznak!=colsimv-1{
		return "", errors.New("ошибочка")
	}
	if h%2==0{g = true}
	for i := 0; i < len(run); i++ {
		if q>0{
			q--
			continue
		}
		if run[i] == 40 || run[i] == 41 {
			if g {
				sq, f := skobki(str, i)
				slice[col] += sq
				q += f
			} else{
				sq, f := skobki(str, i+1)
				slice[col] += sq
				q += f
			}
			continue
		} else if run[i] >= 48 && run[i] <= 57 {
			slice[col] += string(run[i])
			continue
		} else if run[i] == 43 || run[i] == 45 || run[i] == 42 || run[i] == 47 {
			slice = append(slice, string(run[i]))
			slice = append(slice, "")
			col += 2
			continue
		} else {
			return "", errors.New("ошибочка")
		}
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == "*" {
			num1, _ := strconv.ParseFloat(slice[i-1], 64)
			num2, _ := strconv.ParseFloat(slice[i+1], 64)
			slice[i] = fmt.Sprint(num1 * num2)
			slice = remove(slice, i+1)
			slice = remove(slice, i-1)
			i--
		} else if slice[i] == "/" {
			num1, _ := strconv.ParseFloat(slice[i-1], 64)
			num2, _ := strconv.ParseFloat(slice[i+1], 64)
			if num2 == 0 {
				return "", errors.New("На ноль делить нельзя")
			}
			slice[i] = fmt.Sprint(num1 / num2)
			slice = remove(slice, i+1)
			slice = remove(slice, i-1)
			i--
		}
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == "+" {
			num1, _ := strconv.ParseFloat(slice[i-1], 64)
			num2, _ := strconv.ParseFloat(slice[i+1], 64)
			slice[i] = fmt.Sprint(num1 + num2)
			slice = remove(slice, i+1)
			slice = remove(slice, i-1)
			i--
		} else if slice[i] == "-" {
			num1, _ := strconv.ParseFloat(slice[i-1], 64)
			num2, _ := strconv.ParseFloat(slice[i+1], 64)
			slice[i] = fmt.Sprint(num1 - num2)
			slice = remove(slice, i+1)
			slice = remove(slice, i-1)
			i--
		}
	}
	return slice[0], nil
}
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
func main() {
	fmt.Println(Calc("2*(2+2)"))
}