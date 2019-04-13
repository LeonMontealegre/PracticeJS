package handlers

import (
    // "regexp"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "path/filepath"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strings"
)

type Problem struct {
    Id       string `json:"id"`
    Complete bool   `json:"complete"`
}

type ProblemSet struct {
    Problems []Problem `json:"problems"`
}

// var validPath = regexp.MustCompile("^/(section)/([a-zA-Z0-9]+)$")

var problemSets map[string]ProblemSet = make(map[string]ProblemSet)

// func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
//     m := validPath.FindStringSubmatch(r.URL.Path)
//     if m == nil {
//         http.NotFound(w, r)
//         return "", errors.New("Invalid Page Title")
//     }
//     return m[2], nil // The title is the second subexpression.
// }

func init() {
    // Load sections from JSON files
    err := filepath.Walk("../data/sections/", func(path string, info os.FileInfo, err error) error {
        if info.IsDir() && path[len(path)-1] != '/' {
            var problemSet ProblemSet

            file, err := ioutil.ReadFile(path + "/problems.json")
            if err != nil {
                log.Printf("File error %s: %v\n", path, err)
                os.Exit(1)
            }

            err = json.Unmarshal(file, &problemSet)
            if err != nil {
                log.Printf("Failed to unmarshall json %s: %v\n", path, err)
                os.Exit(1)
            }

            split := strings.Split(path, "/")
            section := split[len(split)-1]

            problemSets[section] = problemSet
        }

        return nil
    })
    if err != nil {
        log.Printf("Failed to walk through directory: %v\n", err)
        os.Exit(1)
    }
}

func SectionHandler(c *gin.Context) {
    if val, ok := problemSets[c.Param("section")]; ok {
        c.HTML(http.StatusOK, "section.tmpl", val)
    } else {
        c.HTML(http.StatusNotFound, "error.tmpl", gin.H{"message": "Invalid session state."})
    }
}
