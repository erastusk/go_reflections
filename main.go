package main

import (
	"github/erastus/reflect/examples"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "GO-WEATHER-APP... ", log.Ldate|log.Ltime|log.Lshortfile)
	// logger.Println("\n====== Types =======")
	// examples.TypeExamples(logger)
	// logger.Println("\n====== Kind =======")
	// examples.KindExamples(logger)
	// logger.Println("\n====== Elem Name =======")
	// examples.ElemExamples(logger)
	// logger.Println("\n====== Structs =======")
	//examples.ReflandSTructs(logger)
	// logger.Println("\n====== ReflandOpenWeatherStruct =======")
	// examples.ReflandOpenWeatherStruct(logger)
	logger.Println("\n====== ReflandOpenWeatherStruct =======")
	examples.ReflandDowntimeStruct(logger)

}
