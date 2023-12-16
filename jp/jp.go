package jp

import "fmt"

func Parse(data string) error {
	if data == "{}" {
		return nil
	}

	return fmt.Errorf("something went wrong")
}
