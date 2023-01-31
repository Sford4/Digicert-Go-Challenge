# Digicert-Go-Challenge

# Instructions:

Using Go write a minimal web based REST API for a fictional public library that can perform the following functions:

    1. List all books in the library

    2. CRUD operations on a single book

The application does not have to be fully functional, but please take the time needed to exhibit your skills. There's no restrictions on what resources or third-party plugins you may use. We are looking for your best thoughts on application architecture, REST API design, maintainability, etc. Please return you submission as a link to a GitHub repo once completed. We look forward to seeing your results!

To Get It All Working:

** Based off of https://blog.logrocket.com/using-sql-database-golang/ **
Run the following commands in a terminal:

1. docker pull mcr.microsoft.com/mssql/server:2019-latest
2. sudo docker run -e "ACCEPT_EULA=Y" -e "SA_PASSWORD=password123" \  
   -p 1433:1433 --name sql1 -h sql1 \
   -d mcr.microsoft.com/mssql/server:2019-latest
3. sudo docker exec -it sql1 "bash"
4. /opt/mssql-tools/bin/sqlcmd -S localhost -U SA -P "password123"
5. CREATE DATABASE Library
6. GO
7. USE Library
8. CREATE TABLE Books ( id int IDENTITY(1, 1), title varchar(75), description varchar(255), author varchar(100), year int)

In the Go project, run:

1. go run main.go

In Postman, make calls to localhost:8000, such as:

1. Creating a book
   curl --location --request POST 'localhost:8000/book' \
   --header 'Content-Type: application/json' \
   --data-raw '{
   "title": "Harry Potter",
   "description": "A book about birds",
   "author": "JK Rowling",
   "year": 2010
   }'
2. Getting a single book
   curl --location --request GET 'localhost:8000/book/1'
3. Getting all books
   curl --location --request GET 'localhost:8000/books'
4. Updating a book
   curl --location --request PUT 'localhost:8000/book/1' \
   --header 'Content-Type: application/json' \
   --data-raw '{
   "id": 1,
   "title": "Harry Potter And the Long Title",
   "description": "A book about birds",
   "author": "JK Rowling",
   "year": 2010
   }'
5. Deleting a book
   curl --location --request DELETE 'localhost:8000/book/11'

Thanks for the challenge! Twas fun.
