package df

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/clh021/text-parser/lib"
)

type commandOutputdrowIdx int

const (
	command_output_row_idx_file_system commandOutputdrowIdx = iota
	command_output_row_idx_size
	command_output_row_idx_used
	command_output_row_idx_available
	command_output_row_idx_used_percent
	command_output_row_idx_mounted_on
)

// headers is the headers in 'df -h' output.
var headers = []string{
	"Filesystem",
	"Size",
	"Used",
	"Avail",
	"Use%",
	// Mounted on
	"Mounted",
	"on",
}

type FileSystemItem struct {
	Filesystem string
	Size       string
	Used       string
	Avail      string
	UsePercent string
	MountedOn  string
}

type FileSystemList struct{}

func GetFileSystem() []FileSystemItem {
	list := NewFileSystemList()
	return list.Run()
}

func NewFileSystemList() *FileSystemList {
	return &FileSystemList{}
}

func (l FileSystemList) Run() []FileSystemItem {
	env := make([]string, 0)
	env = append(env, "LANGUAGE=en_US")
	output, err := lib.ExecGetCmdStdoutWithEnv(env, "df", "-lh")
	if err != nil {
		log.Println(err)
	}
	result, err := l.parse(string(output))
	if err != nil {
		log.Println(err)
	}
	return result
}

// Parse parses 'df -h' command output and returns the rows.
func (l FileSystemList) parse(output string) ([]FileSystemItem, error) {
	lines := strings.Split(output, "\n")
	rows := make([]FileSystemItem, 0, len(lines))
	headerFound := false
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		ds := strings.Fields(strings.TrimSpace(line))
		if ds[0] == headers[0] {
			if !reflect.DeepEqual(ds, headers) {
				return nil, fmt.Errorf(`unexpected 'df -h' command header order (%v, expected %v, output: %q)`, ds, headers, output)
			}
			headerFound = true
			continue
		}

		if !headerFound {
			continue
		}

		row, err := l.parseFileSystemItem(ds)
		if err != nil {
			return nil, err
		}

		rows = append(rows, row)
	}

	return rows, nil
}

func (l FileSystemList) parseFileSystemItem(row []string) (FileSystemItem, error) {
	if len(row) != len(headers)-1 {
		return FileSystemItem{}, fmt.Errorf(`unexpected row column number %v (expected %v)`, row, headers)
	}

	return FileSystemItem{
		Filesystem: strings.TrimSpace(row[command_output_row_idx_file_system]),
		Size:       strings.TrimSpace(row[command_output_row_idx_size]),
		Used:       strings.TrimSpace(row[command_output_row_idx_used]),
		Avail:      strings.TrimSpace(row[command_output_row_idx_available]),
		UsePercent: strings.TrimSpace(strings.Replace(row[command_output_row_idx_used_percent], "%", " %", -1)),
		MountedOn:  strings.TrimSpace(row[command_output_row_idx_mounted_on]),
	}, nil
}
