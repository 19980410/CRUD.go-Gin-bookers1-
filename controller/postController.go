package controller

import (
	"database/sql"
	"fmt"
	//"strconv"
	//"reflect"
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

func BookAll()(booksArray []Book) {
		//Book.all↓
		rows, err := db.Query("SELECT * FROM books")
		if err != nil{
			panic(err.Error())
		}
		if err != nil {
			return
		}
		
		for rows.Next() {
			m := Book{}
			
			err = rows.Scan(&m.ID, &m.Title, &m.Body)

			if err != nil {
				return
			}

			// b := []string{ strconv.Itoa(m.ID), m.Title, m.Body }
			booksArray = append(booksArray, m )

		}
		return booksArray
		//Book.all↑
}

//1つ目の（）は引数、2つ目は戻り値を指定している
func BookFind(id int)(book Book){

	//book.find↓
	var bookFind = Book{}
	var err = db.QueryRow("SELECT * FROM books WHERE id = ?", id).Scan(&bookFind.ID, &bookFind.Title, &bookFind.Body)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("レコードが存在しません")
	case err != nil:
		panic(err.Error())
	default:
		fmt.Println(bookFind.ID, bookFind.Title, bookFind.Body)
	}

	return bookFind
	//book.find↑
}

func CreateBook(title string, body string)(err error){
	newBook, err := db.Prepare("INSERT INTO books (title, body)  VALUES(?,?)")
	_, err = newBook.Exec(title, body)

	if err != nil{
		panic(err.Error())
	}
	return 
}

func DeleteBook(id int)(err error){

	deleteBook, err := db.Prepare("DELETE FROM books where id = ?")
	if err != nil{
		panic(err.Error())
	}

	result, err := deleteBook.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	rowsAffect, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rowsAffect)
	return
}

func UpdateBook(id int, title string, body string)(err error){
	UpdateBook , err := db.Prepare("UPDATE books SET title=?, body=? WHERE id=?")
	if err != nil{
		panic(err.Error())
	}

	result, err := UpdateBook.Exec(title, body, id)
	if err != nil {
		panic(err.Error())
	}

	row, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(row)
	
	return
}