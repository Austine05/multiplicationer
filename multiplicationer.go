package multiplicationer

import (
	"log"
	// "os"
)

func Logg(LogAnswer bool) {
	if LogAnswer {
		file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			log.Println(err)
		}
		log.SetOutput(file)
	}
}

func Multiply(firstValue, secondValue float64) float64 {
	result := firstValue * secondValue
	log.Println("Your answer is", result)
	return result
}
