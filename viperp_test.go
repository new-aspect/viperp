package viperp

import (
	"github.com/new-aspect/viperp/internal/testutil"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadInConfig(t *testing.T) {
	t.Run("config file set", func(t *testing.T) {
		fs := afero.NewMemMapFs()

		err := fs.Mkdir("/etc/viper", 0o777)
		require.NoError(t, err)

		file, err := fs.Create(testutil.AbsFilePath(t, "/etc/viper/config.yaml"))
		require.NoError(t, err)

		_, err = file.Write([]byte(`key:value`))
		require.NoError(t, err)

		file.Close()

		v := New()

		v.SetFs(fs)
		v.SetConfigFile("/etc/viper/config.yaml")

		v.ReadInConfig()
	})
}
