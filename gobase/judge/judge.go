package judge

import "fmt"

func fullthrough() {
	a := 3
	switch a {
	case 3:
		fmt.Println("1")
		fallthrough //直接跳到下一步
	case 1:
		fmt.Println("2")
	default:
		fmt.Println("3")
	}
}

func doubleif() {
	a := 3
	if a > 10 {
		fmt.Println("1")
	} else if a < 10 {
		fmt.Println("2")
	} else if a == 3 {
		fmt.Println("3")
	}
}

func continua() {
	a := 0
	for a < 10 {
		a++
		if a%2 == 0 {
			continue
		}
		fmt.Println(a)

	}
}

func Selectt() {
	select {}
}

func init() {
	//fullthrough()
	//doubleif()
	//continua()
}
