package dependency_injection

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer,"Hello, %s", name)
}

/*
func main() {
	//Greet(os.Stdout, "Elodie")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHandler))
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	Greet(w, "world")
}
*/
