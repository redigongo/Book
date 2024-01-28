package bookapp

import (
	"bookapp/gen/books"
	"bytes"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"os"
	"strings"
)

// BooksCreateDecoderFunc implements the multipart decoder for service "books"
// endpoint "create". The decoder must populate the argument p after encoding.
func BooksCreateDecoderFunc(mr *multipart.Reader, p **books.CreatePayload) error {
	//Create a books.CreatePayload object
	res := books.CreatePayload{}

	//Iterate over each part of multipart request. (infinite loop with a break inside)
	for {
		//Retrieve the next part
		part, err := mr.NextPart()

		// Check if we reached the end of the request
		if err == io.EOF {
			break
		}

		//Handle other errors
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return err
		}

		// Parse the Content-Disposition header to get parameters
		_, params, err := mime.ParseMediaType(part.Header.Get("Content-Disposition"))
		if err != nil {
			continue
		}

		// Parse the Content-Type header to get the disposition type
		disposition, _, err := mime.ParseMediaType(part.Header.Get("Content-Type"))

		// Check if the part corresponds to the "book_cover" field
		if params["name"] == "book_cover" {
			// Check if there's an error or if the part contains an image
			if err != nil || strings.HasPrefix(disposition, "image/") {

				//Read the bytes of the part
				bytes, err := io.ReadAll(part)

				//Handle reading errors
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}

				//Get the name of the file from the parameters
				filename := params["filename"]

				//Create a File object to assign it to BookCover
				imageUpload := books.File{
					Type:  &disposition,
					Bytes: bytes,
					Name:  &filename,
				}
				res.BookCover = &imageUpload
			}
		}
		//Read the contents of the part into a buffer
		buf := new(bytes.Buffer)
		buf.ReadFrom(part)

		//Check if param name is equal to author
		if params["name"] == "author" {
			res.Author = buf.String()
		}

		//Check if param name is equal to title
		if params["name"] == "title" {
			res.Title = buf.String()
		}

		//Check if param name is equal to published_at
		if params["name"] == "published_at" {
			res.PublishedAt = buf.String()
		}

	}

	// Update the pointer to the UpdatePayload with the populated object
	*p = &res

	return nil
}

// BooksCreateEncoderFunc implements the multipart encoder for service "books"
// endpoint "create".
func BooksCreateEncoderFunc(mw *multipart.Writer, p *books.CreatePayload) error {
	// Add multipart request encoder logic here
	return nil
}

// BooksUpdateDecoderFunc implements the multipart decoder for service "books"
// endpoint "update". The decoder must populate the argument p after encoding.
func BooksUpdateDecoderFunc(mr *multipart.Reader, p **books.UpdatePayload) error {
	//Create a books.UpdatePayload object
	res := books.UpdatePayload{}

	//Iterate over each part of multipart request. (infinite loop with a break inside)
	for {
		//Retrieve the next part
		part, err := mr.NextPart()

		// Check if we reached the end of the request
		if err == io.EOF {
			break
		}

		//Handle other errors
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return err
		}

		// Parse the Content-Disposition header to get parameters
		_, params, err := mime.ParseMediaType(part.Header.Get("Content-Disposition"))
		if err != nil {
			continue
		}

		// Parse the Content-Type header to get the disposition type
		disposition, _, err := mime.ParseMediaType(part.Header.Get("Content-Type"))

		// Check if the part corresponds to the "book_cover" field
		if params["name"] == "book_cover" {
			// Check if there's an error or if the part contains an image
			if err != nil || strings.HasPrefix(disposition, "image/") {

				//Read the bytes of the part
				bytes, err := io.ReadAll(part)

				//Handle reading errors
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}

				//Get the name of the file from the parameters
				filename := params["filename"]

				//Create a File object to assign it to BookCover
				imageUpload := books.File{
					Type:  &disposition,
					Bytes: bytes,
					Name:  &filename,
				}
				res.BookCover = &imageUpload
			}
		}

		//Read the contents of the part into a buffer
		buf := new(bytes.Buffer)
		buf.ReadFrom(part)

		//Check if param name is equal to author
		if params["name"] == "author" {
			//Read the value of the buffer and assign it to Author
			author := buf.String()
			res.Author = &author
		}

		//Check if param name is equal to title
		if params["name"] == "title" {
			//Read the value of the buffer and assign it to Title
			title := buf.String()
			res.Title = &title
		}

		//Check if param name is equal to published_at
		if params["name"] == "published_at" {
			//Read the value of the buffer and assign it to PublishedAt
			publishedAt := buf.String()
			res.PublishedAt = &publishedAt
		}

	}

	// Update the pointer to the UpdatePayload with the populated object
	*p = &res

	return nil
}

// BooksUpdateEncoderFunc implements the multipart encoder for service "books"
// endpoint "update".
func BooksUpdateEncoderFunc(mw *multipart.Writer, p *books.UpdatePayload) error {
	return nil
}
