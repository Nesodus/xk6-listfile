// Copyright 2022 Ilya Kobelev
// Use of this source code is governed by the Apache 2.0

package listfile

import (
	"log"
	"os"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/listfile", new(FILELIST))
}

// FILELIST is the k6 extension
type FILELIST struct{}

func (*FILELIST) GetFiles(path string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var slice []string

	for _, file := range files {
		slice = append(slice, file.Name())
	}
	return slice
}
