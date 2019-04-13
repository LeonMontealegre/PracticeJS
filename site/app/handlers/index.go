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
    Id       string `json:"id"`
    Title    string `json:"title"`
    Progress uint   `json:"progress"`
    Colspan  uint
}

type SectionRow struct {
    Colspan  uint
    Sections []Section `json:"sections"`
}

type SectionSet struct {
    Rows []SectionRow `json:"rows"`
}

var sectionSet SectionSet

func init() {
    // Load sections from JSON file
    file, err := ioutil.ReadFile("../data/sections.json")
    if err != nil {
        log.Printf("File error: %v\n", err)
        os.Exit(1)
    }
    err = json.Unmarshal(file, &sectionSet)
    if err != nil {
        log.Printf("Failed to unmarshall json: %v\n", err)
        os.Exit(1)
    }

    // Calculate colspans
    var total uint = 1
    for _, row := range sectionSet.Rows {
        var amt uint = uint(len(row.Sections))
        if total % amt != 0 {
            total *= amt
        }
    }

    // Set spans
    for i, row := range sectionSet.Rows {
        var amt uint = uint(len(row.Sections))
        var span uint = total / amt

        sectionSet.Rows[i].Colspan = total
        for j, _ := range row.Sections {
            sectionSet.Rows[i].Sections[j].Colspan = span
        }
    }
}

func IndexHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", sectionSet)
}
