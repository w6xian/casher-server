package config

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// f可以传目录也可以传文件
func MustLoad(f string, val any, configType string) error {
	readAllConfig(f, configType)
	err := viper.Unmarshal(val)
	if err != nil {
		panic(err)
	}
	return nil
}

func readAllConfig(fPath string, configType string) {
	f, err := os.Stat(fPath)
	if err != nil {
		panic(err)
	}
	viper.SetConfigType(configType)
	if f.IsDir() {
		viper.AddConfigPath(fPath)
		filepath.Walk(fPath, func(f string, info os.FileInfo, err error) error {
			if err != nil {
				panic(err)
			}
			if !info.IsDir() {
				ff := strings.Split(f, string(os.PathSeparator))
				filename := ff[len(ff)-1]
				filesuffix := path.Ext(filename)
				if filesuffix[1:] == configType {
					readConfigFile(filename)
				}
			}
			return nil
		})
	} else {
		filename := fPath
		filenameall := path.Base(filename)
		dir := getCurrentAbPath(filename)
		filesuffix := path.Ext(filename)
		if filesuffix[1:] == configType {
			viper.AddConfigPath(dir)
			readConfigFile(filenameall)
		}

	}
}

func readConfigFile(filename string) {
	ff := strings.Split(filename, ".")
	ff = ff[0 : len(ff)-1]
	filename = strings.Join(ff, ".")
	viper.SetConfigName(filename)
	err := viper.MergeInConfig()
	if err != nil {
		panic(err)
	}
}

func getCurrentAbPath(f string) string {
	var aPath []string
	if strings.Contains(f, "\\") {
		aPath = strings.Split(f, "\\")
	} else {
		aPath = strings.Split(f, "/")
	}
	dir := strings.Join(aPath[0:len(aPath)-1], string(os.PathSeparator))
	if dir == "" {
		dir = "./"
	}
	return dir
}
