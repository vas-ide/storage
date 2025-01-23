package main

import (
	"fmt"
	"log"
	
	"github.com/vas-ide/storage/internal/storage"

)




func main() {
	fmt.Println("OZON - Corsis ! GolanG !")
	
	

	storage :=storage.NewStorage()
	fmt.Println(storage)



	file, err := storage.Upload("Ozon - corsis start package structure", []byte("16.01.2025"))
	if err != nil {
		log.Fatal(err) 
	}
	fmt.Println(file)
	fmt.Println(file.Data)
	fmt.Printf("%s\n", string(file.Data))

	reqFile, err := storage.GetByID(file.ID)
	fmt.Println(reqFile)

}





