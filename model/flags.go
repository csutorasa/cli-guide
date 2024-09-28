package model

import "fmt"

type Flags struct {
	RootDir string
	Quiet   bool
	Verbose bool
}

func (f *Flags) String() string {
	return fmt.Sprintf("rootDir %s, quiet %t, verbose %t", f.RootDir, f.Quiet, f.Verbose)
}
