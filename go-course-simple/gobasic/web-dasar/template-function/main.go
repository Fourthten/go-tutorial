package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Superhero struct {
	Name    string
	Alias   string
	Friends []string
}

func (s Superhero) SayHello(from string, message string) string {
	return fmt.Sprintf("%s said: \"%s\"", from, message)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Superhero{
			Name:    "Clark Kent",
			Alias:   "Superman",
			Friends: []string{"Batman", "Supergirl", "Wonder"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

// Operator    Penjelasan                      Analogi
// eq 	        equal, sama dengan              a == b
// ne          not equal, tidak sama dengan    a != b
// lt          lower than, lebih kecil         a < b
// le          lower than or equal,
//             lebih kecil atau sama dengan    a <= b
// gt          greater than, lebih besar       a > b
// ge          greater than or equal,
//             lebih besar atau sama dengan    a >= b

// print	fmt.Sprint
// printf	fmt.Sprintf
// println	fmt.Sprintln
// printf "%s my name is %s", "Hello", "Setia"
// fmt.Sprintf("%s my name is %s", "Hello", "Setia")
