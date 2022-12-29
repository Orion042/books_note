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
	BookAuthor string
	BookTag    pq.StringArray `gorm:"type:text[]"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func insert_db(db *gorm.DB, book_title string, book_author string, book_tag []string) {

	id, _ := uuid.NewUUID()

	book_info := BookInfo{
		BookID:     id.String(),
		BookTitle:  book_title,
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

func read_db() {

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

//export preprocessing_sql
func preprocessing_sql(book_title, book_author, book_tags *C.char) {
	db := connect_db()

	db.AutoMigrate(&BookInfo{})

	bookTitle := C.GoString(book_title)
	bookAuthor := C.GoString(book_author)
	Tags := C.GoString(book_tags)

	bookTags := strings.Split(Tags, ",")

	insert_db(db, bookTitle, bookAuthor, bookTags)
}

//export check_go
func check_go() {
	fmt.Println("go working")
}

func main() {}
