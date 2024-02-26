package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "time"
)

var env = os.Getenv("TARGET_ENV")
var branch = os.Getenv("BRANCH")

func init() {
    for _, i := range []string{"TARGET_ENV", "BRANCH"} {
        _, present := os.LookupEnv(i)
        if !present {
            fmt.Printf("%v variable was not found. Please set it and try again\n", i)
            os.Exit(1)
        }

        if env != "DEV" && env != "PROD" {
            fmt.Println("TARGET_ENV must be either DEV or PROD. Other values are not allowed.")
            os.Exit(1)
        }

        if branch == "" {
            fmt.Println("BRANCH can't be empty.")
            os.Exit(1)         
        }
    }
}

type Response struct {
    Description string    `json:"description"`
    Environment string    `json:"environment"`
    Time        time.Time `json:"time"`
}

func basicAPI(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    resp := Response{
        Description: fmt.Sprintf("You're viewing the %v api, deployed from branch %v", env, branch),
        Environment: env,
        Time:        time.Now(),
    }
    jsonResp, err := json.Marshal(&resp)
    if err != nil {
        w.Write([]byte("Error"))
        return
    }

    w.Write(jsonResp)
}

func main() {
    http.HandleFunc("/", basicAPI)

    var serverPort string
    if env == "DEV" {
        serverPort = ":4040"
    } else {
        serverPort = ":5050"
    }
    fmt.Printf("Starting server on port %v...", serverPort)
    http.ListenAndServe(serverPort, nil)
}
