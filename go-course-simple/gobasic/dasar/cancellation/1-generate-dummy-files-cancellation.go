package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type FileInfo struct {
	Index       int
	FileName    string
	WorkerIndex int
	Err         error
}

const totalFile = 3000
const contentLength = 5000
const timeoutDuration = 3 * time.Second

var tempPath = filepath.Join(os.Getenv("TEMP"), "module-pipeline-cancellation-context")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.Println("start")
	start := time.Now()

	// with timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()
	generateFilesWithContext(ctx)

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	// time.AfterFunc(timeoutDuration, cancel)
	// generateFilesWithContext(ctx)

	// without timeout
	generateFiles()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func randomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func generateFiles() {
	generateFilesWithContext(context.Background())
	// os.RemoveAll(tempPath)
	// os.MkdirAll(tempPath, os.ModePerm)

	// // pipeline 1: job distribution
	// chanFileIndex := generateFilesIndexes(context.Background())

	// // pipeline 2: the main logic (creating files)
	// createFilesWorker := 100
	// chanFileResult := createFiles(context.Background(), chanFileIndex, createFilesWorker)

	// // track and print output
	// counterTotal := 0
	// counterSuccess := 0

	// for fileResult := range chanFileResult {
	// 	if fileResult.Err != nil {
	// 		log.Printf("error creating file %s. stack trace: %s", fileResult.FileName, fileResult.Err)
	// 	} else {
	// 		counterSuccess++
	// 	}
	// 	counterTotal++
	// }
	// log.Printf("%d/%d of total files created", counterSuccess, counterTotal)
}

func generateFilesIndexes(ctx context.Context) <-chan FileInfo {
	chanOut := make(chan FileInfo)
	go func() {
		for i := 0; i < totalFile; i++ {
			select {
			case <-ctx.Done():
				break
			default:
				chanOut <- FileInfo{
					Index:    i,
					FileName: fmt.Sprintf("file-%d.txt", i),
				}
			}
		}
		close(chanOut)
	}()
	return chanOut
}

func createFiles(ctx context.Context, chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	wg := new(sync.WaitGroup)
	wg.Add(numberOfWorkers)

	go func() {
		for workerIndex := 0; workerIndex < numberOfWorkers; workerIndex++ {
			go func(workerIndex int) {
				for job := range chanIn {
					select {
					case <-ctx.Done():
						break
					default:
						filePath := filepath.Join(tempPath, job.FileName)
						content := randomString(contentLength)
						err := ioutil.WriteFile(filePath, []byte(content), os.ModePerm)

						log.Println("worker", workerIndex, "working on", job.FileName, "file generation")

						chanOut <- FileInfo{
							FileName:    job.FileName,
							WorkerIndex: workerIndex,
							Err:         err,
						}
					}
				}

				wg.Done()
			}(workerIndex)
		}
	}()

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}

func generateFilesWithContext(ctx context.Context) {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)
	done := make(chan int)
	go func() {
		// pipeline 1: job distribution
		chanFileIndex := generateFilesIndexes(ctx)

		// pipeline 2: the main logic (creating files)
		createFilesWorker := 100
		chanFileResult := createFiles(ctx, chanFileIndex, createFilesWorker)

		// track and print output
		counterSuccess := 0
		for fileResult := range chanFileResult {
			if fileResult.Err != nil {
				log.Printf("error creating file %s. stack trace: %s", fileResult.FileName, fileResult.Err)
			} else {
				counterSuccess++
			}
		}
		// notify that the process is complete
		done <- counterSuccess
	}()

	select {
	case <-ctx.Done():
		log.Printf("generation process stopped. %s", ctx.Err())
	case counterSuccess := <-done:
		log.Printf("%d/%d of total files created", counterSuccess, totalFile)
	}
}

// go run 1-generate-dummy-files-cancellation.go
