package viperp

import (
	"bytes"
	"fmt"
	"github.com/new-aspect/viperp/internal/encoding"
	"github.com/spf13/afero"
	"io"
	"path/filepath"
	"strings"
)

type UnsupportedConfigError string

func (str UnsupportedConfigError) Error() string {
	return fmt.Sprintf("Unsupported Config Type %q", string(str))
}

var SupportedExts = []string{"json", "toml", "yaml", "yml", "properties"}

type ViperP struct {
	configFile string
	configType string

	config map[string]interface{}

	// The filesystem to read config from.
	fs afero.Fs

	logger Logger

	decoderRegistry *encoding.DecoderRegistry
}

func New() *ViperP {
	v := new(ViperP)

	// If don't set value, it may have nil point err when somewhere code use it.
	v.fs = afero.NewOsFs()
	return v
}

func (v *ViperP) SetFs(fs afero.Fs) {
	// 值只能内部访问，然后通过方法赋值的方式叫封装，目的是保护对象状态，防止外部代码随意修改属性状态
	v.fs = fs
}

func (v *ViperP) SetConfigFile(in string) {
	if in != "" {
		v.configFile = in
	}
}

func (v *ViperP) ReadInConfig() error {
	v.logger.Info("attemption to read in config file")
	filename, err := v.getConfigFile()
	if err != nil {
		return err
	}

	if !stringInSlice(v.getConfigType(), SupportedExts) {
		return UnsupportedConfigError(v.getConfigType())
	}

	v.logger.Debug("reading file", "file", filename)
	file, err := afero.ReadFile(v.fs, filename)
	if err != nil {
		return err
	}

	config := make(map[string]interface{})

	err = v.unmarshalReader(bytes.NewReader(file), config)

	v.config = config
	return nil
}

func (v *ViperP) getConfigFile() (string, error) {
	if v.configFile == "" {
		// todo 这部分逻辑不实践
		return "", fmt.Errorf("没有找到 configFile 文件")
	}
	return v.configFile, nil
}

func (v *ViperP) getConfigType() string {
	if v.configType != "" {
		return v.configType
	}

	cf, err := v.getConfigFile()
	if err != nil {
		return ""
	}

	// 还可以用Ext直接返回扩展名，注意，这个返回包含"."号，例如".json",".yaml"等
	ext := filepath.Ext(cf)

	// 如果取得了拓展名，也就是 len(ext) > 1 , 那就返回拓展名不带"."号，这是通过ext[1:]实现的，
	// 例如ext是".json", 那么ext[1:]就是"json"
	if len(ext) > 1 {
		return ext[1:]
	}

	return ""
}

func (v *ViperP) unmarshalReader(in io.Reader, c map[string]interface{}) error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(in)

	switch format := strings.ToLower(v.getConfigType()); format {
	case "yaml", "yml", "json", "toml", "hcl", "tfvars", "ini", "properties", "props", "prop", "dotenv", "env":
		err := v.decoderRegistry.Decode(format, buf.Bytes(), c)
		if err != nil {
			return ConfigParseError{err}
		}
	}

	insensitiviseMap(c)
	return nil
}
