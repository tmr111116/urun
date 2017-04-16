package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

type projectVersion struct {
	Version string `yaml:"m_EditorVersion"`
}

func getProjVersion(proj string) string {
	file := filepath.Join(proj, "ProjectSettings", "ProjectVersion.txt")
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return "4"
	}
	pv := projectVersion{}
	yaml.Unmarshal(buf, &pv)

	return pv.Version
}

const projPathArg = "-projectPath"

func getProjPath(args []string) (path string, ok bool) {
	for i := 0; i < len(args)-1; i++ {
		if args[i] == projPathArg {
			return args[i+1], true
		}
	}

	return "", false
}

func getApplicationsPath() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("PROGRAMFILES")
	}

	return "/Applications"
}

func getExePathInUnity() string {
	if runtime.GOOS == "windows" {
		return `Editor\Unity.exe`
	}

	return "Unity.app/Contents/MacOS/Unity"
}

func getUnityExePath(version string) (path string, ok bool) {
	folders := []string{
		"Unity" + version,
		"Unity " + version,
	}
	for _, folder := range folders {
		path := filepath.Join(
			getApplicationsPath(),
			folder,
			getExePathInUnity(),
		)
		_, err := os.Stat(path)
		if err == nil {
			return path, true
		}
	}

	return "", false
}
