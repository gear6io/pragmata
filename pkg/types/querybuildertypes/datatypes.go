package querybuildertypes

import "github.com/gear6io/pragmata/pkg/valuer"

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
