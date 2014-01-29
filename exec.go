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

type Xml struct {
	File string
	Args []string
}

func (g Xml) dumpCmd() cmd {
	return g.newCmd(GccXmlCmd, "-fxml=/dev/stdout", g.File)
}

func (g Xml) macroCmd() cmd {
	return g.newCmd(GccXmlCmd, "--preprocess", "-dM", g.File)
}

func (g Xml) newCmd(name string, arg ...string) cmd {
	return newCmd(name, append(g.Args, arg...)...)
}

func (g Xml) Doc() (gccxml *XmlDoc, err error) {
	err = g.dumpCmd().read(func(r io.Reader) error {
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
	return g.dumpCmd().read(func(r io.Reader) error {
		_, err := io.Copy(w, r)
		return err
	})
}

func (g Xml) PrintMacros() error {
	return g.macroCmd().read(func(r io.Reader) error {
		_, err := io.Copy(os.Stdout, r)
		return err
	})
}

func (g Xml) Macros() (ms Macros, err error) {
	err = g.macroCmd().read(func(r io.Reader) error {
		ms, err = DecodeMacros(r)
		return err
	})
	return
}

func IncludeFiles(file string) (fs []string, err error) {
	err = newCmd(GccCmd, "-M", "-MT", "''", file).read(func(r io.Reader) error {
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

func (c cmd) read(visit func(r io.Reader) error) error {
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

func (c cmd) exec() error {
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
	go io.Copy(os.Stdout, stdout)
	if err := c.Wait(); err != nil {
		return err
	}
	return nil

}
