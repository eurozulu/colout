package colout

import (
	"bytes"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io"
	"strings"
)

// ColumnWriter writes comma-delimited strings into fixed sized columns
type ColumnWriter struct {
	Columns      []Column
	ColumnSpacer string
	Out          io.Writer
}

func (cw ColumnWriter) Write(p []byte) (n int, err error) {
	return cw.Out.Write(p)
}

func (cw ColumnWriter) WriteString(s string) (n int, err error) {
	return cw.Out.Write([]byte(s))
}

func (cw ColumnWriter) WriteStrings(ss []string) (n int, err error) {
	buf := bytes.NewBuffer(nil)
	for i, sz := range ss {
		if i > 0 && cw.ColumnSpacer != "" {
			buf.WriteString(cw.ColumnSpacer)
		}
		col := cw.ColumnAtIndex(i)
		buf.WriteString(col.FormatString(sz))
	}
	buf.WriteRune('\n')
	return cw.Out.Write(buf.Bytes())
}

func (cw ColumnWriter) ColumnNames() []string {
	names := make([]string, len(cw.Columns))
	for i, c := range cw.Columns {
		title := c.Title
		if title == "" {
			title = c.Name
		}
		names[i] = cases.Title(language.English).String(title)
	}
	return names
}

func (cw ColumnWriter) ColumnAtIndex(index int) Column {
	if index >= 0 && index < len(cw.Columns) {
		return cw.Columns[index]
	}
	return DefaultColumn
}

func (cw ColumnWriter) IndexOfColumn(name string) int {
	for i, c := range cw.Columns {
		if strings.EqualFold(c.Name, name) {
			return i
		}
	}
	return -1
}
