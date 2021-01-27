package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/storage"
)

// main function
func main() {
	// Taking filepath param
	filepathPtr := flag.String("filepath", "", "Where the file to be uploaded.")

	// Taking bucket param
	bucketPtr := flag.String("bucket", "", "Your bucket name in Cloud Storage.")

	// Taking directory param
	directoryPtr := flag.String("directory", "", "Directory you want to save in Cloud Storage.")

	// Taking object param
	objectPtr := flag.String("object", "", "Object you want to save in Cloud Storage.")
	flag.Parse()

	filepath := *filepathPtr
	log.Printf("filepath: %s\n", filepath)

	bucket := *bucketPtr
	log.Printf("bucket: %s\n", bucket)

	directory := *directoryPtr
	log.Printf("directory: %s\n", directory)

	object := *objectPtr
	log.Printf("object: %s\n", object)

	uploadFile(nil, filepath, bucket, directory, object)
}

// uploadFile uploads an object
func uploadFile(w io.Writer, filepath, bucket, directory, object string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("storage.NewClient: %v", err)
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	// Open local file
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("os.Open: %v", err)
		return fmt.Errorf("os.Open: %v", err)
	}
	defer f.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Checking wheter directory is empty
	if !strings.EqualFold(directory, "") {
		object = directory + "/" + object
	}
	log.Println("directory + object: " + object)

	// Upload an object with storage.Writer
	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		log.Fatalf("io.Copy: %v", err)
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		log.Fatalf("Writer.Close: %v", err)
		return fmt.Errorf("Writer.Close: %v", err)
	}
	log.Printf("Blob %v uploaded.\n", object)
	return nil
}
