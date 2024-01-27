package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/quotedprintable"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var fileContents []map[string]string
var fileCount int = 0

func creteJsonFile() {
	var alternatedContents []interface{}

	for _, element := range fileContents {
		alternatedContents = append(alternatedContents, map[string]interface{}{"index": map[string]string{"_index": "mails"}})
		alternatedContents = append(alternatedContents, element)
	}
	jsonData, err := json.MarshalIndent(alternatedContents, "", "    ")
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		os.Exit(1)
	}
	err = ioutil.WriteFile("tmp/dataEmails/output"+fmt.Sprint(fileCount)+".json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error al escribir el archivo:", err)
		os.Exit(1)
	}
}

func visitFile(path string, f os.FileInfo, err error) error {
	if err != nil {
		fmt.Println("Error al acceder a la ruta:", path, "El error es:", err)
		os.Exit(1)
		return err
	}
	if f.IsDir() {
		return nil
	}
	file, err := os.Open(path)
	defer file.Close()
	// data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		os.Exit(1)
		return err
	}
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1*1024*1024)
	mail := make(map[string]string)
	bodyData := ""
	isBodyData := false

	lastKey := ""
	for scanner.Scan() {
		line := scanner.Text()
		if !isBodyData {
			if strings.Contains(line, ":") {
				parts := strings.SplitN(line, ":", 2)
				if len(parts) == 2 {
					lastKey = parts[0]
					mail[lastKey] = strings.TrimSpace(parts[1])
				}
				if strings.HasPrefix(line, "X-FileName") {
					isBodyData = true
				}
			} else if lastKey != "" {
				mail[lastKey] += line
			}
		} else {
			bodyData += line + "\n"
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return err
	}
	reader := strings.NewReader(bodyData)
	qpReader := quotedprintable.NewReader(reader)
	decodedBytes, _ := ioutil.ReadAll(qpReader)
	decodedString := string(decodedBytes)
	mail["bodyData"] = decodedString
	fileContents = append(fileContents, mail)
	if len(fileContents) == 25000 {
		fmt.Println("Ruta:", path)
		fileCount++
		creteJsonFile()
		fileContents = fileContents[:0]
	}
	return nil
}

func main() {

	start := time.Now()
	if len(os.Args) < 2 {
		fmt.Println("Por favor, proporciona la ruta del directorio")
		os.Exit(1)
	}

	root := os.Args[1]
	err := filepath.Walk(root, visitFile)

	if len(fileContents) > 0 {
		fileCount++
		creteJsonFile()
	}

	if err != nil {
		fmt.Println("Error al recorrer el directorio:", err)
		os.Exit(1)
	}
	elapsed := time.Since(start)
	fmt.Printf("Tiempo de ejecución: %s\n", elapsed)
}
