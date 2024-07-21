package layouts

import "fmt"

func Asset(name string) string {
	return fmt.Sprintf("../css/%s", name)
}
