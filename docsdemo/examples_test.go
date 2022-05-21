package docsdemo_test

import "github.com/leetnz/gotut/docsdemo"

func Example() {
	docsdemo.PrintKb(2 * docsdemo.BytesPerKb)
	// Output:
	// 2 kB
}

func ExamplePrintKb() {
	docsdemo.PrintKb(5 * docsdemo.BytesPerKb)
	// Output:
	// 5 kB
}

// Prints one megabyte in kilobytes to the terminal.
func ExamplePrintKb_megabyte() {
	docsdemo.PrintKb(docsdemo.BytesPerKb * docsdemo.BytesPerKb)
	// Output:
	// 1024 kB
}
