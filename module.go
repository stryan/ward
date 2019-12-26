package main

type Module interface {
	Init() bool
	Path() string
	Update() bool
	Output(...string) []byte
}
