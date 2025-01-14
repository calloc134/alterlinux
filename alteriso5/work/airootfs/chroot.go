package airootfs

import (
	"log/slog"
	"os"
	"os/exec"
)

type Chroot struct {
	Arch       string
	Dir        string
	Config     string
	initilized bool
}

func GetChrootDir(dir, arch, config string) (*Chroot, error) {
	env := Chroot{
		Arch:   arch,
		Dir:    dir,
		Config: config,
	}

	entry, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			env.initilized = false
			return &env, nil
		} else {
			return nil, err
		}
	}

	if len(entry) > 0 {
		env.initilized = true
	}

	return &env, nil
}

func (e *Chroot) Init(pkgs ...string) error {
	if err := os.MkdirAll(e.Dir, 0755); err != nil {
		return err
	}

	args := []string{"-c", "-C", e.Config, e.Dir}
	args = append(args, pkgs...)

	slog.Debug("pacstrap", "args", args)

	pacstrap := exec.Command("pacstrap", args...)
	pacstrap.Env = append(os.Environ(), "LANG=C")
	pacstrap.Stdout = os.Stdout
	pacstrap.Stderr = os.Stderr
	if err := pacstrap.Run(); err != nil {
		return err
	}

	e.initilized = true

	return nil
}

type kernel struct {
	Linux  string
	Initrd string
}

func (k *kernel) Files() []string {
	return []string{k.Linux, k.Initrd}
}

// func (e *Chroot) FindKernels() ([]kernel, error) {
// 	kernels := []kernel{}

// 	presetsDir := path.Join(e.Dir, "etc", "mkinitcpio.d")
// 	entry, err := os.ReadDir(presetsDir)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, e := range entry {
// 		if e.IsDir() || !strings.HasSuffix(e.Name(), ".preset") {
// 			continue
// 		}

// 		fp := path.Join(presetsDir, e.Name())
// 		env, err := utils.LoadEnvFile(fp)
// 		if err != nil {
// 			continue
// 		}

// 		ker := env["ALL_kver"]
// 		initrd := env["default_image"]

// 		if ker != "" && initrd != "" {
// 			kernels = append(kernels, kernel{
// 				Linux:  ker,
// 				Initrd: initrd,
// 			})
// 		}

// 	}

// 	slog.Debug("FindKernels:", "kernels", kernels)
// 	return kernels, nil
// }

func (e *Chroot) FindKernels() ([]kernel, error) {
	return []kernel{
		{
			Linux:  "/boot/vmlinuz-linux",
			Initrd: "/boot/initramfs-linux.img",
		},
	}, nil
}
