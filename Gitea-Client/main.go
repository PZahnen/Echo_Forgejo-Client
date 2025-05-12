package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	forgejo "codeberg.org/mvdkleijn/forgejo-sdk/forgejo/v2"
	"github.com/joho/godotenv"
)

func main() {

    // .env-Datei laden
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v\n", err)
    }

    // Token aus Umgebungsvariablen abrufen
    token := os.Getenv("token")
    if token == "" {
        log.Fatalf("Token not found in environment variables")
    }

	// Forgejo/Gitea-Client konfigurieren
	client, err := forgejo.NewClient("https://gitea.com", forgejo.SetToken(token))
	if err != nil {
		log.Fatalf("Error creating Forgejo client: %v\n", err)
	}


	// Repository-Informationen
	owner := "PZahnen"
	repo := "Forgejo-SDK"
	branch := "main"

	// Datei erstellen
	createFile(client, owner, repo, branch, "example2.txt", "Hello from Go! 🎉")
}

func createFile(client *forgejo.Client, owner, repo, branch, filePath, content string) {

    // Inhalt base64-kodieren (erforderlich laut CreateFileOptions)
    encodedContent := base64.StdEncoding.EncodeToString([]byte(content))

    // Datei erstellen
    opts := forgejo.CreateFileOptions{
        FileOptions: forgejo.FileOptions{
            Message:    "Create example.txt",
            BranchName: branch,
        },
        Content: encodedContent,
    }

    // Datei erstellen
    _, _, err := client.CreateFile(owner, repo, filePath, opts)
	if err != nil {
		fmt.Printf("❌ Error creating file: %v\n", err)
		return
	}

	fmt.Println("✅ File created successfully!")
}