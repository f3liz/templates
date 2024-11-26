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
		appwrite.WithProject("67464723000bb88c5fc1"),
		appwrite.WithKey("standard_16c2d86852e5945e107e4b6c9ebb5c357fe6000ab5aa63a1e2d94b88ca4b7b42e130ae0aaffd9e560113bd02fe8764fe0ff61ce4fb76238c8abaec6daeb47f6b78a20f501e23f9fa1934b505131c8670d07955469dedc2bd24a0a486b830a6a9e5628a1d649eb6f9ac4b340e473b06e82199b3e68dc67d32051fe0700ad26c53"),
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

