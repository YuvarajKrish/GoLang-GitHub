package main

import (
	"os"
	"log"
	"io"
	"os/exec"
	"bufio"
	"strings"
	"fmt"
	"time"
)

func main(){
	files := getChangedFiles();
	copyFiles(files)
}

func getChangedFiles() string{
	log.Println("Getting changed files..")
	cmd := exec.Command("git", "ls-files","-m")
	//run the command
	op, err := cmd.Output()
	var files string
	if err != nil {
		log.Println("Please check the directory!")
		log.Println(err)
	}else{
		files = string(op)
		log.Print(files)
	}
	return files
}

func copyFiles(files string){
	scanner := bufio.NewScanner(strings.NewReader(files))
	currentTime := time.Now()
	currentTimeStr := currentTime.Format("01-02-2006")
	destFolder := "C:\\Vista_Backup\\"+currentTimeStr
	if _, err := os.Stat(destFolder); os.IsNotExist(err) {
		err = os.MkdirAll(destFolder, 0755)
              if err != nil {
                      panic(err)
              }
	}
	for scanner.Scan() {
		file := scanner.Text()
		file = strings.Replace(file, "/", "\\",-1)
		log.Println("Copying file ",file)
		pwd, _ := os.Getwd()
		srcPath := pwd+"\\"+file
		fileName := srcPath[strings.LastIndex(srcPath, "\\"):len(srcPath)]
		destPath := destFolder+"\\"+fileName
		copy(srcPath,destPath)
	}
	log.Println("Files copied!")
}

func copy(src, dst string) (int64, error) {
	log.Println(src,dst)
	sourceFileStat, err := os.Stat(src)
	if err != nil {
			return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
			return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
			return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
			return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	if err != nil {
		log.Println("Please check the directory!")
		log.Println(err)
	}
	return nBytes, err
}