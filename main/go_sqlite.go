package main

import (
	"C"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
)
import "strings"

type BookInfo struct {
	BookID     string `gorm:"primary_key"`
	BookTitle  string `gorm:"primary_key"`
	FileName   string
	BookAuthor string
	BookTag    pq.StringArray `gorm:"type:text[]"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func insert_db(db *gorm.DB, book_title string, file_name string, book_author string, book_tag []string) {

	id, _ := uuid.NewUUID()

	book_info := BookInfo{
		BookID:     id.String(),
		BookTitle:  book_title,
		FileName:   file_name,
		BookAuthor: book_author,
		BookTag:    book_tag,
	}

	result := db.Create(&book_info)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("ok")
	fmt.Println(result.RowsAffected)
}

//export read_db
func read_db(target *C.char) {
	db := connect_db()

	book_info := []BookInfo{}

	db.Find(&book_info)

	loop := len(book_info)

	for i := 0; i < loop; i++ {
		if i == 0 {
			fmt.Println("========================")
		}
		book := book_info[i]
		fmt.Println("書籍名 : ", book.BookTitle)
		fmt.Println("著者名 : ", book.BookAuthor)
		fmt.Println("タグ名 : ", book.BookTag)
		fmt.Println("========================")
	}
}

func update_db() {

}

func delete_db() {

}

func connect_db() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./db/book_note.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

//export go_sql
func go_sql() {

	//db := connect_db()

	//db.AutoMigrate(&BookInfo{})

	//insert_db(db, "name", "author", tag)

}

//export connect
func connect() {
	db := connect_db()

	db.AutoMigrate(&BookInfo{})
}

//export preprocessing_sql
func preprocessing_sql(book_title, file_name, book_author, book_tags *C.char) {
	db := connect_db()

	db.AutoMigrate(&BookInfo{})

	bookTitle := C.GoString(book_title)
	fileName := C.GoString(file_name)
	bookAuthor := C.GoString(book_author)
	Tags := C.GoString(book_tags)

	bookTags := strings.Split(Tags, ",")

	insert_db(db, bookTitle, fileName, bookAuthor, bookTags)
}

//export check_go
func check_go() {
	fmt.Println("go working")
}

func main() {}
