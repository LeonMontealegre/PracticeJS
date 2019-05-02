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

type ProblemSet struct {
    Problems  []ProblemDescription `json:"problems"`
    SectionID string
}

var problemSets map[string]ProblemSet = make(map[string]ProblemSet)
var problems map[string]Problem = make(map[string]Problem)

func handleError(err error, msg string) {
    if err != nil {
        log.Println(msg)
        log.Println(err)
        os.Exit(1)
    }
}

func init() {
    const PATH = "../data/sections/"

    // Get each section directory
    sections, err := ioutil.ReadDir(PATH)
    handleError(err, fmt.Sprintf("Failed to load sections directory %s\n", PATH))

    // Load each section from JSON files
    for _, info := range sections {
        if info.IsDir() {
            var path = PATH + info.Name()
            var problemSet ProblemSet

            // Load section
            file, err := ioutil.ReadFile(path + "/problems.json")
            handleError(err, fmt.Sprintf("File error %s\n", path))
            err = json.Unmarshal(file, &problemSet)
            handleError(err, fmt.Sprintf("Failed to unmarshall section json %s\n", path))

            section := info.Name()
            problemSet.SectionID = section;
            problemSets[section] = problemSet

            // Get each problem
            problemFiles, err := ioutil.ReadDir(path + "/problems/")
            handleError(err, fmt.Sprintf("Could not open directory '%s'\n", path))

            // Load each problem from JSON files
            for _, info := range problemFiles {
                var problem Problem

                // Load problem
                file, err := ioutil.ReadFile(path + "/problems/" + info.Name())
                handleError(err, fmt.Sprintf("File error %s\n", path))
                err = json.Unmarshal(file, &problem)
                handleError(err, fmt.Sprintf("Failed to unmarshall section json %s\n", path))


                problemName := info.Name()
                problemName = problemName[0:strings.LastIndex(problemName, ".")]
                log.Println("loaded problem: " + section + "/" + problemName)
                problems[section + "/" + problemName] = problem
            }
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

func ProblemHandler(c *gin.Context) {
    log.Println(c.Param("section")+"/"+c.Param("problem"))
    if val, ok := problems[c.Param("section")+"/"+c.Param("problem")]; ok {
        c.HTML(http.StatusOK, "exercise.tmpl", val)
    } else {
        c.HTML(http.StatusNotFound, "error.tmpl", gin.H{"message": "Invalid session state."})
    }
}
