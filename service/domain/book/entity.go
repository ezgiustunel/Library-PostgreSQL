package book

import (
	errors "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-ezgiustunel/internal/library"
)

type Book struct {
	ID          int `gorm:"primaryKey;autoIncrement" json:"Id"`
	StockNumber int
	PageNumber  int
	Price       float64
	Name        string
	StockCode   string
	Isbn        string
	AuthorName  string
}

// DecreaseStockNumber: checks and decreases stock number for the given book
func (b *Book) DecreaseStockNumber(bookNumber int) (*Book, error) {
	if b.StockNumber >= bookNumber {
		b.StockNumber -= bookNumber
		return b, nil
	}

	return nil, errors.ErrStockNotEnough
}
