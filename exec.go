// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gccxml

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var (
	// Modify this variable if gccxml is in another path.
	GccXmlCmd = "gccxml"
	GccCmd    = "gcc"
)

func New(file string) Xml {
	return Xml{
		File:        file,
		CFlags:      nil,
		IncludeDirs: nil,
	}
}

type Xml struct {
	File        string
	CFlags      []string // list of CFlags to pass to gccxml
	IncludeDirs []string // list of include directories to pass to gccxml
}

func (g Xml) cmdArgs(args ...string) []string {
	cmdargs := make([]string, 0, 1+len(args)+len(g.CFlags)+len(g.IncludeDirs))
	if len(args) > 0 {
		cmdargs = append(cmdargs, args...)
	}
	cmdargs = append(cmdargs, g.File)
	if len(g.CFlags) > 0 {
		cmdargs = append(cmdargs, g.CFlags...)
	}
	if len(g.IncludeDirs) > 0 {
		for _, dir := range g.IncludeDirs {
			cmdargs = append(cmdargs, "-I"+dir)
		}
	}
	return cmdargs
}

func (g Xml) dumpCmd() cmd {
	args := g.cmdArgs("-fxml=/dev/stdout")
	return newCmd(GccXmlCmd, args...)
}

func (g Xml) macroCmd() cmd {
	args := g.cmdArgs("--preprocess", "-dM")
	return newCmd(GccXmlCmd, args...)
}

func (g Xml) Doc() (gccxml *XmlDoc, err error) {
	err = g.dumpCmd().run(func(r io.Reader) error {
		return xml.NewDecoder(r).Decode(&gccxml)
	})
	if err != nil {
		return nil, err
	}
	gccxml.file = g.File
	gccxml.prepare()
	return
}

func (g Xml) Save(w io.Writer) error {
	return g.dumpCmd().run(func(r io.Reader) error {
		_, err := io.Copy(w, r)
		return err
	})
}

func (g Xml) PrintMacros() error {
	return g.macroCmd().run(func(r io.Reader) error {
		_, err := io.Copy(os.Stdout, r)
		return err
	})
}

func (g Xml) Macros() (ms Macros, err error) {
	err = g.macroCmd().run(func(r io.Reader) error {
		ms, err = DecodeMacros(r)
		return err
	})
	return
}

func IncludeFiles(file string) (fs []string, err error) {
	err = newCmd(GccCmd, "-M", "-MT", "''", file).run(func(r io.Reader) error {
		buf, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}
		for _, line := range strings.Split(strings.Replace(string(buf), `\`, " ",
			-1), " ") {
			line = strings.TrimSpace(line)
			if line != "" {
				fs = append(fs, line)
			}
		}
		return nil
	})
	return
}

type cmd struct {
	*exec.Cmd
}

func newCmd(name string, arg ...string) cmd {
	return cmd{exec.Command(name, arg...)}
}

func (c cmd) run(visit func(r io.Reader) error) error {
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	defer stdout.Close()
	stderr, err := c.StderrPipe()
	if err != nil {
		return err
	}
	defer stderr.Close()
	if err := c.Start(); err != nil {
		return err
	}
	go io.Copy(os.Stderr, stderr)
	if err := visit(stdout); err != nil {
		return err
	}
	if err := c.Wait(); err != nil {
		return err
	}
	return nil
}
