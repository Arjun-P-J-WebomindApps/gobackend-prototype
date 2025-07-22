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

type UserInput struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	Role     int    `json:"role"`
	IsActive bool   `json:"isActive"`
}

type CompanyInput struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type ModelInput struct {
	CompanyName string `json:"company_name"`
	Name        string `json:"name"`
}

type ModelVariantInput struct {
	ModelName  string `json:"model_name"`
	ModelType  string `json:"model_type"`
	ModelImage string `json:"model_image"`
}

type BrandInput struct {
	Name string `json:"name"`
}

type CategoryInput struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type ProductPartInput struct {
	CompanyName  string `json:"company_name"`
	ModelName    string `json:"model_name"`
	BrandName    string `json:"brand_name"`
	CategoryName string `json:"category_name"`
	PartNo       string `json:"part_no"`
	IsActive     bool   `json:"is_active"`
}

type Variables[T any] struct {
	Input T `json:"input"`
}

type GraphQlRequest[T any] struct {
	Query     string       `json:"query"`
	Variables Variables[T] `json:"variables"`
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

			log.Println("Record extraction error")
			continue
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

func UploadData[T any](filename string, graphqlReq func(p []string) GraphQlRequest[T]) {
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

func Auth() {
	createUsers := func(p []string) GraphQlRequest[UserInput] {
		user := GraphQlRequest[UserInput]{
			Query: "mutation CreateUser($input:CreateUserInput!) {createUser(input:$input) { name username email password mobile role isActive }}",
			Variables: Variables[UserInput]{
				Input: UserInput{
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

func Products() {
	createCompany := func(p []string) GraphQlRequest[CompanyInput] {
		company := GraphQlRequest[CompanyInput]{
			Query: "mutation CreateCompany($input: CreateCompanyInput!) { createCompany(input: $input) { id name status } }",
			Variables: Variables[CompanyInput]{
				Input: CompanyInput{
					Name:   p[1],
					Status: extractInt(p[2]) == 1,
				},
			},
		}

		return company
	}

	UploadData("company_master", createCompany)

	createModel := func(p []string) GraphQlRequest[ModelInput] {
		model := GraphQlRequest[ModelInput]{
			Query: "mutation CreateModel($input: CreateModelInput!) { createModel(input: $input) { id company_id name } }",
			Variables: Variables[ModelInput]{
				Input: ModelInput{
					Name:        p[1],
					CompanyName: p[3],
				},
			},
		}
		return model
	}

	UploadData("model_master", createModel)

	createModelVaraint := func(p []string) GraphQlRequest[ModelVariantInput] {
		modelVariant := GraphQlRequest[ModelVariantInput]{
			Query: "mutation CreateModelVariant($input: CreateModelVariantInput!) { createModelVariant(input: $input) { id model_id model_type model_image } }",
			Variables: Variables[ModelVariantInput]{
				Input: ModelVariantInput{
					ModelName:  p[1],
					ModelType:  p[4],
					ModelImage: p[5],
				},
			},
		}

		return modelVariant
	}

	UploadData("model_master", createModelVaraint)

	createBrand := func(p []string) GraphQlRequest[BrandInput] {
		brand := GraphQlRequest[BrandInput]{
			Query: "mutation CreateBrand($input: CreateBrandInput!) { createBrand(input: $input) { id name } }",
			Variables: Variables[BrandInput]{
				Input: BrandInput{
					Name: p[1],
				},
			},
		}

		return brand
	}

	UploadData("brand_master", createBrand)

	createCategory := func(p []string) GraphQlRequest[CategoryInput] {
		category := GraphQlRequest[CategoryInput]{
			Query: "mutation CreateCategory($input: CreateCategoryInput!) { createCategory(input: $input) { id name image } }",
			Variables: Variables[CategoryInput]{
				Input: CategoryInput{
					Name:  p[1],
					Image: p[8],
				},
			},
		}
		return category
	}

	UploadData("category_master", createCategory)

	createProductPart := func(p []string) GraphQlRequest[ProductPartInput] {
		productPart := GraphQlRequest[ProductPartInput]{
			Query: "mutation CreateProductPart($input: CreateProductPartInput!) { createProductPart(input: $input) { id company_id model_id brand_id category_id part_no is_active } }",
			Variables: Variables[ProductPartInput]{
				Input: ProductPartInput{
					CompanyName:  p[1],
					ModelName:    p[2],
					BrandName:    p[3],
					CategoryName: p[4],
					PartNo:       p[5],
					IsActive:     extractInt(p[34]) == 1,
				},
			},
		}

		return productPart
	}

	UploadData("products", createProductPart)
}

func main() {
	Auth()
	Products()
}
