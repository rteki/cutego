package main

import (
	"fmt"
	"math"

	"github.com/rteki/cutego"
)

var calculator = map[string]func(float64) float64{
	"sqr": func(v float64) float64 {
		return v * v
	},
	"cube": func(v float64) float64 {
		return v * v * v
	},
	"sqrt": func(v float64) float64 {
		return math.Sqrt(v)
	},
}

func onCalculate(val interface{}) {
	hash := val.(map[string]interface{})

	if _, ok := hash["input"].(string); ok {
		handler.Call("show-result", "There's no input! :(")
	}

	result := map[string]string{}

	for field := range hash {

		if field == "input" {
			continue
		}

		if hash[field].(bool) {
			result[field] = fmt.Sprintf("%.2f", calculator[field](hash["input"].(float64)))
		} else {
			result[field] = "disabled"
		}
	}

	handler.Call("show-result", result)

}

var handler *cutego.EventManager

func main() {
	cutego.Init()

	handler = cutego.NewEventManager("handler")

	cutego.LoadQmlEntry("qml/main.qml")

	handler.On("calculate", onCalculate)

	cutego.Start()
}
