package main

type Uploader interface {
	Upload([]byte, string) (string, error)
}