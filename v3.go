package main
import (
	"bufio"
	"fmt"
	"os"
	"io"
	"bytes"
)


func Parse_temp( tempBytes []byte) float64{
	negative := false
	index := 0
	if tempBytes[index] == '-' {
		index ++
		negative = true
	}
	temp := float64(tempBytes[index] - '0')
	index++
	if tempBytes[index] != '.' {
		temp = temp*10 + float64(tempBytes[index]-'0')
		index++
	}
	index++
	temp += float64(tempBytes[index]-'0') / 10
	if negative{
		temp = -temp
	}
	return temp
}


func v3(inputPath string, output io.Writer) error{	
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
		line := scanner.Bytes()
		sep_idx := bytes.IndexByte(line, ';')	// split at semicolon, then slice string to get station name and temperature, should be less computationally expensive opposed to strings.split?
		
		name := string(line[:sep_idx])
		tempBytes := line[sep_idx+1:]
		temp := Parse_temp(tempBytes)
		
		s := stations[name]
		if s == nil {
			stations[name] = &Station{
				Name: name,
				temp_min: temp,
				temp_max: temp,
				temp_sum: temp,
				count: 1,
			}
		} else {
			s.Update(temp)
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