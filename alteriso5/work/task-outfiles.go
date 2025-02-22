package work

import "github.com/FascodeNet/alterlinux/alteriso5/work/boot"

var makeOutFiles *BuildTask = NewBuildTask("makeOutFiles", func(w Work) error {
	return boot.Xorriso.Build(w.Dirs.Iso, w.target.Out, w.profile.BootModes...)
})
