// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type (
	// task that adds license header to source
	// files, if they don't already exist
	addLicenseHeaderTask struct {
		license string  // license header string to add
		config  *config // root directory of the project source
	}

	// command line config params
	config struct {
		licenseFile string
		scanDir     string
		verifyOnly  bool
	}
)

// licenseFileName is the name of the license file
const licenseFileName = "./LICENSE"

// unique prefix that identifies a license header
const licenseHeaderPrefix = "// The MIT License"

var (
	// directories to be excluded
	dirBlocklist = []string{".gen/", ".git/", ".vscode/", ".idea/"}
	// default perms for the newly created files
	defaultFilePerms = os.FileMode(0644)
)

// command line utility that adds license header
// to the source files. Usage as follows:
//
//	./cmd/tools/copyright/licensegen.go
func main() {
	var cfg config
	flag.StringVar(&cfg.licenseFile, "licenseFile", licenseFileName, "directory to scan")
	flag.StringVar(&cfg.scanDir, "scanDir", ".", "directory to scan")
	flag.BoolVar(&cfg.verifyOnly, "verifyOnly", false, "don't automatically add headers, just verify all files")
	flag.Parse()

	task := newAddLicenseHeaderTask(&cfg)
	if err := task.run(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func newAddLicenseHeaderTask(cfg *config) *addLicenseHeaderTask {
	return &addLicenseHeaderTask{
		config: cfg,
	}
}

func (task *addLicenseHeaderTask) run() error {
	data, err := os.ReadFile(task.config.licenseFile)
	if err != nil {
		return fmt.Errorf("error reading license file, errr=%v", err.Error())
	}

	task.license, err = commentOutLines(string(data))
	if err != nil {
		return fmt.Errorf("copyright header failed to comment out lines, err=%v", err.Error())
	}

	err = filepath.Walk(task.config.scanDir, task.handleFile)
	if err != nil {
		return fmt.Errorf("copyright header check failed, err=%v", err.Error())
	}
	return nil
}

func (task *addLicenseHeaderTask) handleFile(path string, fileInfo os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		return nil
	}

	if !mustProcessPath(path) {
		return nil
	}

	if !strings.HasSuffix(fileInfo.Name(), ".go") {
		return nil
	}
	if isFileAutogenerated(path) {
		return nil
	}

	// Used as part of the cli to write licence headers on files, does not use user supplied input so marked as nosec
	// #nosec
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	readLineSucc := scanner.Scan()
	if !readLineSucc {
		return fmt.Errorf("fail to read first line of file %v", path)
	}
	firstLine := strings.TrimSpace(scanner.Text())
	if err := scanner.Err(); err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}

	if strings.Contains(firstLine, licenseHeaderPrefix) {
		return nil // file already has the copyright header
	}

	// at this point, src file is missing the header
	if task.config.verifyOnly {
		return fmt.Errorf("%v missing license header", path)
	}

	// Used as part of the cli to write licence headers on files, does not use user supplied input so marked as nosec
	// #nosec
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return os.WriteFile(path, []byte(task.license+string(data)), defaultFilePerms)
}

func isFileAutogenerated(path string) bool {
	// stringer doesn't have an option to include a license header
	return strings.HasSuffix(path, "_string_gen.go")
}

func mustProcessPath(path string) bool {
	for _, d := range dirBlocklist {
		if strings.HasPrefix(path, d) {
			return false
		}
	}
	return true
}

func commentOutLines(str string) (string, error) {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(str))
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			lines = append(lines, "//\n")
		} else {
			lines = append(lines, fmt.Sprintf("// %s\n", line))
		}
	}
	lines = append(lines, "\n")

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return strings.Join(lines, ""), nil
}
