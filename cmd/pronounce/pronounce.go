package pronounce

import (
	"flag"
	"fmt"
)

func WordPronunciation(cmd *flag.FlagSet, w *string) {
	fmt.Printf("%s \n", string(*w))
	fmt.Println("Coming soon!")
}
