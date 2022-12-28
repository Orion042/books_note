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

type BookInfo struct {
	BookID     string `gorm:"primary_key"`
	BookName   string `gorm:"primary_key"`
	BookAuthor string
	BookTag    pq.StringArray `gorm:"type:text[]"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func insert_db(db *gorm.DB, book_name string, book_author string, book_tag []string) {

	id, _ := uuid.NewUUID()

	book_info := BookInfo{
		BookID:     id.String(),
		BookName:   book_name,
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

//export check_go
func check_go() {
	fmt.Println("go working")
}

//export test_sql
func test_sql() {
	db := connect_db()

	db.AutoMigrate(&BookInfo{})

	tag := []string{"a", "b"}

	insert_db(db, "name", "author", tag)
}

func main() {}
