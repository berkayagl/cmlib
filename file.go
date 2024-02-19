package cmlib

import (
	"bufio"
	"os"
)

// is a function used to read a certain number of lines from a given file.
func ReadLines(s string, i int) ([]string, error) {

	// opens the specified file.
	file, err := os.Open(s)

	// errors control
	if err != nil {
		return nil, err
	}

	// automatically switches off when the function ends.
	defer file.Close()

	// A slice with capacity `i` and no initial(0) elements is created.
	lines := make([]string, 0, i)

	// is used to scan the contents of the specified file.
	nc := bufio.NewScanner(file)

	// It works as long as `x` increases from 0 to `i` and `sc.Scan()` reads the next line.
	for x := 0; x < i && nc.Scan(); x++ {
		lines = append(lines, nc.Text())
	}

	// errors control
	if err := nc.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// is used to read the first line of the specified file.
func ReadFirstLine(s string) (string, error) {

	// opens the specified file.
	file, err := os.Open(s)

	// errors control
	if err != nil {
		return "", err
	}

	// automatically switches off when the function ends.
	defer file.Close()

	// is used to scan the contents of the specified file.
	nc := bufio.NewScanner(file)

	// try to read the next line.
	if nc.Scan() {
		return nc.Text(), nil
	}

	// errors control
	if err := nc.Err(); err != nil {
		return "", err
	}

	return "", nil
}

// read a specific line in the specified file
func ReadLine(s string, i int) (string, error) {

	// opens the specified file.
	file, err := os.Open(s)

	// errors control
	if err != nil {
		return "", err
	}

	// automatically switches off when the function ends.
	defer file.Close()

	// is used to scan the contents of the specified file.
	nc := bufio.NewScanner(file)

	for x := 0; x < i-1 && nc.Scan(); x++ {
		// To skip uninterested lines, `sc.Scan()` is used in the loop to skip `i-1` number of lines.
	}

	// try to read the next line.
	if !nc.Scan() {
		if err := nc.Err(); err != nil {
			return "", err
		}
		return "", nil
	}

	line := nc.Text()
	return line, nil
}

func FileAllLine(path string) ([]string, error) {

	// opens the specified file.
	file, err := os.Open(path)

	// errors control
	if err != nil {
		return nil, err
	}

	// automatically switches off when the function ends.
	defer file.Close()

	// is the slice that will be used to contain the lines of the file.
	var lines []string

	// is used to scan the contents of the specified file.
	nc := bufio.NewScanner(file)

	// reads the next line.
	for nc.Scan() {
		lines = append(lines, nc.Text())
	}

	return lines, nc.Err()
}

// check if a file exists in the given file path
func FileExists(i string) bool {

	// takes the state of the given file path.
	_, err := os.Stat(i)

	// check if the file does not exist, returns `false` if there is no file, otherwise `true`.
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
