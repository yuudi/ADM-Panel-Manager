package containers

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	containers, err := ListContainers()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%d containers", len(containers))
	for _, container := range containers {
		fmt.Println(container.Names)
	}
}
