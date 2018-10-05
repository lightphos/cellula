package transport

import "fmt"
import cp "cellula/info"

// Run code
func Run(code string) {
	fmt.Println("Running  ", code)

	cl := cp.Hello{}

	s := cp.Coder(cl).Code()
	fmt.Println("Call ", s)

	//	o, e := exec.Command("go", "run", "code/"+code+".go").Output()
	// //	fmt.Println("done ", string(o), e)
	// if e != nil {
	// 	fmt.Println("", e)
	// } else {
	// 	fmt.Println("", string(o))
	// }
}
