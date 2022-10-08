package main

// I am not a huge fan of GOLANG regardless, but I found implementing this was quite haphazard and wordy.
// However, the performance completely blows Python out of the water, so maybe I could benefit from expanding on the process in Python


import (
	"math/rand"
	"fmt"
	"time"
	"bufio"
	"os"
	"strings"
)

var seed = rand.New(rand.NewSource(time.Now().UnixNano()))
var alphabet = "abcdefghijklmnopqrstuvwxyz" // Define chars to use in random string

func random_string() string {
    tmp := make([]byte, 24) // create byte slice
    for i := range tmp { // iterate over range of byte slice (24)
        tmp[i] = alphabet[seed.Intn(len(alphabet)-1)] // add to byte slice in position of i random letter from alphabet
    }
    return string(tmp)
}

func concat_strings() string {
	// Return 2 random strings delimited
	return random_string() + "," + random_string()
}

func generate_data() []string {
	random_string()
	var lines []string

	// Measure time then generate 5m random strings appended to a string slice
	t1 := time.Now()
	fmt.Println("Generating data")
	for len(lines) < 5000000 {
		lines = append(lines, concat_strings())
	}

	t2 := time.Since(t1)
	fmt.Printf("Done, duration: %s\n", t2)

	return lines
}

func write_file(data []string) {
	// Writes a slice of data to test_data-go.txt
	t1 := time.Now()

	fmt.Println("Writing data to test_data-go.txt")
    f, error := os.Create("test_data-go.txt") // Open test_data-go.txt
    if error != nil {
    	fmt.Println(error) // Throw an error if it cant
    }


    writer := bufio.NewWriter(f) // Open file writer
    for _, line := range data { // Iterate over data then parse the string
    	_, _ = writer.WriteString(line + "\n")
    }
    // Close writer stream + file
    writer.Flush()
    f.Close()

	t2 := time.Since(t1)
	fmt.Printf("Done, duration: %s\n", t2)
}


func read_file() {
	// Reads test_data-go.txt and tests how fast it can parse all the lines (5m into 10m)
	t1 := time.Now()

	fmt.Println("Reading and parsing data from test_data-go.txt")
	var lines []string

	f, error := os.Open("test_data-go.txt") // Opens file etc.
	if error != nil {
		fmt.Println(error)
	}

	reader := bufio.NewScanner(f) // Creates new bufio scanner

	reader.Split(bufio.ScanLines) // Pretty sure this splits the file into a slice
	for reader.Scan() { // Loops over file contents
		line := strings.Split(reader.Text(), ",") // Parses string
		lines = append(lines, line[0]) // Adds it to slice
		lines = append(lines, line[1])
	}
	f.Close()

	t2 := time.Since(t1)
	fmt.Printf("Done, duration: %s\n", t2)

}
func main() {
	start := time.Now()

	lines := generate_data()
	write_file(lines)
	read_file()

	end := time.Since(start)
	fmt.Printf("TOTAL DURATION: %s\n", end)
}