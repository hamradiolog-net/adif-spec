package morsekey

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/spec"
)

var (
	// EnumMorseKeyMap contains ALL records, including un-released and import-only records
	EnumMorseKeyMap EnumMorseKeyItemMap

	// EnumMorseKeyListAll contains ALL records, including un-released and import-only records
	EnumMorseKeyListAll EnumMorseKeyItemList

	// EnumMorseKeyList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumMorseKeyList EnumMorseKeyItemList
)

func init() {
	var err error
	EnumMorseKeyListAll, err = spec.ParseADISpecTSV[*EnumMorseKeyItem](spec.EnumMorseKeyTSV)
	if err != nil {
		panic(err)
	}

	EnumMorseKeyList = make([]*EnumMorseKeyItem, 0, len(EnumMorseKeyListAll))
	for _, item := range EnumMorseKeyListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			EnumMorseKeyList = append(EnumMorseKeyList, item)
		}
	}
	EnumMorseKeyList = slices.Clip(EnumMorseKeyList)

	EnumMorseKeyMap = make(EnumMorseKeyItemMap, len(EnumMorseKeyList))
	for _, item := range EnumMorseKeyList {
		EnumMorseKeyMap[item.ID] = item
	}
}

// MorseKey represents a morse key type identifier
type MorseKey string

// EnumMorseKeyItemList represents a collection of morse key definitions
type EnumMorseKeyItemList []*EnumMorseKeyItem

// EnumMorseKeyItemMap maps MorseKey to its definition
type EnumMorseKeyItemMap map[MorseKey]*EnumMorseKeyItem

// EnumMorseKeyItem represents a morse key item
type EnumMorseKeyItem struct {
	shared.ImportRecordRoot
	ID          MorseKey `csv:"Abbreviation"` // The value that is stored in the MORSE_KEY_TYPE and MY_MORSE_KEY_TYPE ADIF fields.
	Description string   `csv:"Meaning"`
}
