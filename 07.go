package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type commonError struct {
	errorCode int
	errorMsg  string
}

func (ce *commonError) Error() string {
	return ce.errorMsg
}

type MyError struct {
	err error
	msg string
}

func (me MyError) Error() string {
	return me.err.Error() + me.msg
}

func ReadFile(filename string) ([]byte, error) {
	open, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer open.Close()
	// todo
	return []byte{}, nil
}

func connectMySQL(ip, username, password string) {
	if ip == "" {
		panic("ip不可为空")
	}
	// ...
}

func main() {
	atoi, err := strconv.Atoi("a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(atoi)
	}
	//err = errors.New("具体错误信息")

	c := &commonError{
		errorCode: 1,
		errorMsg:  "错误信息",
	}
	println(c.Error())

	err = error(c)
	if c2, ok := err.(*commonError); ok {
		fmt.Printf("错误代码：%d,错误信息为：%s\n", c2.errorCode, c2.errorMsg)
	} else {
		fmt.Println("其他类型错误")
	}

	myError := MyError{err, "数据问题"}
	fmt.Println(myError)

	err = errors.New("原始错误")
	//err = fmt.Errorf("包装了一层错误：%w", err)
	//err = fmt.Errorf("包装了一层错误：%w", myError)
	fmt.Println(err)

	fmt.Println(errors.Unwrap(err))

	var cm *commonError
	if errors.As(c, &cm) {
		fmt.Printf("错误代码为：%d，错误信息为：%s\n", cm.errorCode, cm.errorMsg)
	} else {
		fmt.Println("其他错误")
	}

	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	connectMySQL("", "", "")
}
