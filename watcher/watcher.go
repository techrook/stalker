package watcher

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fsnotify/fsnotify"
)

func Watch(path string) error {
	log.Printf("🔍 Creating watcher for: %s", path)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	if err := watcher.Add(path); err != nil {
		return err
	}

	// Channel to listen for OS interrupt (e.g., Ctrl+C)
	done := make(chan bool)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Printf("📁 Event: %s", event)

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Printf("⚠️ Error: %s", err)
			case <-sig:
				log.Println("🛑 Termination signal received. Exiting watcher.")
				done <- true
				return
			}
		}
	}()

	<-done
	return nil
}
