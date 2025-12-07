package main

import (
	"fmt"
	"log"
	"os"
)

type Contact struct {
	Name      string `json:"name"`
	Lable     string `json:"lable"`
	ContactNo int    `json:"contactNo"`
}
type Contacts struct {
	Contacts []Contact `json:"contacts"`
}

func main() {

	fmt.Println("Welcome to Global Contact Manager CLI")
	contactsFile, err := os.Open("contacts.json")

	if err != nil {
		log.Fatal("Error opening file -> ", err)
	}
	fmt.Println("File Opened successfully")
	defer contactsFile.Close()

	// fmt.Println("Reading Structured data from a file using Structs:")

	// contactsFile, err := os.Open("contacts.json")

	// if err != nil {
	// 	log.Fatal("Error opening file: ", err)
	// }
	// fmt.Println("Opeed file successfully")

	// defer contactsFile.Close()

	// // read our opened jsonFile as a byte array.
	// byteArray, err := io.ReadAll(contactsFile) // reads the file and returns byte array from that
	// if err != nil {
	// 	log.Fatal("Error reading file: ", err)
	// }
	// fmt.Println("Byte Array: ", byteArray) // Returns those bytes in a []byte slice -> file on disk  →  [123, 34, 110, ...]

	// var contacts Contacts
	// json.Unmarshal(byteArray, &contacts)
	// /* Think of &contacts like handing someone the address of your house
	// so they can deliver furniture inside.

	// If you just give them a photo of your house (contacts),
	// they can’t actually put anything inside it.
	// */
	// fmt.Println(contacts)
	// // isRunning := true
	// for i := 0; i < len(contacts.Contacts); i++ {
	// 	fmt.Println("Name: ", contacts.Contacts[i].Name)
	// }

	// fmt.Println("Reading UnStructured data from a file using Structs:")

	// var contactsNew map[string]interface{}
	// json.Unmarshal([]byte(byteArray), &contactsNew)
	// fmt.Println("Contacts: ", contactsNew["contacts"])
	// for isRunning {
	// 	fmt.Println("1. Add Contact")
	// 	fmt.Println("2. Update Contact")
	// 	fmt.Println("3. Delete Contact")
	// 	fmt.Println("4. Exit App")

	// 	var operation int
	// 	fmt.Println("Select a option: ")
	// 	fmt.Scanln(&operation)
	// 	switch operation {
	// 	case 1:
	// 		fmt.Println("Contact Added successfully")
	// 	case 4:
	// 		isRunning = false
	// 		fmt.Println("Exited from app")
	// 	default:
	// 		fmt.Println("Invalid Operation. Select a valid option from below:")

	// 	}
	// }

}
