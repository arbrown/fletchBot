package settings

import (
    "os"
    "encoding/json"
    "log"
    "fmt"
)

// FletchBotSettings for settings in fletchbotsettings.json file
type FletchBotSettings struct {
    AppID, AppSecret string
    UserName, Password, UserAgent string
    
    MinAge, MaxAge, Interval int
    MaxPosts, MaxComments int
    
    CommentQuotes map[string]string
    
}

// Read in settings from a specified file
func ReadSettings(path string) (FletchBotSettings, error) {
    config := FletchBotSettings{}
    file, err := os.Open(path)

    if err != nil {
        return config, err
    }
    
    decoder  := json.NewDecoder(file)
    err = decoder.Decode(&config)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("%s\n%+v\n", "Settings:", config)
    return config, nil
}