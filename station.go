package main 

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