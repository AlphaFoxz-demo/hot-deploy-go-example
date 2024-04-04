package generator

import (
	"github.com/fsnotify/fsnotify"
	"github.com/leaanthony/slicer"
	"os"
	"path/filepath"
)

func NewWatcher(dir string) (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	dirs, err := getSubdirectories(dir)
	if err != nil {
		return nil, err
	}
	for _, dir := range dirs.AsSlice() {
		err := watcher.Add(dir)
		if err != nil {
			return nil, err
		}
	}
	return watcher, nil
}

//func getAllFiles(rootDir string) (*slicer.StringSlicer, error) {
//	var result slicer.StringSlicer
//
//}

func getSubdirectories(rootDir string) (*slicer.StringSlicer, error) {
	var result slicer.StringSlicer

	// Iterate root dir
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// If we have a directory, save it
		if info.IsDir() {
			result.Add(path)
		}
		return nil
	})
	return &result, err
}
