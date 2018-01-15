package cmd

import (
	"github.com/gsantandrea/catbeat/beater"

	cmd "github.com/elastic/beats/libbeat/cmd"
)

// Name of this beat
var Name = "catbeat"

// RootCmd to handle beats cli
var RootCmd = cmd.GenRootCmd(Name, "", beater.New)
