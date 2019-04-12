package handlers

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

type Section struct {
    Title    string `json:"title"`
    Progress uint   `json:"progress"`
    Colspan  uint
}

type Row struct {
    Colspan  uint
    Sections []Section `json:"sections"`
}

type ProblemSet struct {
    Rows []Row `json:"rows"`
}

var problemSet ProblemSet

func init() {
    // Load sections from JSON file
    file, err := ioutil.ReadFile("../data/problemsets.json")
    if err != nil {
        log.Printf("File error: %v\n", err)
        os.Exit(1)
    }
    err = json.Unmarshal(file, &problemSet)
    if err != nil {
        log.Printf("Failed to unmarshall json: %v\n", err)
        os.Exit(1)
    }

    // Calculate colspans
    var total uint = 1
    for _, row := range problemSet.Rows {
        var amt uint = uint(len(row.Sections))
        if total % amt != 0 {
            total *= amt
        }
    }

    // Set spans
    for i, row := range problemSet.Rows {
        var amt uint = uint(len(row.Sections))
        var span uint = total / amt

        problemSet.Rows[i].Colspan = total
        for j, _ := range row.Sections {
            problemSet.Rows[i].Sections[j].Colspan = span
        }
    }
}

func IndexHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", problemSet)
}
