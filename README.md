# Colout
### A simple column formatter for text output

Create Columns of a fixed width with optional names and Title text.  

Create A ColumnWriter using the Columns.

Use ColumnWriter `WriteStrings` (Note plural) method to write out slice of strings into columns.  

```aiignore

func main() {
	text := "abc123,22/12/2025, Joe, Soap, ¢233.55"
	cw := colout.ColumnWriter{
		Columns: []colout.Column{
			{
				Name:      "id",
				Alignment: colout.Right,
				Width:     10,
			},
			{
				Name:  "created",
				Width: 10,
			},
			{
				Name:  "first-name",
				Width: 25,
			},
			{
				Name:  "surname",
				Width: 25,
			},
			{
				Name:      "amount",
				Width:     10,
				Alignment: colout.Centre,
			},
		},
		ColumnSpacer: "  ",
		Out:          os.Stdout,
	}
	cw.WriteStrings(cw.ColumnNames())
	cw.WriteStrings(strings.Split(text, ","))
}

```  

Outputs:
```aiignore
        Id  Created     First-Name                 Surname                      Amount  
    abc123  22/12/2025  Joe                        Soap                        ¢233.55 

```