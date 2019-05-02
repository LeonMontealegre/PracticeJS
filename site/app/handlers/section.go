package handlers

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strings"
)

type ProblemDescription struct {
    Id       string `json:"id"`
    Complete bool   `json:"complete"`
}

type ProblemSet struct {
    Problems []ProblemDescription `json:"problems"`
}

var problemSets map[string]ProblemSet = make(map[string]ProblemSet)

func init() {
    const PATH = "../data/sections/"

    // Get each section directory
    sections, err := ioutil.ReadDir(PATH)
    if err != nil {
        log.Printf("Could not open directory '%s': %v\n", PATH, err)
        os.Exit(1)
    }

    // Load each section from JSON files
    for _, file := range sections {
        if file.IsDir() {
            var path = PATH + file.Name()

            file, err := ioutil.ReadFile(path + "/problems.json")
            if err != nil {
                log.Printf("File error %s: %v\n", path, err)
                os.Exit(1)
            }
            
            var problemSet ProblemSet

            err = json.Unmarshal(file, &problemSet)
            if err != nil {
                log.Printf("Failed to unmarshall json %s: %v\n", path, err)
                os.Exit(1)
            }

            split := strings.Split(path, "/")
            section := split[len(split)-1]

            problemSets[section] = problemSet
        }
    }
}

func SectionHandler(c *gin.Context) {
    if val, ok := problemSets[c.Param("section")]; ok {
        c.HTML(http.StatusOK, "section.tmpl", val)
    } else {
        c.HTML(http.StatusNotFound, "error.tmpl", gin.H{"message": "Invalid session state."})
    }
}
