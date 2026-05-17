package querybuildertypes

import (
	"fmt"
	"strings"

	"github.com/gear6io/pragmata/pkg/valuer"
)

type FieldDataType struct {
	valuer.String
}

var (
	FieldDataTypeString          = FieldDataType{valuer.NewString("string")}
	FieldDataTypeInt64           = FieldDataType{valuer.NewString("int64")}
	FieldDataTypeFloat64         = FieldDataType{valuer.NewString("float64")}
	FieldDataTypeBool            = FieldDataType{valuer.NewString("bool")}
	FieldDataTypeDateTime64      = FieldDataType{valuer.NewString("datetime64")}
	FieldDataTypeDate            = FieldDataType{valuer.NewString("date")}
	FieldDataTypeArrayString     = FieldDataType{valuer.NewString("[]string")}
	FieldDataTypeArrayInt64      = FieldDataType{valuer.NewString("[]int64")}
	FieldDataTypeArrayFloat64    = FieldDataType{valuer.NewString("[]float64")}
	FieldDataTypeArrayBool       = FieldDataType{valuer.NewString("[]bool")}
	FieldDataTypeArrayDateTime64 = FieldDataType{valuer.NewString("[]datetime64")}
	FieldDataTypeArrayDate       = FieldDataType{valuer.NewString("[]date")}

	FieldDataTypeDynamic      = FieldDataType{valuer.NewString("dynamic")}
	FieldDataTypeJSON         = FieldDataType{valuer.NewString("json")}
	FieldDataTypeArrayDynamic = FieldDataType{valuer.NewString("[]dynamic")}
	FieldDataTypeArrayJSON    = FieldDataType{valuer.NewString("[]json")}
)

// Field describes one column of a Source table.
type Field struct {
	Name string        `json:"name"`
	Type FieldDataType `json:"type"`
}

// fieldDataTypeToClickHouse maps each FieldDataType to its ClickHouse DDL type string.
var fieldDataTypeToClickHouse = map[FieldDataType]string{
	FieldDataTypeString:          "String",
	FieldDataTypeInt64:           "Int64",
	FieldDataTypeFloat64:         "Float64",
	FieldDataTypeBool:            "Bool",
	FieldDataTypeDateTime64:      "DateTime64(3)",
	FieldDataTypeDate:            "Date",
	FieldDataTypeArrayString:     "Array(String)",
	FieldDataTypeArrayInt64:      "Array(Int64)",
	FieldDataTypeArrayFloat64:    "Array(Float64)",
	FieldDataTypeArrayBool:       "Array(Bool)",
	FieldDataTypeArrayDateTime64: "Array(DateTime64(3))",
	FieldDataTypeArrayDate:       "Array(Date)",
	FieldDataTypeDynamic:         "Dynamic",
	FieldDataTypeJSON:            "JSON",
	FieldDataTypeArrayDynamic:    "Array(Dynamic)",
	FieldDataTypeArrayJSON:       "Array(JSON)",
}

// clickHouseToFieldDataType maps ClickHouse native type strings to FieldDataType.
// Types that share a mapping (e.g. Int32 → int64) are listed explicitly.
var clickHouseToFieldDataType = map[string]FieldDataType{
	"String":       FieldDataTypeString,
	"Int8":         FieldDataTypeInt64,
	"Int16":        FieldDataTypeInt64,
	"Int32":        FieldDataTypeInt64,
	"Int64":        FieldDataTypeInt64,
	"UInt8":        FieldDataTypeInt64,
	"UInt16":       FieldDataTypeInt64,
	"UInt32":       FieldDataTypeInt64,
	"UInt64":       FieldDataTypeInt64,
	"Float32":      FieldDataTypeFloat64,
	"Float64":      FieldDataTypeFloat64,
	"Bool":         FieldDataTypeBool,
	"DateTime64":   FieldDataTypeDateTime64,
	"Date":         FieldDataTypeDate,
	"Date32":       FieldDataTypeDate,
	"Array(String)": FieldDataTypeArrayString,
	"Array(Int64)":  FieldDataTypeArrayInt64,
	"Array(Float64)": FieldDataTypeArrayFloat64,
	"Array(Bool)":   FieldDataTypeArrayBool,
	"JSON":          FieldDataTypeJSON,
}

// ClickHouseType returns the ClickHouse DDL type string for this FieldDataType.
func (dt FieldDataType) ClickHouseType() (string, error) {
	if chType, ok := fieldDataTypeToClickHouse[dt]; ok {
		return chType, nil
	}
	return "", fmt.Errorf("unknown data type %q", dt.StringValue())
}

// FieldDataTypeFromClickHouse maps a ClickHouse native type string to a FieldDataType.
// Unknown types fall back to FieldDataTypeDynamic.
func FieldDataTypeFromClickHouse(chType string) FieldDataType {
	// Exact match first.
	if dt, ok := clickHouseToFieldDataType[chType]; ok {
		return dt
	}
	// Prefix matches for parameterised types (e.g. FixedString(64), DateTime64(6)).
	switch {
	case strings.HasPrefix(chType, "FixedString"):
		return FieldDataTypeString
	case strings.HasPrefix(chType, "DateTime64"):
		return FieldDataTypeDateTime64
	}
	return FieldDataTypeDynamic
}
