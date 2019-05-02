package handlers

import (
    // "regexp"
    // "encoding/json"
    "github.com/gin-gonic/gin"
    // "path/filepath"
    // "io/ioutil"
    // "log"
    "net/http"
    // "os"
    // "strings"
)

type Test struct {
    TestCode string `json:"test"`
    Expected string `json:"expected"`
}

type Problem struct {
    Id          string   `json:"id"`
    Description string   `json:"description"`
    Examples    []string `json:"examples"`
    Tests       []Test   `json:"tests"`
    Answer      []string `json:"answer"`
}

var problems map[string]Problem = make(map[string]Problem)

func init() {
    // Load sections from JSON files
    // err := filepath.Walk("../data/sections/", func(path string, info os.FileInfo, err error) error {
    //     log.Printf("%s", path);
    //     if info.IsDir() && path[len(path)-1] != '/' {
    //         var problemSet ProblemSet
    //
    //         file, err := ioutil.ReadFile(path + "/problems.json")
    //         if err != nil {
    //             log.Printf("File error %s: %v\n", path, err)
    //             os.Exit(1)
    //         }
    //
    //         err = json.Unmarshal(file, &problemSet)
    //         if err != nil {
    //             log.Printf("Failed to unmarshall json %s: %v\n", path, err)
    //             os.Exit(1)
    //         }
    //
    //         split := strings.Split(path, "/")
    //         section := split[len(split)-1]
    //
    //         problemSets[section] = problemSet
    //     }
    //
    //     return nil
    // })
    // if err != nil {
    //     log.Printf("Failed to walk through directory: %v\n", err)
    //     os.Exit(1)
    // }
}

func ProblemHandler(c *gin.Context) {
    if val, ok := problemSets[c.Param("section")]; ok {
        c.HTML(http.StatusOK, "section.tmpl", val)
    } else {
        c.HTML(http.StatusNotFound, "error.tmpl", gin.H{"message": "Invalid session state."})
    }
}
