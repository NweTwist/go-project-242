package goproject242

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, humanReadable bool, includeHidden bool, recursive bool) int64 {
	var size int64

	err := filepath.WalkDir(path, func(
		p string,
		entry fs.DirEntry,
		err error,
	) error {
		if err != nil {
			return err
		}

		if !includeHidden && p != path &&
			strings.HasPrefix(entry.Name(), ".") {
			if entry.IsDir() {
				return fs.SkipDir
			}
			return nil
		}

		if !recursive && p != path && entry.IsDir() {
			return fs.SkipDir
		}

		if !entry.IsDir() {
			info, err := entry.Info()

			if err != nil {
				return err
			}
			size += info.Size()
		}

		return nil
	})

	if err != nil {
		return 0
	}

	return size
}
