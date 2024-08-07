package services

import (
	"library_management/models"
	"errors"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Books map[int]models.Book
	Members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books : make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

func (l *Library) AddMember(member models.Member) {
	l.Members[member.ID] = member
}

func (l *Library) AddBook(book models.Book) {
	l.Books[book.ID] = book;
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.Books, bookID)

	for memberId, member := range l.Members {
		for i, borrowedBook := range member.BorrowedBooks {
			if borrowedBook.ID == bookID {
				member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i + 1:]...)
				l.Members[memberId] = member
				break
			}
		}
	}
}

func (l *Library) BorrowBook(bookID, memberID int) error {
	book, ok := l.Books[bookID]
	if !ok {
		return errors.New("book not found")
	}

	member, ok := l.Members[memberID]
	if !ok {
		return errors.New("member not found")
	}

	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}

	book.Status = "Borrowed"
	l.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member
	return nil
}

func (l *Library) ReturnBook(bookID, memberID int) error {
	member, ok := l.Members[memberID]
	if !ok {
		return errors.New("member not found")
	}


	for i, borrowedBook := range member.BorrowedBooks {
		if borrowedBook.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i + 1: ]...)
			l.Members[memberID] = member
			book := l.Books[bookID]
			book.Status = "Available"
			l.Books[bookID] = book
			return nil
		}
	}

	return errors.New("book not borrowed by user")
}


func (l *Library) ListAvailableBooks() []models.Book {
	availableBooks := make([]models.Book, 0)

	for _, book := range l.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
			
		}
	}

	return availableBooks

}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, ok := l.Members[memberID]
	if !ok {
		return []models.Book{}
	}
	return member.BorrowedBooks
}