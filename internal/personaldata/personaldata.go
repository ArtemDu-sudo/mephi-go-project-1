package personaldata

import "fmt"

// Personal хранит персональные данные пользователя.
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print выводит данные пользователя на экран.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
