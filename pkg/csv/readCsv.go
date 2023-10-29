package csv

import (
	"fmt"
	"github.com/davidul/buffers/davidul/seekbuffer"
	"os"
)

type CsvReader struct {
	Parser  Parser
	Records Records
}

type Records struct {
	Header  *Header
	Records []Record
	Index   map[string][]Record
}

type Header struct {
	Cols []string
}

type Record struct {
	Value []string
}

func NewCsvReader() *CsvReader {
	return &CsvReader{
		Parser: &BaseParser{},
	}
}

var ReadFile = func(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func ReadFileAsByteBuffer(filePath string) *seekbuffer.SeekBuffer {
	file, err := ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	b := new(seekbuffer.SeekBuffer)
	_, err = b.Write(file)

	if err != nil {
		fmt.Println(err)
	}

	return b
}

func CreateIndex() {
	m2 := make(map[string][]Record)
	for i := range table.Records {
		record, ok := m2[table.Records[i].Value[0]]
		if ok {
			m2[table.Records[i].Value[0]] = append(record, table.Records[i])
		} else {
			m2[table.Records[i].Value[0]] = []Record{table.Records[i]}
		}
	}
	table.Index = m2
}

func GetRecord(n int) Record {
	return table.Records[n]
}

func Count() int {
	return len(table.Records)
}
