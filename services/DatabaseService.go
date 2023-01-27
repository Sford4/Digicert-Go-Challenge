package services

func (d *DatabaseService) DBGetAllBooksDB() (books []data.Book, err error) {
	books = make([]data.Book, 0)

	query := `
	SELECT
		id,
		title,
		description,
		author,
		year
	FROM Book
	ORDER BY title ASC`

	rows, err := d.readDB.Query(query)

	if err != nil {
		err = fmt.Errorf("readDB.Query: %w", err)
		return
	}

	for rows.Next() {
		var book data.Book

		err = rows.Scan(
			&.BookID,
			&.HouseID,
			&.ProviderID,
			&.ProviderBookID,
			&.ProviderBookToken,
			&.AccountNumberSuffix,
			&.AccountType,
			&.Creditable,
			&.Debitable,
			&.Enabled,
			&.DefaultAccount,
			&.HolderName,
			&.RoutingNumber,
			&.Status,
			&.NickName)

		if err != nil {
			err = fmt.Errorf("rows.Scan: %w", err)
			return
		}

		books = append(books, book)
	}

	return
}