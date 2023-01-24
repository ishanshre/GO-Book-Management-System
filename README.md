# GO-Book-Management-System
A book management system API with Golang and MySql


# Project Structure

1. CMD (Folder)
    - main.go (go file)
2. PKG (Folder)
    - config (Folder)
        - app.go (go file used for connecting with the database)
    - controllers (Folder)
        - Handler function in go that deals with the request and responses
    - models (Folder)
        - book.go (consist of structs and models that is used by database)
    - routes (Folder)
        - handling routes in go
    - utils (Folder)
        - utils.go (marshelling and unmarshelling JSON)


# Routes

    | GetBooks |<--------------| /book/all |<---------------------| GET |
    | GetBook |<--------------| /book/{id} |<---------------------| GET |
    | CreateBook |<--------------| /book/ |<---------------------| POST |
    | UpdateBook |<--------------| /book/{id}/update |<---------------------| PUT |
    | DeleteBook |<--------------| /book/{id}/delete |<---------------------| DELETE |

# Pakages
    - GORM - It is an orm - go get "github.com/jinzhu/gorm"
    - MYSQL - Connect to mysql using GORM - go get "github.com/jinzhu/gorm/dialects/mysql"
    - Gorilla/mux - Routers - go get "github.com/gorilla/mux"