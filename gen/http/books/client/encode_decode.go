// Code generated by goa v3.14.1, DO NOT EDIT.
//
// books HTTP client encoders and decoders
//
// Command:
// $ goa gen bookapp/design

package client

import (
	books "bookapp/gen/books"
	booksviews "bookapp/gen/books/views"
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildFindAllRequest instantiates a HTTP request object with method and path
// set to call the "books" service "findAll" endpoint
func (c *Client) BuildFindAllRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: FindAllBooksPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("books", "findAll", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeFindAllResponse returns a decoder for responses returned by the books
// findAll endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeFindAllResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body BookCollection
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("books", "findAll", err)
			}
			p := NewFindAllBookCollectionOK(body)
			view := "default"
			vres := booksviews.BookCollection{Projected: p, View: view}
			if err = booksviews.ValidateBookCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("books", "findAll", err)
			}
			res := books.NewBookCollection(vres)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("books", "findAll", resp.StatusCode, string(body))
		}
	}
}

// BuildFindByIDRequest instantiates a HTTP request object with method and path
// set to call the "books" service "findById" endpoint
func (c *Client) BuildFindByIDRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*books.FindByIDPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("books", "findById", "*books.FindByIDPayload", v)
		}
		if p.ID != nil {
			id = *p.ID
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: FindByIDBooksPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("books", "findById", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeFindByIDResponse returns a decoder for responses returned by the books
// findById endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeFindByIDResponse may return the following errors:
//   - "NotFound" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeFindByIDResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body FindByIDResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("books", "findById", err)
			}
			p := NewFindByIDBookOK(&body)
			view := "default"
			vres := &booksviews.Book{Projected: p, View: view}
			if err = booksviews.ValidateBook(vres); err != nil {
				return nil, goahttp.ErrValidationError("books", "findById", err)
			}
			res := books.NewBook(vres)
			return res, nil
		case http.StatusNotFound:
			var (
				body FindByIDNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("books", "findById", err)
			}
			err = ValidateFindByIDNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("books", "findById", err)
			}
			return nil, NewFindByIDNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("books", "findById", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "books" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateBooksPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("books", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the books create
// server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*books.CreatePayload)
		if !ok {
			return goahttp.ErrInvalidType("books", "create", "*books.CreatePayload", v)
		}
		if err := encoder(req).Encode(p); err != nil {
			return goahttp.ErrEncodingError("books", "create", err)
		}
		return nil
	}
}

// NewBooksCreateEncoder returns an encoder to encode the multipart request for
// the "books" service "create" endpoint.
func NewBooksCreateEncoder(encoderFn BooksCreateEncoderFunc) func(r *http.Request) goahttp.Encoder {
	return func(r *http.Request) goahttp.Encoder {
		body := &bytes.Buffer{}
		mw := multipart.NewWriter(body)
		return goahttp.EncodingFunc(func(v any) error {
			p := v.(*books.CreatePayload)
			if err := encoderFn(mw, p); err != nil {
				return err
			}
			r.Body = io.NopCloser(body)
			r.Header.Set("Content-Type", mw.FormDataContentType())
			return mw.Close()
		})
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the books
// create endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("books", "create", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("books", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateRequest instantiates a HTTP request object with method and path
// set to call the "books" service "update" endpoint
func (c *Client) BuildUpdateRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*books.UpdatePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("books", "update", "*books.UpdatePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateBooksPath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("books", "update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRequest returns an encoder for requests sent to the books update
// server.
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*books.UpdatePayload)
		if !ok {
			return goahttp.ErrInvalidType("books", "update", "*books.UpdatePayload", v)
		}
		if err := encoder(req).Encode(p); err != nil {
			return goahttp.ErrEncodingError("books", "update", err)
		}
		return nil
	}
}

// NewBooksUpdateEncoder returns an encoder to encode the multipart request for
// the "books" service "update" endpoint.
func NewBooksUpdateEncoder(encoderFn BooksUpdateEncoderFunc) func(r *http.Request) goahttp.Encoder {
	return func(r *http.Request) goahttp.Encoder {
		body := &bytes.Buffer{}
		mw := multipart.NewWriter(body)
		return goahttp.EncodingFunc(func(v any) error {
			p := v.(*books.UpdatePayload)
			if err := encoderFn(mw, p); err != nil {
				return err
			}
			r.Body = io.NopCloser(body)
			r.Header.Set("Content-Type", mw.FormDataContentType())
			return mw.Close()
		})
	}
}

// DecodeUpdateResponse returns a decoder for responses returned by the books
// update endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeUpdateResponse may return the following errors:
//   - "NotFound" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("books", "update", err)
			}
			return body, nil
		case http.StatusNotFound:
			var (
				body UpdateNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("books", "update", err)
			}
			err = ValidateUpdateNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("books", "update", err)
			}
			return nil, NewUpdateNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("books", "update", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteRequest instantiates a HTTP request object with method and path
// set to call the "books" service "delete" endpoint
func (c *Client) BuildDeleteRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*books.DeletePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("books", "delete", "*books.DeletePayload", v)
		}
		if p.ID != nil {
			id = *p.ID
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteBooksPath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("books", "delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDeleteResponse returns a decoder for responses returned by the books
// delete endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeDeleteResponse may return the following errors:
//   - "NotFound" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeDeleteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("books", "delete", err)
			}
			return body, nil
		case http.StatusNotFound:
			var (
				body DeleteNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("books", "delete", err)
			}
			err = ValidateDeleteNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("books", "delete", err)
			}
			return nil, NewDeleteNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("books", "delete", resp.StatusCode, string(body))
		}
	}
}

// unmarshalBookToBooksviewsBookView builds a value of type
// *booksviews.BookView from a value of type *Book.
func unmarshalBookToBooksviewsBookView(v *Book) *booksviews.BookView {
	res := &booksviews.BookView{
		ID:                v.ID,
		Title:             v.Title,
		Author:            v.Author,
		BookCoverBytes:    v.BookCoverBytes,
		BookCoverFileName: v.BookCoverFileName,
		BookCoverFileType: v.BookCoverFileType,
		PublishedAt:       v.PublishedAt,
	}

	return res
}

// marshalBooksFileToFileRequestBody builds a value of type *FileRequestBody
// from a value of type *books.File.
func marshalBooksFileToFileRequestBody(v *books.File) *FileRequestBody {
	if v == nil {
		return nil
	}
	res := &FileRequestBody{
		Type:  v.Type,
		Bytes: v.Bytes,
		Name:  v.Name,
	}

	return res
}

// marshalFileRequestBodyToBooksFile builds a value of type *books.File from a
// value of type *FileRequestBody.
func marshalFileRequestBodyToBooksFile(v *FileRequestBody) *books.File {
	if v == nil {
		return nil
	}
	res := &books.File{
		Type:  v.Type,
		Bytes: v.Bytes,
		Name:  v.Name,
	}

	return res
}
