package main

import "fmt"

type Season int64

const (
	Undefined Season = iota
	Spring
	Summer
	Autumn
	Winter
)

func (s Season) String() string {
	switch s {
	case Undefined:
		return ""
	case Spring:
		return "spring"
	case Summer:
		return "summer"
	case Autumn:
		return "autumn"
	case Winter:
		return "winter"
	}
	return "unknown"
}

func main() {
	season1 := Spring
	season2 := Winter

	fmt.Printf(
		"%s == %s is %t\n",
		season1, season2, season1 == season2,
	)
}
