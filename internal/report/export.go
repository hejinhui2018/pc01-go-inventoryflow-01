package report

import (
	"encoding/csv"
	"io"
	"strconv"
)

func WriteCSV(w io.Writer, rows []Row) error {
	c := csv.NewWriter(w)
	if err := c.Write([]string{"sku", "warehouse", "on_hand", "reserved", "available"}); err != nil {
		return err
	}
	for _, r := range rows {
		if err := c.Write([]string{r.SKU, r.Warehouse, strconv.FormatInt(r.OnHand, 10), strconv.FormatInt(r.Reserved, 10), strconv.FormatInt(r.Available, 10)}); err != nil {
			return err
		}
	}
	c.Flush()
	return c.Error()
}
