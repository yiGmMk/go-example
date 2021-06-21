package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func md5path() {
	m, err := md5all(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x %s\n", m[path], path)
	}
}

func md5all(root string) (map[string][md5.Size]byte, error) {
	m := make(map[string][md5.Size]byte)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}

		return nil
	})
	if err != nil {
		return m, nil
	}
	return m, nil
}
