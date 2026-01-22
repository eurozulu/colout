package colout

import (
	"bufio"
	"bytes"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io"
	"os"
	"strings"
)

type ColumnWriter struct {
	Output          io.Writer
	Columns         []Column
	StringDelimiter string
	ColumnSpacer    string
}

func (c ColumnWriter) Write(p []byte) (n int, err error) {
	out := c.Output
	if out == nil {
		out = os.Stdout
	}
	return out.Write(p)
}

func (c ColumnWriter) WriteString(s string) (n int, err error) {
	scn := bufio.NewScanner(strings.NewReader(s))
	buf := bytes.NewBuffer(nil)
	if c.StringDelimiter == "" {
		c.StringDelimiter = ","
	}
	for scn.Scan() {
		buf.WriteString(c.formatString(scn.Text()))
		buf.WriteString("\n")
	}
	return c.Write(buf.Bytes())
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

func (c ColumnWriter) columnAtIndex(i int) Column {
	if i < 0 || i >= len(c.Columns) {
		return DefaultColumn
	}
	return c.Columns[i]
}

func (c ColumnWriter) formatString(s string) string {
	buf := bytes.NewBuffer(nil)

	for i, sz := range strings.Split(s, c.StringDelimiter) {
		if i > 0 {
			buf.WriteString(c.ColumnSpacer)
		}
		sz = c.columnAtIndex(i).FormatString(sz)
		buf.WriteString(sz)
	}
	return buf.String()
}
