package helper

import "fmt"

func ReturnError(err error) {
	if err != nil {
		fmt.Println(err)
	}

}
