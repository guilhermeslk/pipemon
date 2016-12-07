package main

import (
    "fmt"
    "strings"
    "time"
    "os"
    "os/exec"
    "database/sql"
    "encoding/json"
    "github.com/fatih/color"
    _ "github.com/lib/pq"
)

var db_jerico *sql.DB
var db_jis *sql.DB
var db_jps *sql.DB

func main() {
    db_jerico, err := sql.Open("postgres", "user=postgres dbname=cloudification_development sslmode=disable")
    checkErr(err)

    db_jis, err = sql.Open("postgres", "user=postgres dbname=cloudification_jis_development sslmode=disable")
    checkErr(err)

    db_jps, err = sql.Open("postgres", "user=postgres dbname=cloudification_jps_development sslmode=disable")
    checkErr(err)

    for range time.Tick(time.Second *5) {
        go queryPipelines(db_jerico)
    }
}

func queryPipelines(db *sql.DB) {
    clearScr()

    timestamp := time.Now().Local()
    fmt.Println(timestamp.String())

    rows, err := db.Query("SELECT id, state from pipelines ORDER BY id")
    checkErr(err)

    for rows.Next() {
        var id int
        var state string

        err = rows.Scan(&id, &state)
        checkErr(err)

        fmt.Println("%v", id)
        queryPipelineSteps(id, db, "")
    }

}

func queryPipelineSteps(pipelineId int, db *sql.DB, padding string) {
    steps, err := db.Query("SELECT id, step_class, state, async_result from pipeline_steps WHERE pipeline_id = $1 ORDER BY id", pipelineId)

    for steps.Next() {
        var id int
        var step_class string
        var state string
        var async_result string
        var async_result_data map[string]interface{}

        steps.Scan(&id, &step_class, &state, &async_result)
        checkErr(err)

        fmt.Printf(padding)
        printId(id)
        printSeparator()
        printStep(step_class)
        printSeparator()
        printState(state)
        fmt.Printf("\n")

        json.Unmarshal([]byte(async_result), &async_result_data)
        checkErr(err)

        if async_result_data["external_pipeline_id"] != nil {
            var external_pipeline_id float64
            var external_pipeline_name string

            external_pipeline_id = async_result_data["external_pipeline_id"].(float64)
            external_pipeline_name = async_result_data["external_pipeline_name"].(string)

            if external_pipeline_name == "JIS" {
                queryPipelineSteps(int(external_pipeline_id), db_jis, "     |")
            } else if external_pipeline_name == "JPS" {
               queryPipelineSteps(int(external_pipeline_id), db_jps, "     |")
            }
        }
    }
}

func printId(id int) {
    white := color.New(color.FgWhite).PrintfFunc()
    white("%4v", id)
}

func printStep(step_class string) {
    blue := color.New(color.FgBlue).PrintfFunc()
    blue("%-65s", step_class)
}

func printSeparator() {
    fmt.Printf(" | ")
}

func printState(state string) {
    red := color.New(color.FgRed).PrintfFunc()
    green := color.New(color.FgGreen).PrintfFunc()
    cyan := color.New(color.FgCyan).PrintfFunc()

    if state == "running" {
        cyan("%v", strings.ToUpper(state))
    } else if state == "failed" {
        red("%v", strings.ToUpper(state))
    } else if state == "done" {
        green("%v", strings.ToUpper(state))
    } else {
        fmt.Printf("PENDING")
    }
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func clearScr() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}
