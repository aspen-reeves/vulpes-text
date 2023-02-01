package stuff

import (
	"bytes"
	"errors"
	"os"
)

func OpenFile() ([]byte, error) {
	var err error
	if len(os.Args) < 2 {
		err = errors.New("please provide a file name")
		panic(err)
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		err = errors.New("cant read file")
		panic(err)
		//crash the fucking program if god forbid  the file doesnt exist
	}
	return data, err
}

func SaveFile(lines []string) error {
	data := StrToByte(lines)
	if len(os.Args) < 2 {
		err := errors.New("please provide a file name")
		panic(err)
	}
	err := os.WriteFile(os.Args[1], data, 0644) //write to file
	if err != nil {
		err = errors.New("cant write to file")
		panic(err)
		//crash the fucking program if god forbid  the file doesnt exist
	}

	return nil
}
func ByteToStr(data []byte) []string {
	//convert file to array of strings
	lines := make([]string, 0)
	for _, line := range bytes.Split(data, []byte{'\n'}) { //split data by newline
		lines = append(lines, string(line)) // add line to lines
	}

	return lines

}
func StrToByte(lines []string) []byte {
	//convert array of strings to byte array
	data := make([]byte, 0)           //create byte array
	for i := 0; i < len(lines); i++ { //loop through lines
		data = append(data, []byte(lines[i])...) //add line to data
		data = append(data, '\n')                //add newline to data
	}
	return data //return data

}
