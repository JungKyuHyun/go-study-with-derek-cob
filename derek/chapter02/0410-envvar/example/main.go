package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter2/envvar"
)

// json 파일과 환경변수로부터 읽어온 값을 저장.
type Config struct {
	Version string `json:"version" required:"true"`
	IsSafe  bool   `json:"is_safe" default:"true"`
	Secret  string `json:"secret"`
}

func main() {
	var err error

	// temp 파일 생성.
	tf, err := ioutil.TempFile("", "tmp")
	if err != nil {
		panic(err)
	}
	defer tf.Close()
	defer os.Remove(tf.Name())

	secrets := `{
        "secret": "so so secret"
    }`
	// tf 에 secrets json을 string으로 작성.
	if _, err = tf.Write(bytes.NewBufferString(secrets).Bytes()); err != nil {
		panic(err)
	}

	// os.Setenv 로 직접적으로 환경변수를 조작할 수 있음.
	if err = os.Setenv("EXAMPLE_VERSION", "1.0.0"); err != nil {
		panic(err)
	}
	if err = os.Setenv("EXAMPLE_ISSAFE", "false"); err != nil {
		panic(err)
	}

	c := Config{}
	if err = envvar.LoadConfig(tf.Name(), "EXAMPLE", &c); err != nil {
		panic(err)
	}

	fmt.Println("secrets file contains =", secrets)

	// os.Getenv로 쉽게 환경변수 값을 가져올 수 있음.
	fmt.Println("EXAMPLE_VERSION =", os.Getenv("EXAMPLE_VERSION"))
	fmt.Println("EXAMPLE_ISSAFE =", os.Getenv("EXAMPLE_ISSAFE"))

	fmt.Printf("Final Config: %#v\n", c)

}
