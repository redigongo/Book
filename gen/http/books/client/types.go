// Code generated by goa v3.14.1, DO NOT EDIT.
//
// books HTTP client types
//
// Command:
// $ goa gen bookapp/design

package client

import (
	books "bookapp/gen/books"
	booksviews "bookapp/gen/books/views"

	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "books" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// Title of the book
	Title string `form:"title" json:"title" xml:"title"`
	// Author of the book
	Author string `form:"author" json:"author" xml:"author"`
	// Book cover image
	BookCover *FileRequestBody `form:"book_cover,omitempty" json:"book_cover,omitempty" xml:"book_cover,omitempty"`
	// Published date of the book
	PublishedAt string `form:"published_at" json:"published_at" xml:"published_at"`
}

// UpdateRequestBody is the type of the "books" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	// Title of the book
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Author of the book
	Author *string `form:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// Book cover image
	BookCover *FileRequestBody `form:"book_cover,omitempty" json:"book_cover,omitempty" xml:"book_cover,omitempty"`
	// Published date of the book
	PublishedAt *string `form:"published_at,omitempty" json:"published_at,omitempty" xml:"published_at,omitempty"`
}

// BookCollection is the type of the "books" service "findAll" endpoint HTTP
// response body.
type BookCollection []*Book

// FindByIDResponseBody is the type of the "books" service "findById" endpoint
// HTTP response body.
type FindByIDResponseBody BookResponseBody

// FindByIDNotFoundResponseBody is the type of the "books" service "findById"
// endpoint HTTP response body for the "NotFound" error.
type FindByIDNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// UpdateNotFoundResponseBody is the type of the "books" service "update"
// endpoint HTTP response body for the "NotFound" error.
type UpdateNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// DeleteNotFoundResponseBody is the type of the "books" service "delete"
// endpoint HTTP response body for the "NotFound" error.
type DeleteNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// Book is used to define fields on response body types.
type Book struct {
	// ID of book
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Title of book
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Author of book
	Author *string `form:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// Image File base64
	BookCoverBytes []byte `form:"book_cover_bytes,omitempty" json:"book_cover_bytes,omitempty" xml:"book_cover_bytes,omitempty"`
	// Book cover file name
	BookCoverFileName *string `form:"book_cover_file_name,omitempty" json:"book_cover_file_name,omitempty" xml:"book_cover_file_name,omitempty"`
	// Book cover file type
	BookCoverFileType *string `form:"book_cover_file_type,omitempty" json:"book_cover_file_type,omitempty" xml:"book_cover_file_type,omitempty"`
	// Published date of the book
	PublishedAt *string `form:"published_at,omitempty" json:"published_at,omitempty" xml:"published_at,omitempty"`
}

// BookResponseBody is used to define fields on response body types.
type BookResponseBody struct {
	// ID of book
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Title of book
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Author of book
	Author *string `form:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// Image File base64
	BookCoverBytes []byte `form:"book_cover_bytes,omitempty" json:"book_cover_bytes,omitempty" xml:"book_cover_bytes,omitempty"`
	// Book cover file name
	BookCoverFileName *string `form:"book_cover_file_name,omitempty" json:"book_cover_file_name,omitempty" xml:"book_cover_file_name,omitempty"`
	// Book cover file type
	BookCoverFileType *string `form:"book_cover_file_type,omitempty" json:"book_cover_file_type,omitempty" xml:"book_cover_file_type,omitempty"`
	// Published date of the book
	PublishedAt *string `form:"published_at,omitempty" json:"published_at,omitempty" xml:"published_at,omitempty"`
}

// FileRequestBody is used to define fields on request body types.
type FileRequestBody struct {
	Type  *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	Bytes []byte  `form:"bytes,omitempty" json:"bytes,omitempty" xml:"bytes,omitempty"`
	Name  *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// NewCreateRequestBody builds the HTTP request body from the payload of the
// "create" endpoint of the "books" service.
func NewCreateRequestBody(p *books.CreatePayload) *CreateRequestBody {
	body := &CreateRequestBody{
		Title:       p.Title,
		Author:      p.Author,
		PublishedAt: p.PublishedAt,
	}
	if p.BookCover != nil {
		body.BookCover = marshalBooksFileToFileRequestBody(p.BookCover)
	}
	return body
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "books" service.
func NewUpdateRequestBody(p *books.UpdatePayload) *UpdateRequestBody {
	body := &UpdateRequestBody{
		Title:       p.Title,
		Author:      p.Author,
		PublishedAt: p.PublishedAt,
	}
	if p.BookCover != nil {
		body.BookCover = marshalBooksFileToFileRequestBody(p.BookCover)
	}
	return body
}

// NewFindAllBookCollectionOK builds a "books" service "findAll" endpoint
// result from a HTTP "OK" response.
func NewFindAllBookCollectionOK(body BookCollection) booksviews.BookCollectionView {
	v := make([]*booksviews.BookView, len(body))
	for i, val := range body {
		v[i] = unmarshalBookToBooksviewsBookView(val)
	}

	return v
}

// NewFindByIDBookOK builds a "books" service "findById" endpoint result from a
// HTTP "OK" response.
func NewFindByIDBookOK(body *FindByIDResponseBody) *booksviews.BookView {
	v := &booksviews.BookView{
		ID:                body.ID,
		Title:             body.Title,
		Author:            body.Author,
		BookCoverBytes:    body.BookCoverBytes,
		BookCoverFileName: body.BookCoverFileName,
		BookCoverFileType: body.BookCoverFileType,
		PublishedAt:       body.PublishedAt,
	}

	return v
}

// NewFindByIDNotFound builds a books service findById endpoint NotFound error.
func NewFindByIDNotFound(body *FindByIDNotFoundResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewUpdateNotFound builds a books service update endpoint NotFound error.
func NewUpdateNotFound(body *UpdateNotFoundResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewDeleteNotFound builds a books service delete endpoint NotFound error.
func NewDeleteNotFound(body *DeleteNotFoundResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateFindByIDNotFoundResponseBody runs the validations defined on
// findById_NotFound_response_body
func ValidateFindByIDNotFoundResponseBody(body *FindByIDNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateUpdateNotFoundResponseBody runs the validations defined on
// update_NotFound_response_body
func ValidateUpdateNotFoundResponseBody(body *UpdateNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateDeleteNotFoundResponseBody runs the validations defined on
// delete_NotFound_response_body
func ValidateDeleteNotFoundResponseBody(body *DeleteNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateBook runs the validations defined on Book
func ValidateBook(body *Book) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Author == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("author", "body"))
	}
	if body.PublishedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("published_at", "body"))
	}
	return
}

// ValidateBookResponseBody runs the validations defined on BookResponseBody
func ValidateBookResponseBody(body *BookResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Author == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("author", "body"))
	}
	if body.PublishedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("published_at", "body"))
	}
	return
}