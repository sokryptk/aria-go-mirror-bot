package env

import "C"
import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var Config Configuration

func init() {
	//Reads from json and gives back the required values
	Config.ReadFromFile()
}

//Model file to unmarshal the json file

type Configuration struct {
	BotToken	string `json:"botToken"`
	FolderId	string `json:"folderID"`
	AriaArgs	[]string `json:"ariaArgs"`
	RpcSecret	string	`json:"rpcSecret"`
	Commands	CommandsConfig	`json:"commands"`
}

type CommandsConfig struct {
	MirrorCommands	string `json:"mirrorCommand"`
	MirrorTarCommands	string `json:"mirrorTarCommand"`
	ListCommands	string `json:"listCommand"`
	StatusCommands	string	`json:"statusCommand"`
}

func (c *Configuration) ReadFromFile() {
	file, err := os.OpenFile("config.json", os.O_RDONLY, os.ModePerm)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	fileBody, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	_ = json.Unmarshal(fileBody, &c)
}