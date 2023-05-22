package viperp

import "github.com/spf13/afero"

type ViperP struct {
	configFile string

	// The filesystem to read config from.
	fs afero.Fs

	logger Logger
}

func New() *ViperP {
	v := new(ViperP)

	// If don't set value, it may have nil point err when somewhere code use it.
	v.fs = afero.NewOsFs()
	return v
}

func (v *ViperP) SetFs(fs afero.Afero) {
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
}
