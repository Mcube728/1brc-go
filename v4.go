package main
import (
	"fmt"
	"os"
	"io"
	"sort"
	"bytes"
)


func Parse_temp_optimised(tempBytes []byte) (float64, int){
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
	return temp, index
}


func v4(inputPath string, output io.Writer) error{	
	// Read File
	file, err := os.Open(inputPath)
	if err != nil{
		fmt.Println("File reading error: ", err)
		return err
	}
	defer file.Close()

	stations := make(map[string]*Station) 	// create an empty map for stations

	// Read Lines
	buffer := make([]byte, 1024*1024)
	readStart := 0
	for {
		n, err := file.Read(buffer[readStart:])
		if err != nil && err != io.EOF { 
			return err
		}
		if readStart + n == 0 { 	// reached end of file
			break
		}
		chunk := buffer[:readStart+n]
		newline := bytes.LastIndexByte(chunk, '\n')
		if newline < 0 {
			break
		}
		remaining := chunk[newline+1:]
		chunk = chunk[:newline+1]

		for{
			station, after, hasSemi := bytes.Cut(chunk, []byte(";"))
			if !hasSemi {
				break
			}
			temp, index := Parse_temp_optimised(after)
			chunk = after[index:]
			s := stations[string(station)]
			if s == nil {
				stations[string(station)] = &Station{
					Name: string(station),
					temp_min: temp,
					temp_max: temp,
					temp_sum: temp,
					count: 1,
				}
			} else {
				s.Update(temp)
			}
		}
		readStart = copy(buffer, remaining)
	}
	
	stations_list := make([]string, 0, len(stations))
	for station := range stations{
		stations_list = append(stations_list, station)
	}
	sort.Strings(stations_list)

	fmt.Fprint(output, "{")
	for _, station := range stations_list {
		s := stations[station]
		fmt.Fprintf(output, "%s: min=%.1f, max=%.1f, mean=%.1f,\n", station, s.temp_min, s.temp_max, s.Mean())
	}
	fmt.Fprint(output, "}\n")
	return nil
}