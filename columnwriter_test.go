package colout

import (
	"bytes"
	"strings"
	"testing"
)

func TestColumnWriter_ColumnAtIndex(t *testing.T) {
	testColumns := createTestColumns()
	cw := &ColumnWriter{
		Columns: testColumns,
	}
	c := cw.ColumnAtIndex(0)
	if c != testColumns[0] {
		t.Errorf("ColumnAtIndex(0) = %v; want %v", c, testColumns[0])
	}
	c = cw.ColumnAtIndex(3)
	if c != testColumns[3] {
		t.Errorf("ColumnAtIndex(3) = %v; want %v", c, testColumns[3])
	}
	c = cw.ColumnAtIndex(4)
	if c != DefaultColumn {
		t.Errorf("ColumnAtIndex(4) = %v; want %v", DefaultColumn, c)
	}
	c = cw.ColumnAtIndex(-1)
	if c != DefaultColumn {
		t.Errorf("ColumnAtIndex(4) = %v; want %v", DefaultColumn, c)
	}
}

func TestColumnWriter_ColumnNames(t *testing.T) {
	testColumns := createTestColumns()
	cw := &ColumnWriter{
		Columns: testColumns,
	}
	names := cw.ColumnNames()
	if len(names) != len(testColumns) {
		t.Errorf("len(ColumnNames()) = %d; want %d", len(names), len(testColumns))
	}
	for i, c := range testColumns {
		if names[i] != c.Title {
			t.Errorf("ColumnNames()[%d] = %v; want %v", i, names[i], c.Name)
		}
	}
}

func TestColumnWriter_WriteString(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	cw := &ColumnWriter{
		Columns:      createTestColumns(),
		ColumnSpacer: "",
		Out:          buf,
	}
	if _, err := cw.WriteString("a,b,c,d,e"); err != nil {
		t.Errorf("WriteString() error = %v; want nil", err)
	}
	if buf.String() != "a,b,c,d,e" {
		t.Errorf("buf.String() = %q; want %q", buf.String(), "a,b,c,d,e\n")
	}
}

func TestColumnWriter_WriteStrings(t *testing.T) {
	testColumns := createTestColumns()
	buf := bytes.NewBuffer(nil)
	cw := &ColumnWriter{
		Columns:      testColumns,
		ColumnSpacer: "",
		Out:          buf,
	}
	buf.Reset()
	if _, err := cw.WriteStrings(strings.Split("a,b,c,d,e", ",")); err != nil {
		t.Errorf("WriteString() error = %v; want nil", err)
	}
	if buf.String() != "a      b      cde                   \n" {
		t.Errorf("buf.String() = %q; want %q", buf.String(), "a      b      cde                   \n")
	}

	buf.Reset()
	cw.Columns = setColumnsAlignment(testColumns, Centre)
	if _, err := cw.WriteStrings(strings.Split("a,b,c,d,e", ",")); err != nil {
		t.Errorf("WriteString() error = %v; want nil", err)
	}
	if buf.String() != "  a    b    c  de                   \n" {
		t.Errorf("buf.String() = %q; want %q", buf.String(), "  a    b    c  de                   \n")
	}

	buf.Reset()
	cw.Columns = setColumnsAlignment(testColumns, Right)
	if _, err := cw.WriteStrings(strings.Split("a,b,c,d,e", ",")); err != nil {
		t.Errorf("WriteString() error = %v; want nil", err)
	}
	if buf.String() != "    a    b    cde                   \n" {
		t.Errorf("buf.String() = %q; want %q", buf.String(), "    a    b    cde                   \n")
	}
}

func setColumnsAlignment(cols []Column, alignment ColumnAlignment) []Column {
	colz := make([]Column, len(cols))
	for i, col := range cols {
		colz[i] = Column{
			Name:      col.Name,
			Title:     col.Title,
			Alignment: alignment,
			Width:     col.Width,
		}
	}
	return colz
}

func createTestColumns() []Column {
	return []Column{
		{
			Name:      "one",
			Title:     "One",
			Alignment: Left,
			Width:     5,
		},
		{
			Name:      "two",
			Title:     "Two",
			Alignment: Centre,
			Width:     5,
		},
		{
			Name:      "three",
			Title:     "Three",
			Alignment: Right,
			Width:     5,
		},
		{
			Name:      "four",
			Title:     "Four",
			Alignment: Right,
			Width:     -1,
		},
	}
}
