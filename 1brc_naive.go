package main
import (
	"bufio"
	"fmt"
	"os"
	"time"
	"strconv"
	"strings"
)


type Station struct {
	Name 		string
	temp_min 	float64
	temp_max 	float64
	temp_sum 	float64
	count 		int
}


func (s *Station) Mean() float64{
	// a simple function to calculate a station's mean. 
	return s.temp_sum / float64(s.count)
}


func (s *Station) Update(temp float64){
	if s.count == 0 { // if station is encountered for first time
		s.temp_min = temp
		s.temp_max = temp
	} else {
		s.temp_min = min(s.temp_min, temp)
		s.temp_max = max(s.temp_max, temp)
	}
	s.temp_sum += temp
	s.count += 1
}


func main(){
	t0 := time.Now()
	file_path := "measurements.txt"
	
	// Read File
	file, err := os.Open(file_path)
	if err != nil{
		fmt.Println("File reading error: ", err)
		return
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
		return
	}
	c := 0
	for name, station := range stations {
		if c >= 15 { break }
		fmt.Printf("%s: min=%.1f, max=%.1f, mean=%.1f\n", name, station.temp_min, station.temp_max, station.Mean())
		c++
	}
	elapsed := time.Since(t0)
	fmt.Printf("\nProcessed %d stations in %v\n", len(stations), elapsed)
}