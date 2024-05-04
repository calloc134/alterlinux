package pkg

import (
	"log/slog"
	"os"
	"path"

	"github.com/Hayao0819/nahi/osutils"
)

// 指定されたプロファイルとアーキテクチャに対応するパッケージリストファイルを検索し、そのファイルのパスを返します。
func FindPkgListFiles(profile string, arch string) ([]string, error) {
	findFiles := []string{
		"packages." + arch,
		"packages.any",
	}

	findDirs := []string{
		"packages." + arch + ".d",
		"packages.any.d",
	}

	// パッケージリストファイルを検索
	for _, d := range findDirs {
		dir := path.Join(profile, d)
		slog.Debug("Check pkglist", "subdir", dir)
		if !osutils.IsDir(dir) {
			continue
		}

		slog.Debug("Found pkglist", "subdir", dir)

		files, err := os.ReadDir(dir)
		if err != nil {
			slog.Warn("Failed to list directory", "dir", dir, "error", err)
			continue
		}
		slog.Debug("Found pkglist", "files", files)
		for _, f := range files {
			p := path.Join(profile, d, f.Name())
			slog.Info("Found pkglist", "file",p)
			findFiles = append(findFiles, p)
		}
	}

	retunPaths := []string{}

	for _, f := range findFiles {
		slog.Debug("Check pkglist", "file", path.Join(profile, f))
		if osutils.IsFile(path.Join(profile, f)) {
			retunPaths = append(retunPaths, path.Join(profile, f))
		}
	}

	return retunPaths, nil
}
