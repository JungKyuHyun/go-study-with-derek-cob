package main

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
