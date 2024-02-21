package storage

import (
	"commit-helper/data"
	"commit-helper/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var dataPath string

func InitStorage() string {
    var configDir string

    switch runtime.GOOS {
    case "windows":
        configDir = os.Getenv("APPDATA")
    default:
        configDir = os.Getenv("HOME")
    }

    configDir = filepath.Join(configDir, ".commit-helper")

    if err := os.MkdirAll(configDir, 0755); err != nil {
        log.Fatal(err)
    }

    dataPath = filepath.Join(configDir, "data.json")
    return dataPath
}

func WriteToStorage() {
    file, err := os.Create(dataPath)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    ticketData := data.GetData().TicketData

    err = encoder.Encode(ticketData)
    if err != nil {
        fmt.Println("Error encoding JSON:", err)
        return
    }
}

func LoadFromStorage() models.Storage {
	file, err := os.Open(dataPath)
	if err != nil {
		return models.Storage{}
	}
	defer file.Close()

    var storage models.Storage
	err = json.NewDecoder(file).Decode(&storage)
	if err != nil {
		return models.Storage{}
	}
    data.GetData().SetTicketData(storage)
	return storage
}
