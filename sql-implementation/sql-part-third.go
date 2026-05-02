package sql_implementation

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
)

const (
	COLUMN_USERNAME_SIZE = 32
	COLUMN_EMAIL_SIZE    = 255
	PAGE_SIZE            = 4096
	TABLE_MAX_PAGES      = 100
)

type Row struct {
	ID       uint32
	Username [COLUMN_USERNAME_SIZE]byte
	Email    [COLUMN_EMAIL_SIZE]byte
}

const (
	ID_SIZE       = 4
	USERNAME_SIZE = COLUMN_USERNAME_SIZE
	EMAIL_SIZE    = COLUMN_EMAIL_SIZE

	ID_OFFSET       = 0
	USERNAME_OFFSET = ID_OFFSET + ID_SIZE
	EMAIL_OFFSET    = USERNAME_OFFSET + USERNAME_SIZE

	ROW_SIZE = ID_SIZE + USERNAME_SIZE + EMAIL_SIZE
)

type Table struct {
	numRows uint32
	pages   [TABLE_MAX_PAGES][]byte
}

func NewTable() *Table {
	return &Table{}
}

func rowSlot(table *Table, rowNum uint32) []byte {
	pageNum := rowNum / (PAGE_SIZE / ROW_SIZE)

	if table.pages[pageNum] == nil {
		table.pages[pageNum] = make([]byte, PAGE_SIZE)
	}

	rowOffset := rowNum % (PAGE_SIZE / ROW_SIZE)
	byteOffset := rowOffset * ROW_SIZE

	return table.pages[pageNum][byteOffset : byteOffset+ROW_SIZE]
}

func serializeRow(row *Row, dest []byte) {
	binary.LittleEndian.PutUint32(dest[ID_OFFSET:], row.ID)

	copy(dest[USERNAME_OFFSET:], row.Username[:])
	copy(dest[EMAIL_OFFSET:], row.Email[:])
}

func deserializeRow(src []byte) Row {
	var row Row

	row.ID = binary.LittleEndian.Uint32(src[ID_OFFSET:])
	copy(row.Username[:], src[USERNAME_OFFSET:USERNAME_OFFSET+USERNAME_SIZE])
	copy(row.Email[:], src[EMAIL_OFFSET:EMAIL_OFFSET+EMAIL_SIZE])

	return row
}

func executeInsert(table *Table, row Row) {
	slot := rowSlot(table, table.numRows)
	serializeRow(&row, slot)
	table.numRows++
}

func executeSelect(table *Table) {
	for i := uint32(0); i < table.numRows; i++ {
		slot := rowSlot(table, i)
		row := deserializeRow(slot)

		fmt.Printf("(%d, %s, %s)\n",
			row.ID,
			bytesToString(row.Username[:]),
			bytesToString(row.Email[:]),
		)
	}
}

func bytesToString(b []byte) string {
	n := 0
	for i, v := range b {
		if v == 0 {
			break
		}
		n = i + 1
	}
	return string(b[:n])
}

func run() {
	table := NewTable()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("db > ")
		input, _ := reader.ReadString('\n')

		var id uint32
		var username, email string

		if n, _ := fmt.Sscanf(input, "insert %d %s %s", &id, &username, &email); n == 3 {
			var row Row
			row.ID = id
			copy(row.Username[:], username)
			copy(row.Email[:], email)

			executeInsert(table, row)
			fmt.Println("Executed.")
			continue
		}

		if input[:6] == "select" {
			executeSelect(table)
			continue
		}

		fmt.Println("Unrecognized command")
	}
}
