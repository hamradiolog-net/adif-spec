package arrlsection

import (
	"slices"

	"github.com/hamradiolog-net/adif-spec/src/pkg/enum/dxccentitycode"
	"github.com/hamradiolog-net/adif-spec/src/pkg/shared"
	"github.com/hamradiolog-net/adif-spec/src/pkg/spec"
)

var (
	// EnumARRLSectionMap contains ALL records, including un-released and import-only records
	EnumARRLSectionMap map[ARRLSection]*EnumARRLSectionItem

	// EnumARRLSectionListAll contains ALL records, including un-released and import-only records
	EnumARRLSectionListAll []*EnumARRLSectionItem

	// EnumARRLSectionList
	// is a filtered list.
	// It excludes un-released and import-only records.
	EnumARRLSectionList []*EnumARRLSectionItem
)

func init() {
	var err error
	EnumARRLSectionListAll, err = spec.ParseADISpecTSV[*EnumARRLSectionItem](spec.EnumARRLSectionTSV)
	if err != nil {
		panic(err)
	}

	EnumARRLSectionList = make([]*EnumARRLSectionItem, 0, len(EnumARRLSectionListAll))
	for _, item := range EnumARRLSectionListAll {
		if bool(item.IsReleased) && !bool(item.IsImportOnly) && item.DeletedDate.Time.IsZero() {
			EnumARRLSectionList = append(EnumARRLSectionList, item)
		}
	}
	EnumARRLSectionList = slices.Clip(EnumARRLSectionList)

	EnumARRLSectionMap = make(map[ARRLSection]*EnumARRLSectionItem, len(EnumARRLSectionList))
	for _, item := range EnumARRLSectionList {
		EnumARRLSectionMap[item.ID] = item
	}
}

// ARRLSection represents an ARRL section identifier
type ARRLSection string

// EnumARRLSectionItem represents an ARRL section item
type EnumARRLSectionItem struct {
	shared.ImportRecordRoot
	ID              ARRLSection                       `csv:"Abbreviation"` // The value that is stored in the ARRL_SECTION and MY_ARRL_SECTION ADIF fields.
	Description     string                            `csv:"Section Name"`
	DXCCEntityCodes dxccentitycode.DXCCEntityCodeList `csv:"DXCC Entity Code"`
	FromDate        shared.ADIDate                    `csv:"From Date"`
	DeletedDate     shared.ADIDate                    `csv:"Deleted Date"`
}
