package datapack

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Pack struct {
	PackFormat  int    `json:"pack_format"`
	Description string `json:"description"`
}

type McMeta struct {
	Pack Pack `json:"pack"`
}

func CreateMcMeta(data McMeta, outPath string) error {
	packMetaPath := filepath.Join(outPath, "pack.mcmeta")

	_, err := os.Stat(packMetaPath)
	if err == nil {
		return nil
	}

	err = os.MkdirAll(filepath.Dir(packMetaPath), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directories for pack.mcmeta: %v", err)
	}

	file, err := os.Create(packMetaPath)
	if err != nil {
		return fmt.Errorf("failed to create pack.mcmeta file: %v", err)
	}
	defer file.Close()

	metaBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal pack.mcmeta to JSON: %v", err)
	}

	_, err = file.Write(metaBytes)
	if err != nil {
		return fmt.Errorf("failed to write to pack.mcmeta: %v", err)
	}

	return nil
}
