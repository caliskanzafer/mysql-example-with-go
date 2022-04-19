package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Employee struct {
	ID          string `json:"id"`
	NameSurname string `json:"name_surname"`
}

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{Article{
		Title:   "Test title 1",
		Desc:    "Desc 1",
		Content: "Hello world",
	},
		Article{
			Title:   "Test title 2",
			Desc:    "Desc 2",
			Content: "Hello earth",
		},
	}
	json.NewEncoder(w).Encode(articles)
}

func testPostRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TestPostRequest endpoint hit")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage endpoint hit")
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/testpostrequest", testPostRequest).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
func main() {

	db, err := sql.Open("mysql", "root:1155@tcp(127.0.0.1:3306)/dumansah")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Successfully conntected to MySQL database")

	ui := uuid.New()

	insert, err := db.Query("INSERT INTO employee VALUES ('" + ui.String() + "','zafer caliskan')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	fmt.Println("Successfully inserted into employee table")

	result, err := db.Query("SELECT * from employee")

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		var employee Employee
		err = result.Scan(&employee.ID, &employee.NameSurname)
		fmt.Println(employee.NameSurname)
	}

	handleRequest()
}
