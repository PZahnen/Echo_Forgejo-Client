package main

import (
	"encoding/base64"
	"fmt"
	"log"

	forgejo "codeberg.org/mvdkleijn/forgejo-sdk/forgejo/v2"
)

func main() {
	// Forgejo/Gitea-Client konfigurieren
	client, err := forgejo.NewClient("https://gitea.com", forgejo.SetToken("39b183885374876f47d45cfca2e73f61b4a82457"))
	if err != nil {
		log.Fatalf("Error creating Forgejo client: %v\n", err)
	}


	// Repository-Informationen
	owner := "PZahnen"
	repo := "Forgejo-SDK"
	branch := "main"

	// Datei erstellen
	createFile(client, owner, repo, branch, "example.txt", "Hello from Go! 🎉")
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