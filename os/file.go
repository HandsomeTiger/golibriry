package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// 创建一个文件
func insertLog(str string) error {
	fn, err := fileName()
	if err != nil {
		return err
	}
	fx, err := fileExists(fn)
	if err != nil {
		return err
	}
	var f *os.File
	if fx {
		f, err = os.OpenFile(fn, os.O_RDWR, 0666)
		if err != nil {
			return err
		}
	} else {
		f, err = os.Create(fn)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	data := make([]byte, 10, 100)
	read, err := f.Read(data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(read)
		fmt.Println(data)
	}
	stat, err := f.Stat()
	fmt.Printf("%+v", stat)
	f.Seek(stat.Size(), 0)
	f.WriteString("nrmrs")
	return nil
}

func fileName() (string, error) {
	cdir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("getwd faild: %v", err.Error())
		return "", err
	}
	fileName := cdir + string(filepath.Separator) + "log_" + time.Now().Format("2006-01-02") + ".txt"
	return fileName, nil
}

func fileExists(fileName string) (bool, error) {
	_, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	} else {
		return true, nil
	}
}
