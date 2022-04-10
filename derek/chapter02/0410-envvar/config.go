package envvar

import (
	"encoding/json"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// path에 있는 json 파일을 가져옴.
func LoadConfig(path, envPrefix string, config interface{}) error {
	// 경로에 있는 json 파일 가져옴.
	if path != "" {
		err := LoadFile(path, config)
		if err != nil {
			return errors.Wrap(err, "error loading config from file")
		}
	}
	// envconfig 구조체에 envPrefix 환경변수 접두어를 설정해서 값을 덮어 씌움.
	err := envconfig.Process(envPrefix, config)
	return errors.Wrap(err, "error loading config from env")
}

// LoadFile json 파일을 interface 구조체로 변환.
func LoadFile(path string, config interface{}) error {
	// 경로 파일 오픈.
	configFile, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "failed to read config file")
	}
	// 나중에 실행될놈.
	defer configFile.Close()
	// configFile 을  buffer로 디코드 해서 리턴.
	decoder := json.NewDecoder(configFile)
	// decoder을 json으로 전환.
	if err = decoder.Decode(config); err != nil {
		return errors.Wrap(err, "failed to decode config file")
	}
	return nil
}
