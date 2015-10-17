package letter

func ConcurrentFrequency(texts []string) FreqMap {
	ch := make(chan FreqMap)
	for _, t := range texts {
		go func(x string) {
			m := Frequency(x)
			ch <- m
		}(t)
	}

	m := <-ch
	for i := 0; i < len(texts)-1; i++ {
		mx := <-ch
		for k, v := range mx {
			m[k] += v
		}
	}
	return m
}
