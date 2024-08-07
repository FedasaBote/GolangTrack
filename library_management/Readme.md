# Library Management System

## Table of Contents
1. [Introduction](#introduction)
2. [System Requirements](#system-requirements)
3. [Installation](#installation)
4. [Usage](#usage)
   - [Starting the Application](#starting-the-application)
   - [Command Line Interface](#command-line-interface)
5. [System Overview](#system-overview)
   - [Architecture](#architecture)
   - [Components](#components)
6. [Features](#features)
   - [Add Book](#add-book)
   - [Remove Book](#remove-book)
   - [Borrow Book](#borrow-book)
   - [Return Book](#return-book)
   - [List Available Books](#list-available-books)
   - [List Borrowed Books](#list-borrowed-books)
   - [Add Member](#add-member)
7. [Contributing](#contributing)
8. [Testing](#testing)
9. [FAQ](#faq)
10. [License](#license)
11. [Contact](#contact)

## Introduction
The Library Management System is a console-based application designed to manage a collection of books and members. It supports functionalities like adding and removing books, borrowing and returning books, and listing available and borrowed books.

## System Requirements
- Go version 1.18 or higher
- A terminal or command-line interface

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/library-management-system.git
2. Navigate to the repository:
    ```sh
    cd library-management-system
3. Insall dependencies:
    ```sh
    go mod download

## Usage
### Starting the application
To start the application, run:
    ```sh
    go run main.go

### Command Line Interface
The application provides a simple menu-driven interface. Users can select options by entering the corresponding number.

## System Overview
### Architecture
The system is built using the Model-View-Controller (MVC) architecture:
- **Model**: Contains the business logic, including structures for books and members.
- **Controller**: Manages the flow of data and handles user input.
- **View**: As this is a console application, the view is represented by the output and user prompts

### Components
- **models**: Contains definitions for `Book`, `Member`, and `Library`.
- **controllers**: Contains the `LibraryController` which manages interactions between the user and the models.

## Features
### Add Book
- Prompts the user to enter a book ID and title to add to the library.
### Remove Book
- Allows removal of a book by ID.
### Borrow Book
- Enables a member to borrow a book by entering their member ID and the book ID.
### List Available Books
- Lists all books that are not currently borrowed.
### List Borrowed Books
- Shows all books borrowed by a specific member.
### Add Member
- Adds a new member to the library system.
