package store

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
)

type Csv struct {
	writer *csv.Writer
	file   *os.File
	mutex  sync.Mutex
}

func NewCsv(filepath string) (*Csv, error) {

	c := new(Csv)

	// Open the CSV file in append mode
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return &Csv{}, fmt.Errorf("NewCsv() %w", err)
	}

	c.file = file

	// Truncate the file to remove any existing content
	err = c.file.Truncate(0)
	if err != nil {
		return &Csv{}, fmt.Errorf("NewCsv() %w", err)
	}

	// Create a new CSV writer and write a row
	c.writer = csv.NewWriter(file)

	// add a mutex so multiple go routines can write to the same file
	c.mutex = sync.Mutex{}

	return c, nil

}

func (c *Csv) AddRow(row []string) error {

	fmt.Println("adding a row")

	c.mutex.Lock()
	defer c.mutex.Unlock()

	err := c.writer.Write(row)
	if err != nil {
		return fmt.Errorf("Csv.AddRow() %w", err)
	}

	c.writer.Flush()

	return nil

}

func (c *Csv) Close() {
	c.file.Close()
}
