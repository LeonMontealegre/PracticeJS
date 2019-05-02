package handlers

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strings"
    "fmt"
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

type ProblemDescription struct {
    Id       string `json:"id"`
    Complete bool   `json:"complete"`
}

type Section struct {
    SectionID           string
    Problems            map[string]Problem
    ProblemDescriptions []ProblemDescription `json:"problems"`
}

var sections map[string]Section = make(map[string]Section)

func handleError(err error, msg string) {
    if err != nil {
        log.Println(msg)
        log.Println(err)
        os.Exit(1)
    }
}

func loadProblem(path string, sectionName string, problemName string) {
    var problem Problem

    // Load problem
    file, err := ioutil.ReadFile(path + "/problems/" + problemName + ".json")
    handleError(err, fmt.Sprintf("File error %s\n", path))
    err = json.Unmarshal(file, &problem)
    handleError(err, fmt.Sprintf("Failed to unmarshall section json %s\n", path))

    sections[sectionName].Problems[problemName] = problem
}

func loadSection(path string, sectionName string) {
    var section Section

    // Load section
    file, err := ioutil.ReadFile(path + "/problems.json")
    handleError(err, fmt.Sprintf("File error %s\n", path))
    err = json.Unmarshal(file, &section)
    handleError(err, fmt.Sprintf("Failed to unmarshall section json %s\n", path))

    section.SectionID = sectionName;
    section.Problems = make(map[string]Problem)
    sections[sectionName] = section
}

func init() {
    const PATH = "../data/sections/"

    // Get each section directory
    sectionFiles, err := ioutil.ReadDir(PATH)
    handleError(err, fmt.Sprintf("Failed to load sections directory %s\n", PATH))

    // Load each section from JSON files
    for _, info := range sectionFiles {
        if info.IsDir() {
            sectionName := info.Name()
            path := PATH + sectionName

            loadSection(path, sectionName)

            // Get each problem
            problemFiles, err := ioutil.ReadDir(path + "/problems/")
            handleError(err, fmt.Sprintf("Could not open directory '%s'\n", path))

            // Load each problem from JSON files
            for _, info := range problemFiles {
                problemName := info.Name()
                problemName = problemName[0:strings.LastIndex(problemName, ".")] // Remove extension
                loadProblem(path, sectionName, problemName)
            }
        }
    }
}

func SectionHandler(c *gin.Context) {
    if val, ok := sections[c.Param("section")]; ok {
        c.HTML(http.StatusOK, "section.tmpl", val)
    } else {
        c.HTML(http.StatusNotFound, "error.tmpl", gin.H{"message": "Invalid session state."})
    }
}

func ProblemHandler(c *gin.Context) {
    if val, ok := sections[c.Param("section")].Problems[c.Param("problem")]; ok {
        c.HTML(http.StatusOK, "problem.tmpl", val)
    } else {
        c.HTML(http.StatusNotFound, "error.tmpl", gin.H{"message": "Invalid session state."})
    }
}
