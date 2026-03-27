package main

import (
	"strings"
)

type Celsius float64
type Fahrenheit float64

func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

type Language string

func (l Language) IsValid() bool {
	valid := []string{"Go", "Java", "Python", "Rust"}
	for _, v := range valid {
		if strings.EqualFold(string(l), v) {
			return true
		}
	}
	return false
}

func (l Language) Upper() Language {
	return Language(strings.ToUpper(string(l)))
}

type IntSlice []int

func (s IntSlice) Sum() int {
	total := 0
	for _, v := range s {
		total += v
	}
	return total
}

func (s IntSlice) Max() int {
	if len(s) == 0 {
		return 0
	}
	max := s[0]
	for _, v := range s[1:] {
		if v > max {
			max = v
		}
	}
	return max
}
