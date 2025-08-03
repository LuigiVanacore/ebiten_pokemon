package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
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
        varName = capitalizeFirst(varName) // Capitalize the first letter

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




func capitalizeFirst(s string) string {
    if s == "" {
        return s
    }
    r := []rune(s)
    r[0] = unicode.ToUpper(r[0])
    return string(r)
}

func generateVarNameAndValue(filePath, baseDir string, nameCount map[string]int) (string, string) {
    relPath := strings.TrimPrefix(filepath.ToSlash(filePath), filepath.ToSlash(baseDir)+"/")
    parts := strings.Split(relPath, "/")
    if len(parts) < 2 {
        return "", ""
    }
    dir := parts[0]
    file := parts[1]
    fileName := strings.TrimSuffix(file, filepath.Ext(file))

    // Check if fileName is numeric
    isNumeric, _ := regexp.MatchString(`^\d+$`, fileName)
    var varName, value string

    if isNumeric {
        // e.g. other/star_animation/01.png -> star_animation_01
        value = fmt.Sprintf("%s_%s", dir, fileName)
        varName = fmt.Sprintf("%s_%s", capitalizeFirst(dir), fileName)
    } else {
        // e.g. tilesets/water0.png -> Tilesets_water0
        value = fileName
        varName = fmt.Sprintf("%s_%s", capitalizeFirst(dir), fileName)
    }

    // Ensure unique variable names for repeated base names
    key := varName
    if count, exists := nameCount[key]; exists {
        nameCount[key] = count + 1
        varName = fmt.Sprintf("%s%d", varName, count)
        value = fmt.Sprintf("%s%d", fileName, count)
    } else {
        nameCount[key] = 1
    }

    return varName, value
}

func writeFileListGo(filePaths []string, outputFile string, baseDir string) error {
    fmt.Println("Creating file list:", outputFile)

    file, err := os.Create(outputFile)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return err
    }
    defer file.Close()

    _, err = file.WriteString("package resources\n\nvar (\n")
    if err != nil {
        return err
    }

nameCount := make(map[string]int)
for _, filePath := range filePaths {
    varName, value := generateVarNameAndValue(filePath, baseDir, nameCount)
    if varName == "" {
        continue
    }
    _, err := file.WriteString(fmt.Sprintf("\t%s string = \"%s\"\n", varName, value))
    if err != nil {
        return err
    }
}

    _, err = file.WriteString(")\n")
    return err
}

func main() {
     // Hardcoded root directory
    rootDirectory := "C:\\Go_projects\\ebiten_pokemon\\graphics"
    outputFile := filepath.Join(rootDirectory, "assets.go")
    fileListOutput := filepath.Join(rootDirectory, "filelist.go")

    fmt.Println("Traversing directory:", rootDirectory)

    filePaths, err := traverseDirectory(rootDirectory)
    if err != nil {
        fmt.Println("Error traversing directory:", err)
        return
    }

    if len(filePaths) == 0 {
        fmt.Println("No files found in the directory.")
        return
    }

    // Write the output file for embedded assets
    err = writeGoFile(filePaths, outputFile, rootDirectory)
    if err != nil {
        fmt.Println("Error writing Go file:", err)
    }

    // Write the file list file
    err = writeFileListGo(filePaths, fileListOutput, rootDirectory)
    if err != nil {
        fmt.Println("Error writing file list Go file:", err)
    }
}