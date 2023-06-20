package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Product struct {
	Name  string
	Desc  string
	Price float64
	Count int
}


// Стартова точка програми. Запускається локальний сервер на 8000 порту. Прослуховується базова сторінка і пейдж 1 і 2
func main() {
	http.HandleFunc("/", page1)
	http.HandleFunc("/page2/", page2)
	http.HandleFunc("/page3/", page3)

	http.ListenAndServe("https://weblaba.onrender.com", nil)
}


func page1(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("header.html", "page1.html", "footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	page.ExecuteTemplate(w, "page1", nil)
}



func page2(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("header.html", "page2.html", "footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	//i:=6
	//Добавлянння 10ти елементів

	list := []Product{}
	list = append(list, Product{"Product 1", "Desc Desc Desc Desc Desc Desc Desc", 1500.0, 1})
	list = append(list, Product{"Product 2", "Desc Desc Desc Desc Desc Desc Desc", 200.0, 2})
	list = append(list, Product{"Product 3", "Desc Desc Desc Desc Desc Desc Desc", 300.0, 3})
	list = append(list, Product{"Product 4", "Desc Desc Desc Desc Desc Desc Desc", 400.0, 4})
	list = append(list, Product{"Product 5", "Desc Desc Desc Desc Desc Desc Desc", 500.0, 5})
	list = append(list, Product{"Product 6", "Desc Desc Desc Desc Desc Desc Desc", 600.0, 6})
	list = append(list, Product{"Product 7", "Desc Desc Desc Desc Desc Desc Desc", 700.0, 7})
	list = append(list, Product{"Product 8", "Desc Desc Desc Desc Desc Desc Desc", 800.0, 8})
	list = append(list, Product{"Product 9", "Desc Desc Desc Desc Desc Desc Desc", 900.0, 9})
	list = append(list, Product{"Product 10", "Desc Desc Desc Desc Desc Desc Desc", 1000.0, 10})

	//list = append(list[:i], list[i+1:])	//видалення і-того елемента

	page.ExecuteTemplate(w, "page2", list)
}

func page3(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("header.html", "page3.html", "footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	page.ExecuteTemplate(w, "page3", nil)
}
