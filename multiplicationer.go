package multiplicationer

import (
	"log"
	"os"
)

type Multiplyer struct {
	LogAnswer bool
}

func NewMultiplyer(mul bool) *Multiplyer {
	m := Multiplyer{mul}
	if mul {
		file, err := os.OpenFile("logResult.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			log.Println(err)
		}
		log.SetOutput(file)
	}
	return &m
}

func (m *Multiplyer) Multiply(firstValue, secondValue float64) float64 {
	result := firstValue * secondValue
	if m.LogAnswer {
		log.Println("Your answer is", result)
	}
	return result
}
