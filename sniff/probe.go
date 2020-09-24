package sniff

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/sirius1024/overseer/config"
	"github.com/sirius1024/overseer/encrypt"
	"github.com/sirius1024/overseer/models"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

// Pong 应答web ping
func Pong(context *gin.Context) {
	conf := config.GetConfig()
	encrypted, err := context.GetRawData()
	if err != nil {
		log.Fatalf(err.Error())
	}
	de, err := encrypt.Decrypt(encrypted, []byte(conf.Key))
	if err != nil {
		log.Fatalf(err.Error())
	}
	var ping models.Ping
	err = json.Unmarshal(de, &ping)
	if err != nil {
		// fake request, reject
		log.Fatalf(err.Error())
	}
	context.JSON(200, ping)
}

// Probe 探测所有endpoint
func Probe() {
	conf := config.GetConfig()
	for _, probe := range conf.Overseer.Probes {
		if !strings.HasSuffix(probe.Endpoint, "/") {
			probe.Endpoint = probe.Endpoint + "/"
		}
		probe.Endpoint += "ping"
		resp, err := ping(probe.Endpoint)

		// TODO: 发送至nsq
		if err != nil {
			// Log输出
			log.WithFields(log.Fields{
				"from": strings.Join([]string{conf.Cloud, conf.NetworkZone, conf.PrivateIP}, "-"),
				"to":   strings.Join([]string{probe.EndpointName, probe.Endpoint}, "-"),
				"type": "network",
			}).Error(err.Error())
		} else {
			log.WithFields(log.Fields{
				"from": strings.Join([]string{conf.Cloud, conf.NetworkZone, conf.PrivateIP}, "-"),
				"to":   strings.Join([]string{probe.EndpointName, probe.Endpoint}, "-"),
				"type": "network",
			}).Info(resp)
		}
	}
}

// ping 发起web ping，用于检测网络健康
func ping(endpoint string) (result string, err error) {
	// payload is who am I and where I am.
	// to which IP and port.
	// payload must be encrypted.
	// response is payload and server-self informations.
	// callback to nsq to record network checking.
	// nsq installed locally.

	conf := config.GetConfig()
	payload := conf.ToPing()
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf(err.Error())
	}
	encrypted, err := encrypt.Encrypt(payloadBytes, []byte(conf.Key))

	request, _ := http.NewRequest("POST", endpoint, bytes.NewReader(encrypted))
	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		// 请求失败
		// fmt.Printf("post data error:%v\n", err)
		return "", err
	}
	// 请求200
	respBody, _ := ioutil.ReadAll(resp.Body)
	// 反序列化成string先
	// fmt.Printf("response data:%v\n", string(respBody))
	return string(respBody), err

}
