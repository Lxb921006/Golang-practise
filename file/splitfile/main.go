package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	var rb = make([]byte, 536870912)
	file := "D:\\工作工具\\SQLServer2019-x64-CHS.iso"
	path := "C:\\Users\\Administrator\\Desktop\\update"

	f, err := os.Open(file)
	if err != nil {
		return
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			return
		}
	}(f)

	num := 0

	for {
		n, err := f.Read(rb)
		if err == io.EOF {
			break
		}

		if err != nil {
			return
		}

		output := filepath.Join(path, fmt.Sprintf("split_%d", num))

		fn, err := os.Create(output)
		if err != nil {
			return
		}

		wn, err := fn.Write(rb[:n])
		if err != nil {
			return
		}

		num++
		fmt.Println("already write byte >>>", wn)
	}

	fmt.Println("done")
}
