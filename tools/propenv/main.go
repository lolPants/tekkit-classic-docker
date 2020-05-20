package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	propertiesFile = "server.properties"
)

var (
	boolType   = reflect.TypeOf(true)
	stringType = reflect.TypeOf("")
	intType    = reflect.TypeOf(10)
)

// Variables represents all server variables
type Variables struct {
	AllowNether        bool   `env:"ALLOW_NETHER" prop:"allow-nether"`
	LevelName          string `env:"LEVEL_NAME" prop:"level-name"`
	EnableQuery        bool   `env:"ENABLE_QUERY" prop:"enable-query"`
	AllowFlight        bool   `env:"ALLOW_FLIGHT" prop:"allow-flight"`
	RconPassword       string `env:"RCON_PASSWORD" prop:"rcon.password"`
	ServerPort         int    `env:"SERVER_PORT" prop:"server-port"`
	LevelType          string `env:"LEVEL_TYPE" prop:"level-type"`
	EnableRcon         bool   `env:"ENABLE_RCON" prop:"enable-rcon"`
	LevelSeed          string `env:"LEVEL_SEED" prop:"level-seed"`
	ServerIP           string `env:"SERVER_IP" prop:"server-ip"`
	MaxBuildHeight     int    `env:"MAX_BUILD_HEIGHT" prop:"max-build-heigh"`
	SpawnNPCs          bool   `env:"SPAWN_NPCS" prop:"spawn-npcs"`
	Debug              bool   `env:"DEBUG" prop:"debug"`
	WhiteList          bool   `env:"WHITE_LIST" prop:"white-list"`
	SpawnAnimals       bool   `env:"SPAWN_ANIMALS" prop:"spawn-animals"`
	OnlineMode         bool   `env:"ONLINE_MODE" prop:"online-mode"`
	PvP                bool   `env:"PVP" prop:"pvp"`
	Difficulty         int    `env:"DIFFICULTY" prop:"difficulty"`
	Gamemode           int    `env:"GAMEMODE" prop:"gamemode"`
	MaxPlayers         int    `env:"MAX_PLAYERS" prop:"max-players"`
	RconPort           int    `env:"RCON_PORT" prop:"rcon.port"`
	SpawnMonsters      bool   `env:"SPAWN_MONSTERS" prop:"spawn-monsters"`
	GenerateStructures bool   `env:"GENERATE_STRUCTURES" prop:"generate-structures"`
	ViewDistance       int    `env:"VIEW_DISTANCE" prop:"view-distance"`
	Motd               string `env:"MOTD" prop:"motd"`
}

func main() {
	props, err := readProperties()
	if err != nil {
		return
	}

	vars := Variables{}
	v := reflect.ValueOf(vars)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag

		envVar := tag.Get("env")
		propTag := tag.Get("prop")

		value, ok := os.LookupEnv(envVar)
		if ok == false {
			continue
		}

		var resolved interface{}
		var err error

		switch field.Type {
		case boolType:
			resolved, err = handleBool(value)
			break

		case stringType:
			resolved, err = handleString(value)
			break

		case intType:
			resolved, err = handleInt(value)
			break
		}

		if err != nil {
			continue
		}

		parsed := fmt.Sprintf("%v", resolved)
		found := false

		for _, p := range *props {
			if p[0] == propTag {
				p[1] = parsed
				found = true
			}
		}

		if found == false {
			pair := []string{propTag, parsed}

			slice := *props
			p := append(slice, pair)
			props = &p
		}
	}

	writeProperties(props)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func readProperties() (*[][]string, error) {
	if fileExists(propertiesFile) == false {
		empty := make([][]string, 0)
		return &empty, nil
	}

	bytes, err := ioutil.ReadFile(propertiesFile)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	fields := make([][]string, len(lines))

	for i, line := range lines {
		fields[i] = strings.Split(line, "=")
	}

	return &fields, nil
}

func writeProperties(fields *[][]string) {
	lines := make([]string, len(*fields))
	for i, pair := range *fields {
		lines[i] = strings.Join(pair, "=")
	}

	file := strings.TrimSpace(strings.Join(lines, "\n"))
	file += "\n"

	bytes := []byte(file)
	ioutil.WriteFile(propertiesFile, bytes, 0755)
}

func handleBool(value string) (bool, error) {
	lower := strings.ToLower(value)

	if lower == "true" || lower == "t" || lower == "1" || lower == "yes" || lower == "y" {
		return true, nil
	}

	if lower == "false" || lower == "f" || lower == "0" || lower == "no" || lower == "n" {
		return false, nil
	}

	err := errors.New("Invalid boolean")
	return false, err
}

func handleString(value string) (string, error) {
	return value, nil
}

func handleInt(value string) (int, error) {
	return strconv.Atoi(value)
}
