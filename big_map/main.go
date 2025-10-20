package main

func main() {
	m := make(map[int]int)
	for i := range 50_000_000 {
		m[i] = i
	}
}
