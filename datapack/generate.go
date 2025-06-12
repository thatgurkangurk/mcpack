package datapack

import (
	"path/filepath"

	"github.com/thatgurkangurk/mcpack/utils"
)

func CreatePack(packName, packId, outDir string, forceClean bool) error {
	path := filepath.Join(outDir, "datapack")

	err := utils.CheckOrCreateFolder(path)

	if err != nil {
		return err
	}

	err = CreateMcMeta(McMeta{
		Pack: Pack{
			PackFormat:  71,
			Description: packName,
		},
	}, path)

	if err != nil {
		return err
	}

	err = utils.CheckOrCreateFolder(filepath.Join(outDir, "datapack", "data", packId))

	if err != nil {
		return err
	}

	return nil
}
