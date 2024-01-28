package bookapp

import (
	"bookapp/database"
	"bookapp/gen/books"
	"bookapp/models"
	"context"
	"fmt"
	"log"
	"time"
)

type bookssrvc struct {
	logger *log.Logger
}

// NewBooks returns the books service implementation.
func NewBooks(logger *log.Logger) books.Service {
	return &bookssrvc{logger}
}

// FindAll Books
func (s *bookssrvc) FindAll(ctx context.Context) (res books.BookCollection, err error) {
	bookList := []models.Book{}

	//Find all books that are saved on database
	result := database.DB.Find(&bookList)

	if result.Error != nil {
		return nil, result.Error
	}

	//Iterate through the list of books retrieved from database
	for _, b := range bookList {
		//add the converted Book to the list
		res = append(res, toBook(b))
	}

	return res, nil

}

// Find book by id
func (s *bookssrvc) FindByID(ctx context.Context, p *books.FindByIDPayload) (res *books.Book, err error) {
	//Create an instance of models.Book
	book := models.Book{}

	//Find the book where ID = p.ID
	//result := database.DB.First(&book, uint(*p.ID))

	result := database.DB.First(&book, p.ID)

	if result.Error != nil {

		return nil, books.MakeNotFound(result.Error)
	}

	//Convert to response Object
	res = toBook(book)

	return res, nil
}

// Create a new book.
func (s *bookssrvc) Create(ctx context.Context, p *books.CreatePayload) (res string, err error) {
	//Create an instance of models.Book
	bookToSave := &models.Book{}

	//Copy data from payload to bookToSave
	bookToSave.Title = p.Title
	bookToSave.Author = p.Author

	//Parse the date from String(yyyy-mm-dd) to time.Time
	publishedAtDate, _ := time.Parse("2006-01-02", p.PublishedAt)
	bookToSave.PublishedAt = publishedAtDate

	if p.BookCover != nil {
		bookToSave.BookCoverBytes = p.BookCover.Bytes
		bookToSave.CoverFileType = *p.BookCover.Type
		bookToSave.CoverFileName = *p.BookCover.Name
	}

	//Save the book to the database
	result := database.DB.Create(&bookToSave)

	//Check for errors
	if result.Error != nil {
		return "Failed", result.Error
	}

	return "Created", nil
}

// Update a book.
func (s *bookssrvc) Update(ctx context.Context, p *books.UpdatePayload) (res string, err error) {
	//Create an instance of models.Book
	bookToUpdate := models.Book{}

	//Find the book from database that matches with bookToUpdate
	database.DB.First(&bookToUpdate)
	result := database.DB.First(&bookToUpdate, p.ID)

	if result.Error != nil {
		return "Failed", books.MakeNotFound(result.Error)
	}

	//Check if the payload contains the Title and update
	if p.Title != nil {
		bookToUpdate.Title = *p.Title
	}

	//Check if the payload contains the Author and update
	if p.Author != nil {
		bookToUpdate.Author = *p.Author
	}

	//Check if the payload contains the PublishedAt and update
	if p.PublishedAt != nil {
		//Parse the date from String(yyyy-mm-dd) to time.Time
		publishedAtDate, _ := time.Parse("2006-01-02", *p.PublishedAt)
		bookToUpdate.PublishedAt = publishedAtDate
	}

	//Check if the payload contains the BookCover and update
	if p.BookCover != nil {
		//Split the Book Cover image into Bytes,Type and Name
		bookToUpdate.BookCoverBytes = p.BookCover.Bytes
		bookToUpdate.CoverFileType = *p.BookCover.Type
		bookToUpdate.CoverFileName = *p.BookCover.Name
	}

	//Save the result on database
	result = database.DB.Save(&bookToUpdate)

	//Handle any error
	if result.Error != nil {
		return "Failed", result.Error
	}

	return "Updated", nil
}

// Delete a book.
func (s *bookssrvc) Delete(ctx context.Context, p *books.DeletePayload) (res string, err error) {
	bookToDelete := models.Book{}

	//Check if the record with ID = p.ID exists
	result := database.DB.First(&bookToDelete, p.ID)
	if result.Error != nil {
		return "Book with ID: %v does not Exist", books.MakeNotFound(result.Error)
	}
	//Delete from the database where the ID = p.ID
	result = database.DB.Delete(&bookToDelete)

	//Handle any error
	if result.Error != nil {
		return "Failed", result.Error
	}

	return fmt.Sprintf("Book with ID: %v deleted", *p.ID), nil
}

// Convert from books.Book to models.Book object
func toBook(book models.Book) *books.Book {

	newRes := &books.Book{}
	newRes.ID = book.Id
	newRes.Title = book.Title
	newRes.Author = book.Author
	newRes.BookCoverBytes = book.BookCoverBytes
	newRes.BookCoverFileName = &book.CoverFileName
	newRes.BookCoverFileType = &book.CoverFileType
	newRes.PublishedAt = book.PublishedAt.String()

	return newRes

}
