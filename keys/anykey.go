package keys

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
)

func ExitOnAnyKey() {
	if err := keyboard.Open(); err != nil {
		fmt.Printf("Error attaching keyboard for input: %v\n", err)
		return
	}

	defer func() {
		_ = keyboard.Close()
	}()

	go func() {
		_, _, _ = keyboard.GetSingleKey()
		os.Exit(0)
	}()
}
