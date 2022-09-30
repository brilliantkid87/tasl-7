// package main

// import "fmt"

// func main() {

// 	fmt.Println("Brian")

// 	var name string
// 	name = "Brian"

// 	var studentName = "Lauren"
// 	studentName = "Bagus"

// 	batchName := "Batch 40"

// 	var bialngan1 = 20
// 	var bilangan2 = 50

// 	fmt.Println(bialngan1 + bilangan2)

// 	fmt.Println(studentName)
// 	fmt.Println(name)
// 	fmt.Println(batchName)

// }
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	// path folder public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// routing
	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/project", project).Methods("GET")
	route.HandleFunc("/project-detail/{name}", projectDetail).Methods("GET")
	route.HandleFunc("/form-project", formAddProject).Methods("GET")
	route.HandleFunc("/add-project", addProject).Methods("POST")

	fmt.Println("Server running on port 5000")
	http.ListenAndServe("localhost:5000", route)

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		w.Write([]byte("message :" + err.Error()))
		return

	}

	tmpl.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("message :" + err.Error()))
		return

	}

	tmpl.Execute(w, nil)
}

func project(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/project.html")

	if err != nil {
		w.Write([]byte("message :" + err.Error()))
		return

	}

	tmpl.Execute(w, nil)
}

func addProject(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Title : " + r.PostForm.Get("input-project"))
	fmt.Println("Description : " + r.PostForm.Get("input-description"))
	fmt.Println("Date : " + r.PostForm.Get("input-date"))

	http.Redirect(w, r, "/project", http.StatusMovedPermanently)
}

func formAddProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/add-project.html")

	if err != nil {
		w.Write([]byte("message :" + err.Error()))
		return

	}

	tmpl.Execute(w, nil)
}

func projectDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/project-detail.html")

	if err != nil {
		w.Write([]byte("message :" + err.Error()))
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	fmt.Println(id)

	data := map[string]interface{}{
		"Title":   "A Fullstack Enginner",
		"Content": "Lorem ipsum dolor sit amet consectetur adipisicing elit. Qui, necessitatibus consequuntur. Dolor ipsam itaque nam?",
		"Id":      "Id",
	}

	tmpl.Execute(w, data)
}
