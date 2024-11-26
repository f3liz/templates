package handler

import (
	"os"
	"strconv"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/open-runtimes/types-for-go/v4/openruntimes"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/databases"
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
		appwrite.WithProject("67462c7b000e193b99aa"),
		appwrite.WithKey("standard_17e9b2c441aa0f15dc1f186f0752a5555364196501843194a2e0ec43fa2c09f21bee4bde69919f0850ac935d2872bb2da8f6bb79412d1e62e46030a14c306f9c43bcae27985b36683732e3e1ef42e7ed1a1e25daaf3d03e2ea225cbc65987d7c7939934126ff8f48a564dc4c4130e21a34569dff04256f85246befde44f252ab"),
	)

type Response struct {
	Motto       string `json:"motto"`
	Learn       string `json:"learn"`
	Connect     string `json:"connect"`
	GetInspired string `json:"getInspired"`
}

// This Appwrite function will be executed every time your function is triggered
func Main(Context openruntimes.Context) openruntimes.Response {
	// You can use the Appwrite SDK to interact with other services
	// For this example, we're using the Users service
	client := appwrite.NewClient(
		appwrite.WithEndpoint(os.Getenv("APPWRITE_FUNCTION_API_ENDPOINT")),
		appwrite.WithProject(os.Getenv("APPWRITE_FUNCTION_PROJECT_ID")),
		appwrite.WithKey(Context.Req.Headers["x-appwrite-key"]),
	)
	users := appwrite.NewUsers(client)

	response, err := users.List()
	if err != nil {
		Context.Error("Could not list users: " + err.Error())
	} else {
		// Log messages and errors to the Appwrite Console
		// These logs won't be seen by your end users
		Context.Log("Total users: " + strconv.Itoa(response.Total))
	}

	// The req object contains the request data
	if Context.Req.Path == "/ping" {
		// Use res object to respond with text(), json(), or binary()
		// Don't forget to return a response!
		return Context.Res.Text("Pong")
	}

	return Context.Res.Json(Response{
		Motto:       "Build like a team of hundreds_",
		Learn:       "https://appwrite.io/docs",
		Connect:     "https://appwrite.io/discord",
		GetInspired: "https://builtwith.appwrite.io",
	})
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

