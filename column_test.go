package colout

import (
	"fmt"
	"strings"
	"testing"
)

const smallValue = "abc"
const midValue = "abcdefghij"
const bigValue = "abcdefghijklmnopqrstuvwxyz1234567890"

func TestColumn_FormatStringWidth(t *testing.T) {
	if err := testColWidth(smallValue, 10, paddedLeftString(smallValue, 10)); err != nil {
		t.Fatal(err)
	}
	if err := testColWidth(midValue, 10, midValue); err != nil {
		t.Fatal(err)
	}
	if err := testColWidth(bigValue, 10, bigValue[:7]+"..."); err != nil {
		t.Fatal(err)
	}

	if err := testColWidth(smallValue, 40, paddedLeftString(smallValue, 40)); err != nil {
		t.Fatal(err)
	}
	if err := testColWidth(midValue, 40, paddedLeftString(midValue, 40)); err != nil {
		t.Fatal(err)
	}
	if err := testColWidth(bigValue, 40, paddedLeftString(bigValue, 40)); err != nil {
		t.Fatal(err)
	}

	if err := testColWidth(smallValue, 1, smallValue[:1]); err != nil {
		t.Fatal(err)
	}
	if err := testColWidth(midValue, 1, midValue[:1]); err != nil {
		t.Fatal(err)
	}
	if err := testColWidth(bigValue, 1, bigValue[:1]); err != nil {
		t.Fatal(err)
	}

	if err := testColWidth(smallValue, -1, smallValue); err != nil {
		t.Fatal(err)
	}
	if err := testColWidth(midValue, -1, midValue); err != nil {
		t.Fatal(err)
	}
	if err := testColWidth(bigValue, -1, bigValue); err != nil {
		t.Fatal(err)
	}

	if err := testColWidth(smallValue, 0, ""); err != nil {
		t.Fatal(err)
	}
	if err := testColWidth(midValue, 0, ""); err != nil {
		t.Fatal(err)
	}
	if err := testColWidth(bigValue, 0, ""); err != nil {
		t.Fatal(err)
	}

}

func TestColumn_FormatStringAlignment(t *testing.T) {
	if err := testColAlign(smallValue); err != nil {
		t.Fatal(err)
	}
	if err := testColAlign(midValue); err != nil {
		t.Fatal(err)
	}
	if err := testColAlign(bigValue); err != nil {
		t.Fatal(err)
	}
}

func testColAlign(s string) error {
	col := Column{
		Name:  "",
		Width: 40,
	}
	for _, align := range []ColumnAlignment{Left, Right, Centre} {
		col.Alignment = align
		var expect string
		switch align {
		case Left:
			expect = paddedLeftString(s, col.Width)
		case Right:
			expect = paddedRightString(s, col.Width)
		case Centre:
			expect = paddedMidString(s, col.Width)
		}
		result := col.FormatString(s)
		if result != expect {
			return fmt.Errorf("unexpected format string with column alignment %s. %q, got %q", align, expect, result)
		}
	}
	return nil
}

func testColWidth(s string, width int, expect string) error {
	expect = paddedLeftString(expect, width)
	col := Column{
		Name:      "",
		Alignment: Left,
		Width:     width,
	}
	result := col.FormatString(s)
	if result != expect {
		return fmt.Errorf("unexpected format string with column width %d. %q, got %q", width, expect, result)
	}
	return nil
}

func paddedLeftString(s string, width int) string {
	size := width - len(s)
	if size <= 0 {
		return s
	}
	return strings.Join([]string{s, strings.Repeat(" ", size)}, "")
}
func paddedRightString(s string, width int) string {
	size := width - len(s)
	if size <= 0 {
		return s
	}
	return strings.Join([]string{strings.Repeat(" ", size), s}, "")
}
func paddedMidString(s string, width int) string {
	padSize := width - len(s)
	if padSize <= 0 {
		return s
	}
	l := strings.Repeat(" ", padSize/2)
	r := strings.Repeat(" ", padSize-(padSize/2))
	return strings.Join([]string{l, s, r}, "")
}
