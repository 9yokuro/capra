package main

import (
	"encoding/json"
	"errors"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

func jsonToMap(data []byte) (map[string]any, error) {
	var m map[string]any

	if err := json.Unmarshal(data, &m); err != nil {
		return make(map[string]any), err
	}

	return m, nil
}

func mapToJSON(m *map[string]any) ([]byte, error) {
	data, err := json.MarshalIndent(m, "", "  ")

	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

func tomlToMap(data []byte) (map[string]any, error) {
	var m map[string]any

	if err := toml.Unmarshal(data, &m); err != nil {
		return make(map[string]any), err
	}

	return m, nil
}

func mapToTOML(m *map[string]any) ([]byte, error) {
	data, err := toml.Marshal(m)

	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

func yamlToMap(data []byte) (map[string]any, error) {
	var m map[string]any

	if err := yaml.Unmarshal(data, &m); err != nil {
		return make(map[string]any), err
	}

	return m, nil
}

func mapToYAML(m *map[string]any) ([]byte, error) {
	data, err := yaml.Marshal(m)

	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

func ToMap(content []byte) (map[string]any, error) {
	if m, err := jsonToMap(content); err == nil {
		return m, err
	}

	if m, err := yamlToMap(content); err == nil {
		return m, err
	}

	if m, err := tomlToMap(content); err == nil {
		return m, err
	}

	return make(map[string]any), errors.New("Invalid input")
}

func ToMaps(inputs []string) []map[string]any {
	var data []map[string]any

	if len(inputs) == 0 {
		content, err := readPipe()

		if err != nil {
			Eprintln(err)
		}

		if d, err := ToMap(content); err != nil {
			Eprintln(err)
		} else {
			data = append(data, d)
		}
	} else {
		for _, input := range inputs {
			content, err := readFile(input)

			if err != nil {
				Eprintln(err)
			}

			if d, err := ToMap(content); err != nil {
				Eprintln(err)
			} else {
				data = append(data, d)
			}
		}
	}

	return data
}

func ToBytes(maps []map[string]any, format string) ([][]byte, error) {
	var bytes [][]byte

	switch format {
	case "json":
		for _, m := range maps {
			if json, err := mapToJSON(&m); err != nil {
				Eprintln(err)
			} else {
				bytes = append(bytes, json)
			}
		}
	case "yml", "yaml":
		for _, m := range maps {
			if yaml, err := mapToYAML(&m); err != nil {
				Eprintln(err)
			} else {
				bytes = append(bytes, yaml)
			}
		}
	case "toml":
		for _, m := range maps {
			if toml, err := mapToTOML(&m); err != nil {
				Eprintln(err)
			} else {
				bytes = append(bytes, toml)
			}
		}
	default:
		return [][]byte{}, errors.New("Invalid file format")
	}

	return bytes, nil
}
