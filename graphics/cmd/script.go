package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func generateGoVarName(filePath string, baseDir string) string {
    // Normalize the file path and base directory to use forward slashes
    normalizedPath := filepath.ToSlash(filePath)
    normalizedBaseDir := filepath.ToSlash(baseDir)

    // Remove the base directory from the file path
    relativePath := strings.TrimPrefix(normalizedPath, normalizedBaseDir+"/")

    // Replace all slashes (/) with underscores (_)
    varName := strings.ReplaceAll(relativePath, "/", "_")

    // Remove the file extension
    extension := filepath.Ext(varName)
    varName = strings.TrimSuffix(varName, extension)

    return varName
}


func traverseDirectory(rootDir string) ([]string, error) {
    var filePaths []string
    err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            // Skip directories or files that cannot be accessed
            fmt.Printf("Skipping %s: %v\n", path, err)
            return nil
        }
        if !info.IsDir() {
            filePaths = append(filePaths, path)
        }
        return nil
    })
    return filePaths, err
}

func writeGoFile(filePaths []string, outputFile string, baseDir string) error {
    // Debug: Print the output file path
    fmt.Println("Creating output file:", outputFile)

    // Create the output file
    file, err := os.Create(outputFile)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return err
    }
    defer file.Close()

    // Write the package declaration and imports
    _, err = file.WriteString(`package resources

import (
    _ "embed"
)

`)
    if err != nil {
        fmt.Println("Error writing package declaration:", err)
        return err
    }

    // Write the variable declaration
    _, err = file.WriteString("var (\n")
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return err
    }

    // Write each file path as an embedded variable
    for _, filePath := range filePaths {
        varName := generateGoVarName(filePath, baseDir)

        _, err := file.WriteString(fmt.Sprintf("\t//go:embed %s\n\t%s []byte\n\n", filePath, varName))
        if err != nil {
            fmt.Println("Error writing file path to file:", err)
            return err
        }
    }

    // Close the variable declaration
    _, err = file.WriteString(")\n")
    if err != nil {
        fmt.Println("Error closing variable declaration:", err)
        return err
    }

    fmt.Println("File written successfully:", outputFile)
    return nil
}

func main() {
    // Hardcoded root directory
    rootDirectory := "C:\\Go_projects\\ebiten_pokemon\\graphics"
    outputFile := filepath.Join(rootDirectory, "assets.go")

    // Debug: Print the root directory
    fmt.Println("Traversing directory:", rootDirectory)

    // Traverse the directory
    filePaths, err := traverseDirectory(rootDirectory)
    if err != nil {
        fmt.Println("Error traversing directory:", err)
        return
    }

    if len(filePaths) == 0 {
        fmt.Println("No files found in the directory.")
        return
    }

    // Write the output file
    err = writeGoFile(filePaths, outputFile, rootDirectory)
    if err != nil {
        fmt.Println("Error writing Go file:", err)
    }
}