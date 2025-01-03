package secondaryadministrativesubdivision

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// EnumSecondaryAdministrativeSubdivisionMap contains ALL records, including un-released and import-only records
	EnumSecondaryAdministrativeSubdivisionMap map[SecondaryAdministrativeSubdivision]*EnumSecondaryAdministrativeSubdivisionItem

	// EnumSecondaryAdministrativeSubdivisionListAll contains ALL records, including un-released and import-only records
	EnumSecondaryAdministrativeSubdivisionListAll []*EnumSecondaryAdministrativeSubdivisionItem

	// EnumSecondaryAdministrativeSubdivisionList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumSecondaryAdministrativeSubdivisionList []*EnumSecondaryAdministrativeSubdivisionItem
)

func init() {
	var err error
	EnumSecondaryAdministrativeSubdivisionListAll, err = spec.ParseADISpecTSV[*EnumSecondaryAdministrativeSubdivisionItem](spec.EnumSecondaryAdministrativeSubdivisionTSV)
	if err != nil {
		panic(err)
	}

	EnumSecondaryAdministrativeSubdivisionList = make([]*EnumSecondaryAdministrativeSubdivisionItem, 0, len(EnumSecondaryAdministrativeSubdivisionListAll))
	for _, item := range EnumSecondaryAdministrativeSubdivisionListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) {
			EnumSecondaryAdministrativeSubdivisionList = append(EnumSecondaryAdministrativeSubdivisionList, item)
		}
	}
	EnumSecondaryAdministrativeSubdivisionList = slices.Clip(EnumSecondaryAdministrativeSubdivisionList)

	EnumSecondaryAdministrativeSubdivisionMap = make(map[SecondaryAdministrativeSubdivision]*EnumSecondaryAdministrativeSubdivisionItem, len(EnumSecondaryAdministrativeSubdivisionList))
	for _, item := range EnumSecondaryAdministrativeSubdivisionList {
		EnumSecondaryAdministrativeSubdivisionMap[item.ID] = item
	}
}

// SecondaryAdministrativeSubdivision represents an antenna path abbreviation
type SecondaryAdministrativeSubdivision string

// EnumSecondaryAdministrativeSubdivisionItem represents an antenna path item
type EnumSecondaryAdministrativeSubdivisionItem struct {
	shared.ImportRecordRoot
	ID                                 SecondaryAdministrativeSubdivision `csv:"Code"` // The value that is stored in the ANT_PATH ADIF field.
	SecondaryAdministrativeSubdivision string                             `csv:"Secondary Administrative Subdivision"`
	DXCCEntityCode                     dxccentitycode.DXCCEntityCode      `csv:"DXCC Entity Code"`
	AlaskaJudicialDistrict             string                             `csv:"Alaska Judicial District"`
	Deleted                            shared.Deleted                     `csv:"Deleted"`
}
