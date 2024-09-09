package service

import (
	"context"
	"golang-crud/data/request"
	"golang-crud/data/response"
	"golang-crud/helper"
	"golang-crud/model"
	"golang-crud/repository"
)

type BookServiceImple struct {
	BookRepository repository.BookRepository
}

func NewBookServiceImple(bookrepository repository.BookRepository) BookService {
	return &BookServiceImple{BookRepository: bookrepository}
}

// Create implements BookService.
func (b *BookServiceImple) Create(ctx context.Context, request request.BookCreateRequest) {
	book := model.Book{
		Name: request.Name,
	}
	b.BookRepository.Save(ctx, book)
}

// Delete implements BookService.
func (b *BookServiceImple) Delete(ctx context.Context, bookId int) {
	book, err := b.BookRepository.FindById(ctx, bookId)
	helper.PanicIfError(err)
	b.BookRepository.Delete(ctx, book)
}

// FindAll implements BookService.
func (b *BookServiceImple) FindAll(ctx context.Context) []response.BookResponse {
	books := b.BookRepository.FindAll(ctx)

	var bookRes []response.BookResponse

	for _, value := range books {
		book := response.BookResponse{Id: value.Id, Name: value.Name}
		bookRes = append(bookRes, book)
	}
	return bookRes
}

// FindById implements BookService.
func (b *BookServiceImple) FindById(ctx context.Context, bookId int) response.BookResponse {
	book, err := b.BookRepository.FindById(ctx, bookId)
	helper.PanicIfError(err)

	return response.BookResponse(book)
}

// Update implements BookService.
func (b *BookServiceImple) Update(ctx context.Context, request request.BookUpdateRequest) {
	book, err := b.BookRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	book.Name = request.Name
	b.BookRepository.Update(ctx, book)
}
