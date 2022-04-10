package envvar

import (
	"encoding/json"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// LoadConfig 함수는 path에 저장된 json 파일로부터 파일을 선택적으로 읽은 다음,
// envconfig 구조체를 기반으로 읽어온 값을 덮어 쓴다
// envPrefix는 환경 변수에 접두어를 설정하는 방법이다.
func LoadConfig(path, envPrefix string, config interface{}) error {
	if path != "" {
		err := LoadFile(path, config)
		if err != nil {
			// Wrap은 스택 추적으로 err에 주석을 추가하는 오류를 반환한다.
			return errors.Wrap(err, "error loading config from file")
		}
	}
	err := envconfig.Process(envPrefix, config)
	return errors.Wrap(err, "error loading config from file")
}

// LoadFile 함수는 json 파일을 구성 구조체로 변환한다
func LoadFile(path string, config interface{}) error {
	configFile, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "failed to read config file")
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	if err = decoder.Decode(config); err != nil {
		return errors.Wrap(err, "failed to decode config file")
	}
	return nil
}
