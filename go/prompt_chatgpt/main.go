package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/databases"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/models"
)

var (
	appwriteClient    client.Client
	todoDatabase      *models.Database
	todoCollection    *models.Collection
	appwriteDatabases *databases.Databases
)

func main() {
	appwriteClient = appwrite.NewClient(
		appwrite.WithProject("674644eb00302038176d"), // project ID
		appwrite.WithKey("standard_e41e04fc9782068ba1f4b276025d11d2e3ebb36853595119b66ec2e2240f680a7294d1925201b5b008a36f9efa72a8a8ca2c24c4b7e01b480da5c6e13c9ba3f2db17c9ceb9aaa2fd2a6787540ad4742af5cfe45da6c11a63772c9d1f6f9c19e6ef45d9baac5993ea426a939ea291742b5f1b62318b6a42fef7aeed3f206960b7"), // Personal API key
	)

	prepareDatabase()
	seedDatabase()
	getTodos()
}

func prepareDatabase() {
	appwriteDatabases = appwrite.NewDatabases(appwriteClient)

	todoDatabase, _ = appwriteDatabases.Create(
		id.Unique(),
		"TodosDB",
	)

	todoCollection, _ = appwriteDatabases.CreateCollection(
		todoDatabase.Id,
		id.Unique(),
		"Todos",
	)

	appwriteDatabases.CreateStringAttribute(
		todoDatabase.Id,
		todoCollection.Id,
		"title",
		255,
		true,
	)

	appwriteDatabases.CreateStringAttribute(
		todoDatabase.Id,
		todoCollection.Id,
		"description",
		255,
		false,
	)

	appwriteDatabases.CreateBooleanAttribute(
		todoDatabase.Id,
		todoCollection.Id,
		"isComplete",
		true,
	)
}

func seedDatabase() {
	testTodo1 := map[string]interface{}{
		"title":       "Buy apples",
		"description": "At least 2KGs",
		"isComplete":  true,
	}

	testTodo2 := map[string]interface{}{
		"title":      "Wash the apples",
		"isComplete": true,
	}

	testTodo3 := map[string]interface{}{
		"title":       "Cut the apples",
		"description": "Don't forget to pack them in a box",
		"isComplete":  false,
	}

	appwriteDatabases.CreateDocument(
		todoDatabase.Id,
		todoCollection.Id,
		id.Unique(),
		testTodo1,
	)

	appwriteDatabases.CreateDocument(
		todoDatabase.Id,
		todoCollection.Id,
		id.Unique(),
		testTodo2,
	)

	appwriteDatabases.CreateDocument(
		todoDatabase.Id,
		todoCollection.Id,
		id.Unique(),
		testTodo3,
	)
}

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsComplete  bool   `json:"isComplete"`
}

type TodoList struct {
	*models.DocumentList
	Documents []Todo `json:"documents"`
}

func getTodos() {
	todoResponse, _ := appwriteDatabases.ListDocuments(
		todoDatabase.Id,
		todoCollection.Id,
	)

	var todos TodoList
	todoResponse.Decode(&todos)

	for _, todo := range todos.Documents {
		fmt.Printf("Title: %s\nDescription: %s\nIs Todo Complete: %t\n\n", todo.Title, todo.Description, todo.IsComplete)
	}
}
