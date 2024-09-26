package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func Checkboxes(label string, opts []string) []string {
	res := []string{}
	prompt := &survey.MultiSelect{
		Message: label,
		Options: opts,
	}
	survey.AskOne(prompt, &res)

	return res
}

func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func GetPackage(i string) string {
	listPackage := map[string]string{
		"mysql":                      "gorm.io/driver/mysql",
		"fiber":                      "github.com/gofiber/fiber/v2",
		"snap-validator":             "github.com/apelweb15/snap-validator",
		"redis":                      "github.com/go-redis/redis/v8",
		"resty":                      "github.com/go-resty/resty/v2",
		"golang-jwt":                 "github.com/golang-jwt/jwt/v5",
		"google-uuid":                "github.com/google/uuid",
		"nats":                       "github.com/nats-io/nats.go",
		"cron":                       "github.com/robfig/cron/v3",
		"viper":                      "github.com/spf13/viper",
		"gorm":                       "gorm.io/gorm",
		"bayarind-signature-service": "repos.bayarind.id/common/go-signature-service",
		"bayarind-common-utilitites": "repos.bayarind.id/common/go-utilities",
		"validator":                  "github.com/go-playground/validator/v10",
	}

	return listPackage[i]
}

func createFile(dirpath, fileName string) {
	filePath := filepath.Join(dirpath, fileName)

	err := os.MkdirAll(dirpath, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

}

func createFileWithContent(dirpath, fileName, content string) {
	filePath := filepath.Join(dirpath, fileName)

	err := os.MkdirAll(dirpath, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func main() {
	projectName := StringPrompt("What is the name of the project?")

	selectedTemplate := Checkboxes(
		"Which project template do you want to use?",
		[]string{
			"Bayarind Pattern",
		},
	)

	selectedFramework := Checkboxes(
		"Which framework do you want to use?",
		[]string{
			"Fiber",
		},
	)

	selectedDatabase := Checkboxes(
		"Which database do you want to use?",
		[]string{
			"MySQL",
		},
	)

	selectedOtherLib := Checkboxes(
		"Is there any other library do you want to use?",
		[]string{
			"snap-validator",
			"validator",
			"redis",
			"resty",
			"golang-jwt",
			"google-uuid",
			"nats",
			"cron",
			"viper",
			"gorm",
			"bayarind-signature-service",
			"bayarind-common-utilities",
		},
	)

	allSelectedPackage := append(selectedOtherLib, selectedDatabase[0], selectedTemplate[0], selectedFramework[0])

	cmd := exec.Command("cmd", "/C", fmt.Sprintf("go mod init %s", projectName))
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Output:", string(output))
		return
	}

	for _, v := range allSelectedPackage {
		cmd := exec.Command("cmd", "/C", fmt.Sprintf("go get -u %s", GetPackage(v)))
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
			fmt.Println("Output:", string(output))
			return
		}
	}

	createFile(".", "Dockerfile")
	createFile(".", "config.json")
	createFile(".", "config.example.json")
	createFileWithContent(".", ".gitignore", "logs/* \n config.json")
	createFileWithContent("./cmd/rest/", "main.go", "package main")
	createFileWithContent("./internal/configs/", "app.go", "package config")
	createFileWithContent("./internal/constants/", "constants.go", "package constants")
	createFileWithContent("./internal/delivery/http/handler", "name_handler.go", "package handler")
	createFileWithContent("./internal/delivery/http/middleware", "name_middleware.go", "package middleware")
	createFileWithContent("./internal/delivery/http/route", "route.go", "package route")
	createFileWithContent("./internal/delivery/http/route", "route.go", "package route")
	createFileWithContent("./internal/entity/", "name_entity.go", "package entity")
	createFileWithContent("./internal/models/", "name_model.go", "package models")
	createFileWithContent("./internal/repository/", "name_repository.go", "package repository")
	createFileWithContent("./internal/repository/", "name_repository_impl.go", "package repository")
	createFileWithContent("./internal/usecase/", "name_usecase.go", "package usecase")
	createFileWithContent("./internal/usecase/", "name_usecase_impl.go", "package usecase")
	createFileWithContent("./internal/utils/", "name_util.go", "package utils")
}
