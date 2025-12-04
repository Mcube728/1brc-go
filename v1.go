package main
import (
	"bufio"
	"fmt"
	"os"
	"io"
	"strconv"
	"strings"
)


func v1(inputPath string, output io.Writer) error {
	// Read File
	file, err := os.Open(inputPath)
	if err != nil{
		fmt.Println("File reading error: ", err)
		return err
	}
	defer file.Close()

	stations := make(map[string]*Station) 	// create an empty map for stations

	// Read Lines
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		parts := strings.Split(line, ";") 	// split at semicolon
		name := parts[0]
		temp, err := strconv.ParseFloat(parts[1], 64)
		if err != nil{
			continue
		}
		if station, exists := stations[name]; exists {
			station.Update(temp)
		} else {
			station := &Station{Name: name}
			station.Update(temp)
			stations[name] = station
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil{
		fmt.Println("Error reading file: ", err)
		return err
	}
	
	fmt.Fprint(output, "{")
	for name, station := range stations {

		fmt.Fprintf(output, "%s: min=%.1f, max=%.1f, mean=%.1f,\n", name, station.temp_min, station.temp_max, station.Mean())
	}
	fmt.Fprint(output, "}\n")
	return nil
}