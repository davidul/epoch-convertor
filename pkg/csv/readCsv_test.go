package csv

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseLine(t *testing.T) {
	parser := NewParser(',', '#', nil)
	line2 := parser.ParseLine([]byte("a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z"))
	assert.NotNil(t, line2)
	assert.Equal(t, 26, len(line2.Value))
}

var testCsv = "col_1, col_2, col_3\n" +
	"1, val_2, 3.8"

func TestReadFileAsByteBuffer(t *testing.T) {
	ReadFile = func(name string) ([]byte, error) {
		return []byte(testCsv), nil
	}
	bytes := ReadFileAsByteBuffer("test.csv")
	assert.NotNil(t, bytes)
	assert.Equal(t, len(testCsv), bytes.Len())
}

func TestParseFile(t *testing.T) {
	ReadFile = func(name string) ([]byte, error) {
		return []byte(testCsv), nil
	}
	parser := NewParser(',', 0, nil)
	file := parser.ParseFile(ReadFileAsByteBuffer("test.csv"))
	assert.NotNil(t, file)
	assert.Equal(t, 2, len(file.Records))
}

func TestFirstLineHeader(t *testing.T) {
	ReadFile = func(name string) ([]byte, error) {
		return []byte(testCsv), nil
	}
	parser := NewParser(',', 0, nil)
	headerParser := NewFirstLineHeaderParser(',', parser)
	records := headerParser.ParseFile(ReadFileAsByteBuffer("test.csv"))
	fmt.Println(records)
	record := records.Records[0]
	fmt.Println(record)
}
