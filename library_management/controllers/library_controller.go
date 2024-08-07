package controllers

import (
	"fmt"
	"library_management/services"
	"library_management/models"
)

type LibraryController struct {
	Library *services.Library
}

func NewLibraryController() *LibraryController{
	return &LibraryController{
		Library: services.NewLibrary(),
	}
}


func (lc *LibraryController) AddBook() {
	var id int
	var title string
	var author string

	for {
		fmt.Print("Enter Book ID: ")
		_, err := fmt.Scan(&id)
		if err != nil {
			fmt.Println("Invalid Input. Please Enter a valid number")
			continue
		}
		break
	}

	for {
		fmt.Print("Enter Book title: ")
		_, err := fmt.Scan(&title)
		if err != nil {
			fmt.Println("Invalid Input. Please Enter a valid title.")
			continue
		}
		break
	}

	for {
		fmt.Print("Enter Book author: ")
		_, err := fmt.Scan(&author)
		if err != nil {
			fmt.Println("Invalid Input. Please Enter a valid author name.")
			continue
		}
		break
	}


	book := models.Book{ID: id, Title: title, Author: author, Status: "Available"}
	lc.Library.AddBook(book)
	fmt.Println("Book added Successfully!")
}

func (lc *LibraryController) RemoveBook() {
	var id int

	for {
		fmt.Print("Enter Book ID: ")
		_, err := fmt.Scan(&id)
		if err != nil {
			fmt.Println("Invalid Input. Please Enter valid book ID.")
			continue
		}
		break
	}

	lc.Library.RemoveBook(id)
	fmt.Println("Book has been removed successfully!")
}

func (lc *LibraryController) AddMember() {
	var memberId int

	for {
		fmt.Print("Enter member ID: ")
		_, err := fmt.Scan(&memberId)
		if err != nil {
			fmt.Println("Invalid number. Please insert valid id")
			continue
		}
		break
	}

	member := models.Member{ID: memberId}
	lc.Library.Members[memberId] = member
	fmt.Println("Member added successfully!")
}

func (lc *LibraryController) BorrowBook() {
	var bookID int
	var memberID int

	for {
		fmt.Print("Enter member ID: ")
		_, err := fmt.Scan(&memberID)
		if err != nil {
			fmt.Println("Invalid Input. Please enter a valid member id.")
			continue
		}
		break
	}

	for {
		fmt.Print("Enter book ID: ")
		_, err := fmt.Scan(&bookID)
		if err != nil {
			fmt.Println("Invalid Input. Please enter a valid book id.")
			continue
		}
		break
	}

	err := lc.Library.BorrowBook(bookID, memberID)

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("Book borrowed Successfully!")
	}

}

func (lc *LibraryController) ReturnBook() {
	var bookID int
	var memberID int

	for {
		fmt.Print("Enter member ID: ")
		_, err := fmt.Scan(&memberID)
		if err != nil {
			fmt.Println("Invalid Input. Please enter a valid member id.")
			continue
		}
		break
	}

	for {
		fmt.Print("Enter book ID: ")
		_, err := fmt.Scan(&bookID)
		if err != nil {
			fmt.Println("Invalid Input. Please enter a valid book id.")
			continue
		}
		break
	}

	err := lc.Library.ReturnBook(bookID, memberID)

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("Book Returned Successfully!")
	}
}

func (lc *LibraryController) ListAvailableBooks() {
	fmt.Println(lc.Library.ListAvailableBooks())
}

func (lc *LibraryController) ListBorrowedBooks() {
	var memberID int
	for {
		fmt.Print("Enter member ID: ")
		_, err := fmt.Scan(&memberID)

		if err != nil {
			fmt.Println("Invalid Input. Please enter a valid id")
			continue
		}
		break
	}
	fmt.Println(lc.Library.ListBorrowedBooks(memberID))
}

func (lc *LibraryController) ShowMenu() {
	fmt.Println("Library Management System")
	fmt.Println("1. Add Book")
	fmt.Println("2. Remove Book")
	fmt.Println("3. Borrow Book")
	fmt.Println("4. Return Book")
	fmt.Println("5. List Available Books")
	fmt.Println("6. List Borrowed Books")
	fmt.Println("7. Add Member")
	fmt.Println("8. Exit")
}

func (lc *LibraryController) Run() {
	for {
		lc.ShowMenu()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1: 
			lc.AddBook()
		case 2:
			lc.RemoveBook()
		case 3:
			lc.BorrowBook()
		case 4:
			lc.ReturnBook()
		case 5:
			lc.ListAvailableBooks()
		case 6:
			lc.ListBorrowedBooks()
		case 7:
			lc.AddMember()
		default:
			fmt.Println("Invalid choice. Please try again!")
		}
	}
}