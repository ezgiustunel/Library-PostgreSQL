package infrastructure

import (
	"fmt"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-ezgiustunel/service/domain/book"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB: connects db
func ConnectDB(conString string) *gorm.DB {
	dbURL := "postgres://ezgiustunel:pass@localhost:5432/library"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		panic(fmt.Sprintf("Cannot connect to database : %s", err.Error()))
	}

	db.AutoMigrate(&book.Book{})

	if err != nil {
		panic(err)
	}

	return db
}
