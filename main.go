package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

type TemplateData struct {
	NamespaceName    string
	CostCenterID     string
	ComputeType      string
	DevGroup         []string
	AdminGroup       []string
	DeploymentName   string
	AppLabel         string
	PolicyName       string
	AWSAccountNumber string
	RepoNameInECR    string
	TagPattern       string
	ReplicaCount     int
	ContainerName    string
	ContainerImage   string
	ContainerPort    int
	ServiceName      string
	ServicePort      int
	IngressName      string
	IngressHost      string
}

func main() {
	// Read JSON data from file
	jsonData, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Printf("Error reading JSON data file: %v", err)
	}

	var data TemplateData
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON data: %v", err)
	}

	// Parse template file
	funcMap := template.FuncMap{
		"toJsonArray": func(v []string) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}
	tmpl, err := template.New("template.yaml").Funcs(funcMap).ParseFiles("template.yaml")
	if err != nil {
		fmt.Printf("Error parsing template: %v", err)
	}

	// Create output file
	outFile, err := os.Create("output.yaml")
	if err != nil {
		fmt.Printf("Error creating output file: %v", err)
	}
	defer outFile.Close()

	// Execute template with data
	err = tmpl.Execute(outFile, data)
	if err != nil {
		fmt.Printf("Error executing template: %v", err)
	}

	fmt.Println("Template rendered successfully and written to output.yaml")
}
