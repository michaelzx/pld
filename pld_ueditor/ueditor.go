package pld_ueditor

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func init() {
	//syscall.Umask(0)
}

// 参考：https://github.com/dazhenghu/gueditor
type UEditor struct {
	config    *Config
	webRoot   string
	uploadDir string
}

func NewUEditor(webRoot, uploadDir string) (*UEditor, error) {
	file, err := os.Open("./resource/ueditor_config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fd, err := ioutil.ReadAll(file)
	cfgJson := string(fd)
	re, _ := regexp.Compile("\\/\\*[\\S\\s]+?\\*\\/")
	cfgJson = re.ReplaceAllString(cfgJson, "")
	cfgJson = strings.Replace(cfgJson, "#{UploadDir}", uploadDir, -1)
	var cfg Config
	err = json.Unmarshal([]byte(cfgJson), &cfg)
	if err != nil {
		return nil, err
	}
	return &UEditor{
		config:    &cfg,
		webRoot:   webRoot,
		uploadDir: uploadDir,
	}, nil
}

func (ue *UEditor) GetConfig() *Config {
	return ue.config
}
