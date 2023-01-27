package database

import (
	"LibraryAPI/data"
	"context"
	"database/sql"
	"fmt"
)

type Database struct {
	SqlDb *sql.DB
}

var dbContext = context.Background()

func (db Database) CreateBook(book *data.Book) (newID int, err error) {

	err = db.SqlDb.PingContext(dbContext)
	if err != nil {
		return -1, err
	}

	queryStatement := `
    INSERT INTO Books(title, description, author, year ) VALUES (@Title, @Description, @Author, @Year);
   `

	query, err := db.SqlDb.Prepare(queryStatement)
	if err != nil {
		return -1, err
	}

	defer query.Close()

	newRecord := query.QueryRowContext(dbContext,
		sql.Named("Title", &book.Title),
		sql.Named("Description", &book.Description),
		sql.Named("Author", &book.Author),
		sql.Named("Year", &book.Year),
	)

	err = newRecord.Scan(&newID)
	if err != nil {
		return -1, err
	}

	return newID, nil
}

func (db Database) UpdateBook(book *data.Book) error {

	err := db.SqlDb.PingContext(dbContext)
	if err != nil {
		return err
	}

	queryStatement := `
    UPDATE Books 
	SET title = @Title, description = @Description, author = @Author, year = @Year,
	WHERE id = @ID;
   `

	query, err := db.SqlDb.Prepare(queryStatement)
	if err != nil {
		return err
	}

	defer query.Close()

	updatedRecord := query.QueryRowContext(dbContext,
		sql.Named("Title", &book.Title),
		sql.Named("Description", &book.Description),
		sql.Named("Author", &book.Author),
		sql.Named("Year", &book.Year),
		sql.Named("ID", &book.ID),
	)

	err = updatedRecord.Scan()
	if err != nil {
		return err
	}

	return nil
}

func (db Database) RetrieveBooks() (books []data.Book, err error) {
	err = db.SqlDb.PingContext(dbContext)
	if err != nil {
		return
	}

	sqlStatement := fmt.Sprintf("SELECT id, title, description, author, year FROM Books ORDER BY title ASC;")

	rows, queryErr := db.SqlDb.QueryContext(dbContext, sqlStatement)
	if queryErr != nil {
		return books, queryErr
	}

	var book data.Book
	for rows.Next() {
		nErr := rows.Scan(&book.ID, &book.Title, &book.Description, &book.Author, &book.Year)
		if nErr != nil {
			return books, nErr
		}

		books = append(books, book)
	}

	return
}

func (db Database) RetrieveBook(id string) (book data.Book, err error) {
	err = db.SqlDb.PingContext(dbContext)
	if err != nil {
		return
	}

	sqlStatement := fmt.Sprintf("SELECT id, title, description, author, year FROM Books WHERE id = ?;")

	rows, queryErr := db.SqlDb.QueryContext(dbContext, sqlStatement, id)
	if queryErr != nil {
		return book, queryErr
	}

	for rows.Next() {
		nErr := rows.Scan(&book.ID, &book.Title, &book.Description, &book.Author, &book.Year)
		if nErr != nil {
			return book, nErr
		}

		return
	}

	return
}

func (db Database) DeleteBook(id string) error {
	var err error

	err = db.SqlDb.PingContext(dbContext)
	if err != nil {
		fmt.Printf("Error checking db connection: %v", err)
	}

	queryStatement := `DELETE FROM Books WHERE id = @ID;`

	_, err = db.SqlDb.ExecContext(dbContext, queryStatement, sql.Named("ID", id))
	if err != nil {
		return err
	}

	return nil
}
