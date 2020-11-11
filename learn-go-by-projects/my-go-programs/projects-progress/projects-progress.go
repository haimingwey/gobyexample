package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const projectRootPath = "learn-go-by-projects"

type project struct {
	FullPath          string
	Total             float64
	Name              string
	ProReadMeFileName string
	Progress          string
	ModuleCnt         float64
}

func main() {
	var proj project = project{
		Name:              "learn-go-with-tests",
		ProReadMeFileName: "readme.md",
		Total:             17,
	}
	setFullPath(&proj)
	setModuleCnt(&proj)
	setProgress(&proj)
	updateProgressFile(proj, getLines(proj))
	fmt.Println("done")

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func setFullPath(proj *project) {
	currentPath, err := os.Getwd()
	check(err)
	proj.FullPath = filepath.Join(currentPath, projectRootPath, proj.Name)
	fmt.Println("FullPath:", proj.FullPath)
}

func setModuleCnt(proj *project) {
	paths, err := ioutil.ReadDir(proj.FullPath)
	check(err)
	var cnt float64
	for _, p := range paths {
		if p.IsDir() {
			cnt++
			fmt.Println("find path:", p.Name())
		}
	}
	proj.ModuleCnt = cnt
	fmt.Println("ModuleCnt:", proj.ModuleCnt)
}

func setProgress(proj *project) {
	tmp := proj.ModuleCnt / proj.Total
	progress, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", tmp), 64)
	proj.Progress = fmt.Sprintf("%.2f", progress*100)
	fmt.Println("progress:", proj.Progress)
}

func updateProgressFile(proj project, lines []string) {
	fileName := filepath.Join(proj.FullPath, proj.ProReadMeFileName)
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, os.ModePerm)
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, line := range lines {
		w.WriteString(line)
	}
	w.Flush()
}

func getLines(proj project) []string {
	line1 := "|日期|进度|百分比|\n"
	line2 := "|--|--|--|\n"
	date := time.Now().Format("2006-01-02")
	line3 := fmt.Sprintf("| %s | %d/%d | %s%s |", date, int(proj.ModuleCnt), int(proj.Total), proj.Progress, "%")
	lines := []string{line1, line2, line3}
	return lines
}
