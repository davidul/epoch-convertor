package csv

import (
	"fmt"
	"github.com/davidul/buffers/davidul/seekbuffer"
	"strings"
)

type Parser interface {
	ParseFile(buffer *seekbuffer.SeekBuffer) Records
	ParseLine(line []byte) Record
}

type FirstLineHeaderParser struct {
	sep    byte
	Parser Parser
}

type BaseParser struct {
	sep     byte
	comment byte
	header  *Header
}

var table Records

// if header nil, guess it
func NewParser(separator, comment byte, header *Header) *BaseParser {
	table = Records{
		Header:  nil,
		Records: make([]Record, 0),
		Index:   nil,
	}
	return &BaseParser{
		sep:     separator,
		comment: comment,
		header:  header,
	}
}

func NewFirstLineHeaderParser(separator byte, parse Parser) *FirstLineHeaderParser {
	return &FirstLineHeaderParser{
		sep:    separator,
		Parser: parse,
	}
}

func (f *FirstLineHeaderParser) ParseFile(buffer *seekbuffer.SeekBuffer) Records {
	line, err := buffer.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	}

	parseLine := f.Parser.ParseLine(line)
	table.Header = &Header{
		Cols: parseLine.Value,
	}

	return f.Parser.ParseFile(buffer)

}

func (p *BaseParser) ParseFile(buffer *seekbuffer.SeekBuffer) Records {

	if p.header != nil {
		table.Header = p.header
	}

	for buffer.Len() > 0 {
		line, err := buffer.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
		}
		//ignore lines that start with #
		if p.comment != 0 {
			if line[0] != p.comment {
				record := p.ParseLine(line)
				table.Records = append(table.Records, record)
			}
		} else {
			record := p.ParseLine(line)
			table.Records = append(table.Records, record)
		}
	}

	return table
}

func (p *BaseParser) ParseLine(line []byte) Record {
	x := len(line)
	tmp := 0
	var record = Record{
		Value: make([]string, 0),
	}
	for i := 0; i < x; i++ {
		if line[i] == p.sep {
			record.Value = append(record.Value, string(line[tmp:i]))
			tmp = i + 1
		}
	}
	record.Value = append(record.Value, strings.TrimRight(string(line[tmp:x]), "\n"))
	return record
}
