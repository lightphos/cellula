package transport

import (
	"fmt"
)

func calculator(a int, b int) {
	fmt.Printf("%d\n", a+b)
}

func msg(s string) {
	fmt.Println("Rcv " + s)
}
func main() {
	// bus := EventBus.New()
	// bus.Subscribe("main:calculator", calculator)
	// bus.Publish("main:calculator", 20, 40)
	// bus.Unsubscribe("main:calculator", calculator)
	// bus.Subscribe("msg", msg)
	// bus.Publish("msg", "===test")
	// bus.Unsubscribe("msg", msg)
	// s := nucleus.Create()
	// fmt.Println("create ", s)
}
