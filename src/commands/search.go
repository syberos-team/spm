package commands

import (
	"core"
	"core/log"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"strings"
	"syscall"
)

//SearchCommand 模糊搜索包
type SearchCommand struct {
	Command
	packageName string
}

func (s *SearchCommand) RegisterArgs(args ...string) {
	if args==nil || len(args)==0 {
		return
	}
	s.packageName = args[0]
}

func (s *SearchCommand) Description() string {
	return "Searches for packages containing the given string"
}

func (s *SearchCommand) Run() error {
	if s.packageName=="" {
		return errors.New("you must enter the package name keyword to be searched")
	}

	client := core.NewSpmClient()
	req := &core.SearchRequest{
		PackageName: s.packageName,
	}
	resp, err := client.Search(req)
	if err!=nil {
		return err
	}
	if core.CODE_ERROR==resp.Code {
		return errors.New(resp.Msg)
	}
	s.printData(resp.Data)
	return nil
}

const (
	columnName        = "Name"
	columnDescription = "Description"
)

func (s *SearchCommand) printData(results []*core.SearchResponseData) {
	if results==nil || len(results)==0 {
		log.Warning("No packages found")
		return
	}
	columnWidths := map[string]int{
		columnName: len(columnName),
		columnDescription: len(columnDescription),
	}
	if len(results) < 1000 {
		// pre-process the list to get column widths
		for _, r := range results {
			columnWidths[columnName] = s.intMin(s.intMax(columnWidths[columnName], len(r.Name)), 40)
			columnWidths[columnDescription] = s.intMin(s.intMax(columnWidths[columnDescription], len(r.Description)), 80)
		}
	} else {
		// Too many results to pre-process so use sensible defaults
		columnWidths[columnName] = 40
		columnWidths[columnDescription] = 60
	}
	width, _, err := terminal.GetSize(int(syscall.Stdout))
	if err != nil {
		fmt.Printf("Couldn't get terminal width: %s\n", err.Error())
		// gracefully fallback to something sensible
		width = 110
	}
	const columnSpacing = 3

	fmt.Println("")
	columns := []string{columnName, columnDescription}
	widths := make([]int, len(columns))
	for i, col := range columns {
		widths[i] = columnWidths[col]
	}

	// Print the headers
	s.printRow(width, columnSpacing, widths, columns)

	// Print a horizontal line
	fmt.Printf("%s\n", strings.Repeat("-", width))

	// Print the search results
	for _, r := range results {
		columns := []string{
			r.Name,
			r.Description,
		}
		s.printRow(width, columnSpacing, widths, columns)
	}
}

func (s *SearchCommand) printRow(screenWidth int, columnSpacing int, columnWidths []int, columns []string) {
	remaining := screenWidth
	for i, col := range columns {
		if remaining <= 0 {
			break
		}
		// convert to []rune since we want the char count not bytes
		runes := []rune(col)
		// truncate the string if we are out of space
		maxLength := s.intMin(remaining, columnWidths[i])
		if len(runes) > maxLength {
			runes = runes[:maxLength]
		}

		fmt.Printf("%s", string(runes))
		w := columnWidths[i]
		remaining -= len(runes)
		toNextCol := s.intMax(s.intMin(w-len(runes)+columnSpacing, remaining), 0)
		fmt.Printf("%s", strings.Repeat(" ", toNextCol))
		remaining -= toNextCol
	}
	fmt.Printf("\n")
}

func (s *SearchCommand) intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (s *SearchCommand) intMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func NewSearchCommand() *SearchCommand{
	return &SearchCommand{}
}



