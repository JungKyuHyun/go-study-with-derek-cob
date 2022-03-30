package main

import filedirs "github.com/jungkyuhyun/go-study-with-derek-cob/cob/0330-filedirs"

func main() {
	if err := filedirs.Operate(); err != nil {
		panic(err)
	}
	if err := filedirs.CapitalizerExample(); err != nil {
		panic(err)
	}
}
