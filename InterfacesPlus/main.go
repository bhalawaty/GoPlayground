package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// Open the file
	body, err := os.Open("bilalTestFile.text")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(body *os.File) {
		err := body.Close()
		if err != nil {
			fmt.Println("Error Closing file:", err)

		}
	}(body) // Ensure the file is closed when done

	// Copy the content to stdout
	//_, err = io.Copy(os.Stdout, body)
	//if err != nil {
	//	fmt.Println("Error copying file content:", err)
	//	return
	//}
	// Create a buffer to store data read from the file
	sliceBody := make([]byte, 32*1024)

	// Read the file in chunks
	for {
		// Read data into sliceBody
		n, err := body.Read(sliceBody)
		if err == io.EOF {
			// Reached end of file, stop reading
			fmt.Println("End of file reached.")
			break
		}
		if err != nil {
			// Handle any other errors (not EOF)
			fmt.Println("Error reading file:", err)
			return
		}

		// Process the chunk of data (n bytes)
		fmt.Printf("Read %d bytes: %s\n", n, string(sliceBody[:n]))
	}

}
