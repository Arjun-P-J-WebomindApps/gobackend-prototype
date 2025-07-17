package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Input struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	Role     int    `json:"role"`
	IsActive bool   `json:"isActive"`
}

type Variables struct {
	Input Input `json:"input"`
}

type GraphQlRequest struct {
	Query     string    `json:"query"`
	Variables Variables `json:"variables"`
}

func extractInt(data string) int {
	i, err := strconv.Atoi(data)
	if err != nil {
		log.Fatal("Error converting string to int", err)
	}

	return i
}

func extractExcelData(path string) [][]string {

	file, err := os.Open(path)

	if err != nil {
		log.Fatal("Cant process file at ", path)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	var table [][]string

	for {
		record, err := reader.Read()

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			log.Fatal("Record extraction error")
		}
		table = append(table, record)
	}
	return table
}

func postData(url string, jsonData []byte, wg *sync.WaitGroup) {

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Print("Error while uploading data", err.Error())
		return
	}

	defer func() {
		if resp.Body != nil {
			resp.Body.Close()

		}
	}()
	fmt.Println("status", resp.Status)
}

func replaceWithNameIfEmpty(username string, fallback string) string {
	if username == "" {
		return fallback
	}
	return strings.ToLower(username)

}

func UploadData(filename string, graphqlReq func(p []string) GraphQlRequest) {
	var wg sync.WaitGroup
	url := "http://localhost:8000/query"

	payloads := extractExcelData("C:/ProgramData/MySQL/MySQL Server 8.0/Uploads/" + filename + ".csv")
	sem := make(chan struct{}, 2)

	for _, p := range payloads {
		wg.Add(1)
		sem <- struct{}{}
		go func(p []string) {
			defer wg.Done()
			defer func() {
				<-sem
			}()
			graphqlData := graphqlReq(p)
			jsonData, err := json.Marshal(graphqlData)
			// fmt.Println(string(jsonData))
			if err != nil {
				log.Fatal("Couldnt add data", jsonData)
			}
			postData(url, jsonData, &wg)
		}(p)
	}

	wg.Wait()
}

func main() {
	createUsers := func(p []string) GraphQlRequest {
		user := GraphQlRequest{
			Query: "mutation CreateUser($input:CreateUserInput!) {createUser(input:$input) { name username email password mobile role isActive }}",
			Variables: Variables{
				Input: Input{
					Name:     p[1],
					Username: replaceWithNameIfEmpty("", p[1]),
					Email:    p[5],
					Password: p[6],
					Mobile:   p[3],
					Role:     extractInt(p[7]),
					IsActive: extractInt(p[8]) == 1,
				},
			},
		}
		return user
	}

	UploadData("muser_master", createUsers)
}
