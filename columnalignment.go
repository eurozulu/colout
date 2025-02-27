package colout

// ColumnAlignment represents where in a column the text will be places when there are less caharacters than column width space.
type ColumnAlignment byte

const (
	Left ColumnAlignment = iota
	Right
	Centre
)

var columnAlignmentNames = [...]string{
	"Left",
	"Right",
	"Centre",
}

func (c ColumnAlignment) String() string {
	i := int(c)
	if i < 0 || i >= len(columnAlignmentNames) {
		return ""
	}
	return columnAlignmentNames[i]
}
