// See https://magefile.org/

//go:build mage

// Build steps for the acceptable API:
package main

import (
	"github.com/magefile/mage/sh"
)

var Default = Build

func Build() error {
	if err := sh.RunV("go", "test", "./..."); err != nil {
		return err
	}
	if err := sh.RunV("gofmt", "-l", "-w", "-s", "."); err != nil {
		return err
	}
	if err := sh.RunV("go", "vet", "./..."); err != nil {
		return err
	}
	return nil
}

func Coverage() error {
	if err := sh.RunV("go", "test", "-cover", "./...", "-coverprofile", "coverage.out", "-coverpkg", "./..."); err != nil {
		return err
	}
	if err := sh.RunV("go", "tool", "cover", "-func", "coverage.out", "-o", "report.out"); err != nil {
		return err
	}
	return nil
}
