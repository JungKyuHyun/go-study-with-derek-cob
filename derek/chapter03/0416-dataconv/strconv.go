package dataconv

import (
	"fmt"
	"strconv"
)

func Strconv() error  {
	s:= "1234"
	// s를 10진수로 64는 얼마나 많은 비트의 정밀도를 파싱할지를 알려줍니다.
	res,err := strconv.ParseInt(s,10,64)
	
	if err != nil{
		return err
	}

	fmt.Println(res)

	res , err = strconv.ParseInt("FF",16,64)
	if err != nil {
		return err
	}

	fmt.Println(res)
	val, err := strconv.ParseBool("true")
	if err != nil {
		return err
	}

	fmt.Println(val)
	return nil
	
}