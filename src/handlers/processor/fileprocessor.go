package processor

import (
	"bufio"
	"fmt"
	"handlers/commandmanager"
	"os"
	"strings"
	"utils/constant"
)

// FileProcessor type
type FileProcessor struct {
	Filename string
	Resp     string
}

// NewFileProcessor - Create a new File Processor
func NewFileProcessor(filename string) *FileProcessor {
	fp := new(FileProcessor)
	fp.Filename = filename

	return fp
}

// Run - Run the file process
func (fp *FileProcessor) Run() {
	f, err := os.Open(fp.Filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	cm := commandmanager.NewCommandManager()

	var input string
	var tempResp []string
	for s.Scan() {
		input = s.Text()
		result, err := cm.Execute(input)
		if err != nil {
			fmt.Println("Error:", err.Error())
		} else {
			fmt.Println(result)
			tempResp = append(tempResp, result)
		}
	}

	if err := s.Err(); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Print(constant.NewLine)
	fp.Resp = strings.Join(tempResp, constant.NewLine)
}
