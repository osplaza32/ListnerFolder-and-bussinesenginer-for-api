package main

import (
	"fmt"
	"github.com/subosito/gotenv"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/radovskyb/watcher"
)

func main() {
	gotenv.Load(".env.gatewaydev")

	w := watcher.New()

	// SetMaxEvents to 1 to allow at most 1 event's to be received
	// on the Event channel per watching cycle.
	//
	// If SetMaxEvents is not set, the default is to send all events.

	// Only notify rename and move events.
	w.FilterOps(watcher.Rename, watcher.Move,watcher.Create,watcher.Write,watcher.Remove,watcher.Chmod)

	// Only files that match the regular expression during file listings
	// will be watched.
	var r = regexp.MustCompile(`(?m)([a-zA-Z0-9\s_\\.\-\(\):])+(.xml)$`)
	w.AddFilterHook(watcher.RegexFilterHook(r, false))

	go func() {
		for {
			select {
			case event := <-w.Event:
				PayOrDie(event)
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch this folder for changes.
	if err := w.Add(os.Getenv("ENV_CLONE")+string(os.PathSeparator)); err != nil {
		log.Fatalln(err)
	}

	// Watch test_folder recursively for changes.
	if err := w.AddRecursive(os.Getenv("ENV_CLONE")+string(os.PathSeparator)+"RootNode"); err != nil {
		log.Fatalln(err)
	}

	// Print a list of all of the files and folders currently
	// being watched and their paths.

	fmt.Println()

	// Trigger 2 events after watcher started.
	go func() {



	}()

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 10); err != nil {
		log.Fatalln(err)
	}
}

func PayOrDie(event watcher.Event) {
	fmt.Println(event.Path)
	fmt.Println(event.Name())
	fmt.Println(event.FileInfo)
	fmt.Println(event.Mode())
	fmt.Println(event.Op)






}