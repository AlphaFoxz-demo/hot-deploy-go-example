package generator

import (
	_ "embed"
	"fmt"
	"github.com/AlphaFoxz/hot-deploy-go-example/generator/handler"
	"github.com/AlphaFoxz/hot-deploy-go-example/generator/utils/customfs"
	"github.com/AlphaFoxz/hot-deploy-go-example/generator/utils/logutils"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

var extensionsToTriggerGen = make(map[string]bool)

func Listen(sourceDir string, targetDir string) {
	extensionsToTriggerGen["go"] = true
	watcher, err := NewWatcher(sourceDir)
	if err != nil {
		return
	}
	defer func(watcher *fsnotify.Watcher) {
		err := watcher.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(watcher)

	logutils.LogGreen("FileSystem Watching: %s", sourceDir)

	quit := false
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, os.Interrupt, syscall.SIGTERM)

	needGen := false
	exitCodeChannel := make(chan int, 1)
	interval := time.Duration(3000) * time.Millisecond
	timer := time.NewTimer(interval)

	for !quit {
		select {
		case exitCode := <-exitCodeChannel:
			if exitCode == 0 {
				quit = true
			}
		case err := <-watcher.Errors:
			logutils.LogDarkYellow(err.Error())
		case item := <-watcher.Events:
			if item.Op&fsnotify.Write == fsnotify.Write {
				if checkEligibleFile(item.Name) {
					fmt.Printf("file change detected: %s\n", item)
					needGen = true
					timer.Reset(interval)
					continue
				}
			}
			if item.Op&fsnotify.Create == fsnotify.Create {
				// If this is a folder, add it to our watch list
				if customfs.DirExists(item.Name) {
					err := watcher.Add(item.Name)
					if err != nil {
						logutils.LogRed("%s", err.Error())
					}
					logutils.LogGreen("Added new directory to watcher: %s", item.Name)
				} else if checkEligibleFile(item.Name) {
					needGen = true
					timer.Reset(interval)
					logutils.LogGreen("update file %s", item)
					continue
				}
			}
		case <-timer.C:
			if needGen {
				needGen = false
				logutils.LogDarkYellow("start to generate ...")
				handler.HandleCommand(sourceDir)
				handler.HandleApi(sourceDir)
				handler.HandleEvent(sourceDir)
			}
		}
	}
}

func checkEligibleFile(fileName string) bool {
	if strings.HasSuffix(fileName, "_gen.go") {
		return false
	}
	// Iterate all file patterns
	ext := filepath.Ext(fileName)
	if ext != "" {
		ext = ext[1:]
		return extensionsToTriggerGen[ext]
	}
	return false
}
