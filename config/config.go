package config

import (
    jira "commit-helper/adapters"
    "commit-helper/models"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "runtime"

    "github.com/c-bata/go-prompt"
)

type Config struct {
    UserDetails models.UserDetails 
    CommitTypes []prompt.Suggest
}

var config Config
var configPath string

func InitConfig() Config {
    initConfigPath()
    config := loadConfig()
    if config.UserDetails.Username == "" {
        userDetails := jira.GetUserDetails()
        config.UserDetails = userDetails
        config.CommitTypes = getCommitTypes() 
        writeToConfigFile(config)
    }
    return config
} 

func initConfigPath() string {
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

    configPath = filepath.Join(configDir, "config.json")
    return configPath
}

func getCommitTypes() []prompt.Suggest {
    return []prompt.Suggest{
        {Text: "🆕 feat", Description: "A new feature"},
        {Text: "🔧 fix", Description: "A bug fix"},
        {Text: "📚 docs", Description: "Documentation only changes"},
        {Text: "💅 style", Description: "Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)"},
        {Text: "🚧 refactor", Description: "A code change that neither fixes a bug nor adds a feature"},
        {Text: "🏎️ perf", Description: "A code change that improves performance"},
        {Text: "🧪 test", Description: "Adding missing tests or correcting existing tests"},
        {Text: "🧰 chore", Description: "Changes to the build process or auxiliary tools and libraries such as documentation generation"},
    }
}


func loadConfig() Config {
    file, err := os.Open(configPath)
    if err != nil {
        return Config{}
    }
    defer file.Close()

    var configuration Config
    err = json.NewDecoder(file).Decode(&configuration)
    if err != nil {
        return Config{}
    }
    return configuration
}

func writeToConfigFile(config Config) {
    file, err := os.Create(configPath)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    fmt.Println("Writing configuration to config.json")

    err = encoder.Encode(config)
    if err != nil {
        fmt.Println("Error encoding JSON:", err)
        return
    }
    fmt.Println("Config written to config.json")
}

func GetConfig() Config {
    return config
}
