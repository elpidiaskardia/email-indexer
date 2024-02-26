package main

import (
	"Api_Go/db/functions"
	"Api_Go/models"
	"flag"
	"log"
	"os"
	"path/filepath"
	"runtime/pprof"
	"sync"
)

func main() {
	flag.Parse()
	f, err := os.Create("profile.pb.gz")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	indexerData, err := functions.CreateIndexerFromJsonFile("./indexs/mails.json")
	if err != nil {
		log.Fatalf("Error reading indexer data: %v", err)
	}

	log.Println("Deleting existing index...")
	if err := functions.DeleteIndexOnZincSearch("emails"); err != nil {
		log.Printf("No existing index found: %v", err)
	}

	log.Println("Creating new index...")
	if err := functions.CreateIndexOnZincSearch(indexerData); err != nil {
		log.Fatalf("Error creating index: %v", err)
	}

	log.Println("Start indexing, this might take a few minutes...")
	var batchSize = 1000
	var records []models.EmailData
	var mu sync.Mutex
	var wg sync.WaitGroup
	maxWorkers :=  500
	semaphore := make(chan struct{}, maxWorkers)

	err = filepath.Walk("./maildir/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			semaphore <- struct{}{}
			wg.Add(1)
			go func(p string) {
				emailData, err := functions.ProcessFile(p)
				if err != nil {
					log.Printf("Error processing file %s: %v", p, err)
					return
				}
				mu.Lock()
				records = append(records, emailData)
				if len(records) >= batchSize {
					functions.SendBulkToZincSearch(records)
					records = nil
				}
				defer mu.Unlock() 
				defer func() {
					<-semaphore
					wg.Done()
				}()
			}(path)
		}
		return nil
	})
	wg.Wait()
	
	if len(records) > 0 {
		functions.SendBulkToZincSearch(records)
	}

	if err != nil {
		log.Fatalf("Error walking directory: %v", err)
	}
	log.Printf("Indexing completed")
}
