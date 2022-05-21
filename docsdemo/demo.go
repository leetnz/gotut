// This package demonstrates go docs.
//
// This will only be in the overview.
package docsdemo

import "fmt"

// BytesPerKb is the number of bytes in a kilobyte.
const BytesPerKb = 1024

// PrintKb prints the size in kb to the terminal.
func PrintKb(size int) {
	fmt.Printf("%d kB\n", size/BytesPerKb)
}
