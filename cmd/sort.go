package cmd

import (
	"fmt"
	"os"
	"wb_lvl2_sort/internal/sorter"

	"github.com/spf13/cobra"
)

var (
	sortColumn uint32
	reverse    bool
	numeric    bool
	unique     bool
)

var sortCmd = &cobra.Command{
	Use: "sort",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := sorter.ReadLines()
		if err != nil {
			return err
		}

		result := sorter.Sort(data, sorter.Options{
			Reverse:    reverse,
			Nuneric:    numeric,
			Unique:     unique,
			SortColumn: sortColumn,
		})

		for _, line := range result {
			_, err = fmt.Fprintln(os.Stdout, line)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(sortCmd)

	sortCmd.Flags().BoolVarP(&reverse, "reverse", "r", false, "сортировка в обратном порядке")
	sortCmd.Flags().BoolVarP(&numeric, "numeric", "n", false, "сортировка по числовому значению")
	sortCmd.Flags().BoolVarP(&unique, "unique", "u", false, "сортировка только уникальных значений")

	sortCmd.Flags().Uint32VarP(&sortColumn, "column", "k", 1, "номер столбца для сортировки")
}
