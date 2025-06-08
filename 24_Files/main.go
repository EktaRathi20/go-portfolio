package main

import (
	"bufio"
	"fmt"
	"os"
)

func getFileInfo(filename string) {
	f, error := os.Open(filename)

	if error != nil {
		panic(error)
	}

	fileInfo, error := f.Stat()
	if error != nil {
		panic(error)
	}

	fileInformation := map[string]interface{}{
		"Name":       fileInfo.Name(),
		"Size":       fileInfo.Size(),
		"Mode":       fileInfo.Mode(),
		"ModTime":    fileInfo.ModTime(),
		"IsDir":      fileInfo.IsDir(),
		"modifyTime": fileInfo.ModTime(),
	}

	fmt.Println("File Information:", fileInformation)
}

// readFile1 reads a specified number of bytes from the file and prints it
// This is useful for larger files where you want to read only a portion of the file.
// It uses a buffer to read a fixed number of bytes.
// It is more efficient for large files as it does not load the entire file into memory.
func readFile1(filename string) {
	f, error := os.Open(filename)
	if error != nil {
		panic(error)
	}

	defer f.Close()

	buffer := make([]byte, 100) // Read up to 100 bytes

	_, error = f.Read(buffer)
	if error != nil {
		panic(error)
	}

	fmt.Println("File Content:", string(buffer))
}

// readFile2 reads the entire file content into memory and prints it
// This is useful for smaller files where you want to read everything at once.
func readFile2(filename string) {
	f, error := os.ReadFile(filename)
	if error != nil {
		panic(error)
	}
	fmt.Println("File Content:", string(f))

}

func getFolderDirectory() {
	dir, error := os.Open("./")

	if error != nil {
		panic(error)
	}

	defer dir.Close()

	files, error := dir.ReadDir(-1) // -1 means read all files

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func createFile(filename string) {
	f, error := os.Create(filename)
	if error != nil {
		panic(error)
	}

	defer f.Close()

	// Write some content to the file
	f.WriteString("Hello, World! \n")
	f.WriteString("This is a new file created using os.Create.")

	fmt.Println("File created successfully:", filename)
}

func transferFileData(sourceFile, destinationFile string) {
	sourceF, error := os.Open(sourceFile)
	if error != nil {
		panic(error)
	}
	defer sourceF.Close()

	destinationF, error := os.Create(destinationFile)
	if error != nil {
		panic(error)
	}
	defer destinationF.Close()

	reader := bufio.NewReader(sourceF)
	writer := bufio.NewWriter(destinationF)

	for {
		b, error := reader.ReadByte()

		if error != nil {
			if error.Error() != "EOF" {
				panic(error)
			}

			break
		}

		err := writer.WriteByte(b)

		if err != nil {
			panic(err)
		}
	}

	writer.Flush()
	fmt.Println("File data transferred from", sourceFile, "to", destinationFile)
}

func deleteFile(filename string) {
	error := os.Remove(filename)
	if error != nil {
		panic(error)
	}
	fmt.Println("File deleted successfully:", filename)
}

func main() {
	getFileInfo("24_Files/example.txt")
	readFile1("24_Files/example.txt")
	readFile2("24_Files/example.txt")
	getFolderDirectory()
	createFile("24_Files/newfile.txt")
	transferFileData("24_Files/example.txt", "24_Files/copiedfile.txt")
	deleteFile("24_Files/newfile.txt")
}
