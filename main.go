package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

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

	response := map[string]interface{}{
		"Blogs": dataBlog,
	}

	tmpl.Execute(w, response)
}

type Blog struct {
	Title       string
	Description string
	Author      string
	Post_Date   string
}

var dataBlog = []Blog{
	{
		Title:       "Hallo Title",
		Description: "Hallo Description",
	},
}

func addProject(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Title : " + r.PostForm.Get("input-project"))
	// fmt.Println("Description : " + r.PostForm.Get("input-description"))
	// fmt.Println("Start Date : " + r.PostForm.Get("input-sDate"))
	// fmt.Println("End Date : " + r.PostForm.Get("input-eDate"))

	var title = r.PostForm.Get("input-project")
	var description = r.PostForm.Get("input-description")

	var newBlog = Blog{
		Title:       title,
		Description: description,
		Author:      "Brilliant",
		Post_Date:   time.Now().String(),
	}

	dataBlog = append(dataBlog, newBlog)
	fmt.Println(dataBlog)

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
