package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"pipemon/models"
	"strings"
	"time"
)

const (
	JIS_PIPELINE_NAME = "JIS"
	JPS_PIPELINE_NAME = "JPS"

	DB_USER        = "postgres"
	JERICO_DB_NAME = "cloudification_development"
	JIS_DB_NAME    = "cloudification_jis_development"
	JPS_DB_NAME    = "cloudification_jps_development"

	RELOAD_INTERVAL = 5
)

var (
	dbJerico *sql.DB
	dbJIS    *sql.DB
	dbJPS    *sql.DB
)

func main() {
	dbJerico = models.InitDB(fmt.Sprintf("user=%s dbname=%s sslmode=disable", DB_USER, JERICO_DB_NAME))
	dbJIS = models.InitDB(fmt.Sprintf("user=%s dbname=%s sslmode=disable", DB_USER, JIS_DB_NAME))
	dbJPS = models.InitDB(fmt.Sprintf("user=%s dbname=%s sslmode=disable", DB_USER, JPS_DB_NAME))

	listPipelines(dbJerico)
}

func listPipelines(db *sql.DB) {
	pipelines, err := models.QueryPipelines(db)
	checkErr(err)

	clearScr()
	fmt.Println(time.Now().Local().String())

	color.Yellow("### PIPEMON ###")

	for _, pipeline := range pipelines {
		printId(pipeline.Id)
		printSeparator()
		fmt.Printf(pipeline.Type)
		printSeparator()
		printState(pipeline.State)
		fmt.Printf("\n")
	}

	fmt.Print("Pipeline to query: ")

	var input int
	_, err = fmt.Scanf("%v\n", &input)
	checkErr(err)

	ticker := time.NewTicker(time.Second * RELOAD_INTERVAL)

	for range ticker.C {
		go showPipelineDetails(input, dbJerico)
	}
}

func showPipelineDetails(pipelineId int, db *sql.DB) {
	clearScr()

	fmt.Println(time.Now().Local().String())
	fmt.Printf("PIPELINE: %v", pipelineId)
	fmt.Printf("\n")

	listPipelineSteps(pipelineId, db, 0)
}

func listPipelineSteps(pipelineId int, db *sql.DB, paddingLength int) {
	steps, err := models.QueryPipelineSteps(pipelineId, db)
	checkErr(err)

	for _, step := range steps {
		printPipelineStep(step, paddingLength)

		json.Unmarshal([]byte(step.AsyncResult), &step.AsyncResultData)
		checkErr(err)

		if step.AsyncResultData["external_pipeline_id"] != nil {
			externalPipelineId := step.AsyncResultData["external_pipeline_id"].(float64)
			externalPipelineName := step.AsyncResultData["external_pipeline_name"].(string)

			if externalPipelineName == JIS_PIPELINE_NAME {
				listPipelineSteps(int(externalPipelineId), dbJIS, 5)
			} else if externalPipelineName == JPS_PIPELINE_NAME {
				listPipelineSteps(int(externalPipelineId), dbJPS, 5)
			}
		}
	}
}

func printPipelineStep(step *models.PipelineStep, paddingLength int) {
	printPadding(paddingLength)
	printId(step.Id)
	printSeparator()
	printStep(step.StepClass)
	printSeparator()
	printState(step.State)
	fmt.Printf("\n")
}

func printPadding(length int) {
	if length <= 0 {
		return
	}

	for i := 0; i < length; i++ {
		fmt.Printf(" ")
	}

	fmt.Printf("|")
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

func clearScr() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
