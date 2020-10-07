package config

import (
	"fmt"
	"io/ioutil"

	_ "github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

type Config struct {

	Application  *application
}


type application struct {
	Basepath string
	Port int
}

func NewConfig(cfile string) (*Config, error){
	buf, err := ioutil.ReadFile(cfile)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", cfile, err)
	}

	return c, nil
	//_, err := os.Stat(cfile)
	//if err != nil {
	//	// errMsg := fmt.Sprintf("Failed to read %s file statistics. Error: %v", cfile, err)
	//	// logger1.Error("Error", zap.String("mesg", errMsg))
	//	fmt.Printf("Failed to read %s file statistics. Error: %v", cfile, err)
	//
	//	return nil, errors.New("Failed to read %s file statistics. Error: %v\", cfile, err")
	//}
	//
	//
	//
	//var config Config
	//if _, err := toml.DecodeFile(cfile, &config); err != nil {
	//	// errMsg := fmt.Sprintf("Failed to decode %s file. Error: %v", cfile, err)
	//	// logger1.Error("Error", zap.String("mesg", errMsg))
	//	fmt.Printf("Failed to decode %s file. Error: %v", cfile, err)
	//
	//	return nil, errors.New("Failed to decode %s file. Error:")
	//}
	//// Validate config content
	//
	//return &config, nil
	////return &config, nil

}