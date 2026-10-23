package actioninfo

import (
	"fmt"
	"log"
)

// DataParser описывает типы, которые умеют разбирать строку данных
// и формировать информацию об активности.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info разбирает каждую строку из dataset с помощью dp и выводит информацию об активности.
// Ошибки разбора и формирования информации логируются.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("ошибка разбора данных %q: %v", data, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка формирования информации для %q: %v", data, err)
			continue
		}

		fmt.Println(info)
	}
}
