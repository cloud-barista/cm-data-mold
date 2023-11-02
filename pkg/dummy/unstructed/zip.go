package unstructed

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/cloud-barista/cm-data-mold/pkg/utils"
)

// ZIP generation function using gofakeit
//
// CapacitySize is in GB and generates zip files
// within the entered dummyDir path.
func GenerateRandomZIP(dummyDir string, capacitySize int) error {
	dummyDir = filepath.Join(dummyDir, "zip")
	if err := utils.IsDir(dummyDir); err != nil {
		return err
	}

	tempPath := filepath.Join(dummyDir, "tmpTxt")
	if err := os.MkdirAll(tempPath, 0755); err != nil {
		return err
	}
	defer os.RemoveAll(tempPath)

	if err := GenerateRandomTXT(tempPath, 1); err != nil {
		return err
	}

	countNum := make(chan int, capacitySize)
	resultChan := make(chan error, capacitySize)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			randomZIPWorker(countNum, dummyDir, tempPath, resultChan)
		}()
	}

	for i := 0; i < capacitySize; i++ {
		countNum <- i
	}
	close(countNum)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for ret := range resultChan {
		if ret != nil {
			return ret
		}
	}

	return nil
}

// txt worker
func randomZIPWorker(countNum chan int, dummyDir, tempPath string, resultChan chan<- error) {
	for num := range countNum {
		w, err := os.Create(filepath.Join(dummyDir, fmt.Sprintf("datamold-dummy-data_%d.zip", num)))
		if err != nil {
			resultChan <- err
		}
		defer w.Close()

		zipWriter := zip.NewWriter(w)
		defer zipWriter.Close()

		if err := gzip(tempPath, zipWriter); err != nil {
			resultChan <- err
		}

		resultChan <- nil
	}
}

func gzip(srcDir string, zipWriter *zip.Writer) error {
	return filepath.Walk(srcDir, func(fp string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fileToZip, err := os.Open(fp)
			if err != nil {
				return err
			}
			defer fileToZip.Close()

			infoHeader, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}

			infoHeader.Name = filepath.Join(filepath.Base(srcDir), filepath.Base(fp))

			writer, err := zipWriter.CreateHeader(infoHeader)
			if err != nil {
				return err
			}

			_, err = io.Copy(writer, fileToZip)

			return err
		}
		return nil
	})
}
