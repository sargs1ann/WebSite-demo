package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func contactsPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "Contact/contact.html")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index/index.html")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "about/about.html")
}
func servicesHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "service/service.html")
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/golang")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO `users` (`name`, `age`, `email`) VALUES('John', 12, 'sdhsahd@mail.ru')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/contacts/", contactsPage)
	http.HandleFunc("/about/", aboutHandler)
	http.HandleFunc("/services/", servicesHandler)
	fmt.Println("Сервер запущен на порту 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
