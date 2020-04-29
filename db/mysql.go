package main

import (
	"database/sql"
	"fmt"
	//使わないけどimportしておく
	_ "github.com/go-sql-driver/mysql"
)

type Book struct{
	ID int
	Title string
	Body string
}

var db *sql.DB

func init(){
	var err error
	db, err = sql.Open("mysql", "root:Runriku1@tcp(127.0.0.1:3306)/bookers_go")
	if err != nil {
		panic(err)
	}
}

func main(){

	index()

	book, err := bookFind(2)
	if err != nil{
		panic(err.Error())
	}
	fmt.Println(book)

	//book.create↓
	createBook("顔", "ミルク")
	//book.create↑


	//book.update↓

	//book.update↑
}


func index(){
		//Book.all↓
		rows, err := db.Query("SELECT * FROM books")
		if err != nil{
			panic(err.Error())
		}
	
		defer rows.Close()
	
		for rows.Next(){
			var book Book
	
			err := rows.Scan(&book.ID, &book.Title, &book.Body)
			if err != nil{
				panic(err.Error())
			}
			fmt.Println(book.ID, book.Title, book.Body)
			
			err = rows.Err()
			if err != nil{
				panic(err.Error())
			}
		}
		//Book.all↑
}

func bookFind(id int)(book Book, err error){

	//book.find↓
	var bookFind = Book{}
	err = db.QueryRow("SELECT * FROM books WHERE id = ?", 1).Scan(&bookFind.ID, &bookFind.Title, &bookFind.Body)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("レコードが存在しません")
	case err != nil:
		panic(err.Error())
	default:
		fmt.Println(bookFind.ID, bookFind.Title, bookFind.Body)
	}

	return
	//book.find↑
}

func createBook(title string, body string)(err error){
	newBook, err := db.Prepare("INSERT INTO books (title, body)  VALUES(?,?)")
	_, err = newBook.Exec(title, body)

	if err != nil{
		panic(err.Error())
	}
	return 
}