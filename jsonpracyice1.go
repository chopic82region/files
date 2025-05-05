package main

import (
	"encoding/json"
	"fmt"
)

// пользователь библиотеки 
// json используется для передачи информации на другой сервис
type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	// слайс книг взятые пользователем
	Books []Book `json:"books"`
}

// книги 
type Book struct {
	BookID int    `json:"book_id"`
	Name   string `json:"name"`
}

// пустая структура 
// в будущем будет класться в канал для синхранизации 
type emptyStruct struct{}

func main() {

	// канал для будущей синхронизации
	ch := make(chan emptyStruct)

	varAndPeace := Book{
		BookID: 1231,
		Name:   "War and Peace",
	}

	math := Book{
		BookID: 1232,
		Name:   "Math",
	}

	peter := User{
		Name:  "Peter",
		Age:   18,
		Email: "Peter@bk.ru",
		Books: []Book{varAndPeace, math},
	}

	// горутина сериализирует информацию о пользователе
	go func() {

		newJson, err := json.Marshal(peter)
		if err != nil {
			panic(err)
		}

		fmt.Println("user information serialized: ", string(newJson))
		// пустая структура кладется в канал для синзранизации горутин
		ch <- emptyStruct{}
	}()
	
	// читатель канала
	// горутина main ждала выполнения горутины выше чтобы прочитать канал
	result1 := <-ch

	fmt.Println("sending information to librery....", result1)
	
	// горутина сообщает что данные о пользователе переданы в библиотеку
	go func() {
		ch <- emptyStruct{}
		fmt.Println("librery got information about user!")
	}()

	// снова main ждет записи в канал для чтения 
	result2 := <-ch
	fmt.Println("END...", result2)
}
