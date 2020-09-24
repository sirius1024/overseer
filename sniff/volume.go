package sniff

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sirius1024/overseer/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	// log.SetOutput()

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

// IO of volume to check health status
func IO() {
	conf := config.GetConfig()
	for _, volume := range conf.Overseer.Volumes {
		content, err := write(volume.Path)
		//TODO: 发送至nsq
		if err != nil {
			log.WithFields(log.Fields{
				"from": strings.Join([]string{conf.Cloud, conf.NetworkZone, conf.PrivateIP}, "-"),
				"to":   strings.Join([]string{volume.Type, volume.Path}, "-"),
				"type": "storage",
			}).Error(err.Error())
		} else {
			log.WithFields(log.Fields{
				"from": strings.Join([]string{conf.Cloud, conf.NetworkZone, conf.PrivateIP}, "-"),
				"to":   strings.Join([]string{volume.Type, volume.Path}, "-"),
				"type": "storage",
			}).Info(content)
		}
	}
}

func write(path string) (result string, err error) {
	if !Exists(path) {
		// create folder
		err = os.MkdirAll(path, os.ModePerm)
	}
	// write file by date

	fileName := (time.Now().Format("20060102")) + ".txt"
	var fileContent = []byte(time.Now().Format("2006-01-02 15:04:05") + "\n")
	logPath := filepath.Join(path, fileName)

	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return logPath, err
	}
	defer f.Close()
	_, err = f.Write(fileContent)

	// err = ioutil.WriteFile(logPath, fileContent, os.ModePerm)
	return logPath, err
}

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
