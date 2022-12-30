package main

import (
	"C"
	"fmt"
	"log"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
)

type BookInfo struct {
	BookID     string `gorm:"primary_key"`
	BookTitle  string `gorm:"primary_key"`
	FileName   string
	BookAuthor string
	BookTag    pq.StringArray `gorm:"type:text[]"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Tags struct {
	BookID  string
	BookTag string
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
	if result.RowsAffected < 1 {
		fmt.Println("挿入件数0件")
	}
}

//export search_bookTitle_db
func search_bookTitle_db(target *C.char) *C.char {
	db := connect_db()

	book_info := BookInfo{}

	search_book := C.GoString(target)

	result := db.Where("book_title = ?", search_book).Find(&book_info)

	if result.RowsAffected < 1 {
		fmt.Println("書籍が存在しません")
		return C.CString("None")
	}

	return C.CString(book_info.FileName)
}

//export search_author_db
func search_author_db(target *C.char) *C.char {
	db := connect_db()

	book_info := []BookInfo{}

	search_book := C.GoString(target)

	result := db.Where("book_author = ?", search_book).Find(&book_info)

	if result.RowsAffected < 1 {
		fmt.Println("著者が存在しません")
		return C.CString("None")
	}

	loop := len(book_info)

	for i := 0; i < loop; i++ {
		if i == 0 {
			fmt.Println("========================")
		}
		book := book_info[i]

		fmt.Println("書籍名 : ", book.BookTitle)
		fmt.Println("著者名 : ", book.BookAuthor)
		fmt.Println("タグ名 : ", array2string(book.BookTag))
		fmt.Println("========================")
	}

	//result := strings.Join(file_name, ",")

	//return C.CString(result)
	return C.CString("Exist")
}

//export show_all_tags
func show_all_tags() {
	db := connect_db()

	book_info := []BookInfo{}

	result := db.Select("book_tag").Find(&book_info)

	if result.RowsAffected < 1 {
		fmt.Println("ファイルが存在しません")
	}

	loop := len(book_info)

	var book []string

	for i := 0; i < loop; i++ {
		book = book_info[i].BookTag
	}
	fmt.Println("============ タグ一覧 ============")
	fmt.Println(array2string(book))
	fmt.Println("==================================")
}

//export search_tags_db
func search_tags_db(target *C.char) {
	db := connect_db()

	book_info := []BookInfo{}

	search_book := C.GoString(target)

	result := db.Where("book_tag = ?", search_book).Find(&book_info)

	if result.RowsAffected < 1 {
		fmt.Println("ファイルが存在しません")
	}

	loop := len(book_info)

	for i := 0; i < loop; i++ {
		if i == 0 {
			fmt.Println("========================")
		}
		book := book_info[i]

		fmt.Println("書籍名 : ", book.BookTitle)
		fmt.Println("著者名 : ", book.BookAuthor)
		fmt.Println("タグ名 : ", array2string(book.BookTag))
		fmt.Println("========================")
	}
}

//export show_all_db
func show_all_db() {
	db := connect_db()

	book_info := []BookInfo{}

	result := db.Find(&book_info)

	if result.RowsAffected < 1 {
		fmt.Println("ファイルが存在しません")
	}

	loop := len(book_info)

	for i := 0; i < loop; i++ {
		if i == 0 {
			fmt.Println("========================")
		}
		book := book_info[i]

		fmt.Println("書籍名 : ", book.BookTitle)
		fmt.Println("著者名 : ", book.BookAuthor)
		fmt.Println("タグ名 : ", array2string(book.BookTag))
		fmt.Println("========================")
	}
}

func array2string(array []string) string {
	tag := strings.Join(array, ",")

	return tag
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
