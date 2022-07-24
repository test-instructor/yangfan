package hrp

import (
	"fmt"
	"os"
)

func DeleteHrpGOMOD() {
	fmt.Println("DeleteHrpGOMOD")
	err := os.Remove("hrp/internal/scaffold/templates/plugin/go.mod")
	if err != nil {
	}
	err = os.Remove("hrp/internal/scaffold/templates/plugin/go.sum")
	if err != nil {
	}
}
