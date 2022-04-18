package tags

import "fmt"

type Person struct {
	Name  string `serialize: "name"`
	City  string `serialize: "city"`
	State string
	Misc  string `serialize: "-"`
	Year  int    `serialize: "year"`
}

// EmptyStruct함수는 태그를 갖는 빈(Empty) 구조체의 직렬화 및 역질력화 방법을 보여준다.
func EmptyStruct() error {
	p := Person{}

	res, err := SerializeStructStrings(&p)
	if err != nil {
		return err
	}

	fmt.Printf("Empty struct: %#v\n", p)
	fmt.Println("Serialize Results:", res)

	newP := Person{}
	if err := DeSerializeStructStrings(res, &newP); err != nil {
		return err
	}
	fmt.Println("Deserialize Results:", res)
	return nil
}

// FullStruct함수는 태그가 있는 데이터가 모두 채워진 구조체의 직렬화 및 역직렬화 방법을 보여준다.
func FullStruct() error {
	p := Person{
		Name:  "Aaron",
		City:  "Seattle",
		State: "WA",
		Misc:  "some fact",
		Year:  2017,
	}
	res, err := SerializeStructStrings(&p)
	if err != nil {
		return err
	}
	fmt.Printf("Full struct: %#v\n", p)
	fmt.Println("Serialize Results:", res)

	newP := Person{}
	if err := DeSerializeStructStrings(res, &newP); err != nil {
		return err
	}
	fmt.Printf("Deserialize results: %#v\n", newP)
	return nil
}
