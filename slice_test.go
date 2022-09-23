package slice_test

import (
	"fmt"
	"testing"

	"github.com/goosepkg/slice"
)

func Test_GetLatestElement(t *testing.T) {
	arr := []string{"test1", "test2", "test3"}

	latest := slice.GetLatestElement(arr)

	fmt.Println(latest)
}

func Test_Negative_GetLatestElement(t *testing.T) {
	arr := []string{}

	latest := slice.GetLatestElement(arr)

	fmt.Println(latest)
}
