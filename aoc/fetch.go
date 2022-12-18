package aoc

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	tokenFile = "token"
	inputFile = "input"
)

type fetcher struct {
	dir    string
	client *http.Client
}

func (f *fetcher) fetchInput(year, day int) io.ReadSeeker {
	filename := filepath.Join(f.dir, strconv.Itoa(year), strconv.Itoa(day), inputFile)

	file, err := os.Open(filename)
	if err == nil {
		return file
	}
	if !errors.Is(err, os.ErrNotExist) {
		log.Panicf("failed to read cached input for %d/%d: %v", year, day, err)
	}

	log.Printf("⬇️ downloading input for %d/%d", year, day)

	data := f.requestInput(year, day)
	defer func() {
		_ = data.Close()
	}()

	file, err = f.cacheInput(filename, data)
	if err != nil {
		log.Panicf("failed to open input cache for %d/%d: %v", year, day, err)
	}

	log.Printf("✔ input for %d/%d downloaded and cached successfully", year, day)

	return file
}

func (f *fetcher) requestInput(year, day int) io.ReadCloser {
	tokenPath, err := filepath.Abs(filepath.Join(f.dir, tokenFile))
	if err != nil {
		log.Panicf("failed to read token file: %v", err)
	}

	token, err := os.ReadFile(tokenPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Panicf("missing token file; you must store your AoC session token in file %s", tokenPath)
		} else {
			log.Panicf("failed to read token file: %v", err)
		}
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Panicf("failed to fetch input for %d/%d: %v", year, day, err)
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: strings.TrimSpace(string(token))})

	client := f.client
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicf("failed to fetch input for %d/%d: %v", year, day, err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Panicf("failed to fetch input for %d/%d: %s", year, day, resp.Status)
	}
	return resp.Body
}

func (f *fetcher) cacheInput(filename string, data io.Reader) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(filename), 0o755); err != nil {
		return nil, err
	}

	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(file, data); err != nil {
		return nil, err
	}
	if err := file.Sync(); err != nil {
		return nil, err
	}
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}
	return file, nil
}
