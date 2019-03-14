package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	_, err := ioutil.ReadAll(strings.NewReader("test"))
	fmt.Println(errors.Wrap(err, "read failed"))
}
