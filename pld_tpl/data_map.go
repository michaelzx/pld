package pld_tpl

type DataMap map[string]interface{}

func NewDataMap() DataMap {
	return make(DataMap)
}
