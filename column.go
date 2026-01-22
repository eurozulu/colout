package colout

import "strings"

// Column represents a single vertical line of text with a fixed width.
// Strings in this column which are greater than the column width are trimmed down.
// Strings smaller are spaced according to the Column Alignment value.
type Column struct {
	// Name of the column by which it may be referenced.
	Name string
	// Title is an optional string used for display of the name.
	Title string
	// Alignment sets where strings are placed in the column when shorter than the column width.
	Alignment ColumnAlignment
	// Width is the number of characters wide the column is.
	// A negative width == non fixed/unlimited width. (Usually reserved for the last column only.)
	// Note: Alignment is ignored when width is zero or less
	Width int
}

var DefaultColumn = Column{
	Name:      "",
	Alignment: Left,
	Width:     -1,
}

func (col Column) FormatString(s string) string {
	if col.Width < 0 {
		return s
	}
	if col.Width == 0 {
		return ""
	}

	s = strings.TrimSpace(s)
	padSize := col.Width - len(s)
	if padSize == 0 {
		return s
	}
	if padSize < 0 {
		// trim if string is too long
		ss := s[:(len(s)+padSize)-len("...")]
		if len(ss) > 4 {
			ss = strings.Join([]string{ss[:len(ss)-3], "..."}, "")
		}
		return ss
	}
	pad := strings.Repeat(" ", padSize)
	switch col.Alignment {
	case Left:
		return strings.Join([]string{s, pad}, "")
	case Centre:
		lpad := pad[:padSize/2]
		rpad := pad[len(lpad):]
		return strings.Join([]string{lpad, s, rpad}, "")
	case Right:
		return strings.Join([]string{pad, s}, "")
	}
	return s
}
