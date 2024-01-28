package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("books", func() {
	Title("Books Service")
	Description("HTTP service for managing books")
	Server("books", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
		})
	})
})
var _ = Service("books", func() {
	Description("Books Service")
	Error("NotFound")

	Method("findAll", func() {
		Description("FindAll Books")
		Result(CollectionOf(BookDTO))
		HTTP(func() {
			GET("/books")
			Response(StatusOK, func() {
				Body(CollectionOf(BookDTO))
			})
		})
	})

	Method("findById", func() {
		Description("Find book by id")
		Payload(func() {
			Field(1, "id", Int, "Book ID")
		})
		Result(BookDTO)
		HTTP(func() {
			GET("/books/{id}")
			Response(StatusOK, func() {
				Body(BookDTO)
			})
			Response("NotFound", StatusNotFound)

		})
	})

	Method("create", func() {
		Description("Create a new book.")
		Payload(func() {
			Attribute("title", String, "Title of the book")
			Attribute("author", String, "Author of the book")
			Attribute("book_cover", File, "Book cover image")
			Attribute("published_at", String, "Published date of the book", func() {
				Example("2024-01-27")
			})
			Required("title", "author", "published_at")
		})
		Result(String)

		HTTP(func() {
			POST("/books")

			MultipartRequest()
			Response(StatusCreated)
		})
	})

	Method("update", func() {
		Description("Update a book.")
		Payload(func() {
			Field(1, "id", Int, "Book ID")
			Attribute("title", String, "Title of the book")
			Attribute("author", String, "Author of the book")
			Attribute("book_cover", File, "Book cover image")
			Attribute("published_at", String, "Published date of the book", func() {
				Example("2023-12-31")
			})
			Required("id")
		})
		Result(String)

		HTTP(func() {
			PUT("/books/{id}")
			MultipartRequest()
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})

	Method("delete", func() {
		Description("Delete a book.")
		Payload(func() {
			Field(1, "id", Int, "Book ID")
		})
		Result(String)
		HTTP(func() {
			DELETE("/books/{id}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})

})

var File = Type("File", func() {
	Description("Image Upload Type")
	Attribute("type", String)
	Attribute("bytes", Bytes)
	Attribute("name", String)
})

var BookDTO = ResultType("application/vnd.book+json", func() {
	Description("A book payload")
	Attribute("id", Int, "ID of book")
	Attribute("title", String, "Title of book")
	Attribute("author", String, "Author of book")
	Attribute("book_cover_bytes", Bytes, "Image File base64")
	Attribute("book_cover_file_name", String, "Book cover file name")
	Attribute("book_cover_file_type", String, "Book cover file type")
	Attribute("published_at", String, "Published date of the book")
	Required("id", "title", "author", "published_at")
})
