package main

import "time"

const(
	FilterType_EQUAL="EQUAL"
	FilterType_GREAT_EQUAL_THAN="GREAT_EQUAL_THAN"
	FilterType_DATE_LESS_EQUAL_THAN="DATE_LESS_EQUAL_THAN"
)
type FilterSet struct {
	Name string
	Type string
	ParamType string
	Format string
	Value interface{}
}
func NewDefaultFilterSet(name string,tp string,val interface{})*FilterSet{
	switch val.(type) {
	case time.Time:
		return &FilterSet{
			Name: name,
			Type: tp,
			ParamType: "date",
			Format: "yyyy-MM-dd",
			Value: val.(time.Time).Format("2006-01-02"),
		}
	case *time.Time:
		return &FilterSet{
			Name: name,
			Type: tp,
			ParamType: "date",
			Format: "yyyy-MM-dd",
			Value: val.(*time.Time).Format("2006-01-02"),
		}
	default:
		return &FilterSet{
			Name: name,
			Type: tp,
			Value: val,
		}
	}
}

type FilterSets struct {
	Sets []*FilterSet
}

func (f *FilterSets) Add(set *FilterSet)  {
	f.Sets=append(f.Sets,set)
}
func (f *FilterSets) Size()int{
	return len(f.Sets)
}
func (f *FilterSets) Maps()[]map[string]interface{}{
	rs:=make([]map[string]interface{},0)
	for _, it := range f.Sets {
		switch it.ParamType {
		case "date":
			rs=append(rs, map[string]interface{}{
				"name":it.Name,
				"type":it.Type,
				"paramType":it.ParamType,
				"format":it.Format,
				"value":it.Value,
			})
		default:
			rs=append(rs, map[string]interface{}{
				"name":it.Name,
				"type":it.Type,
				"value":it.Value,
			})
		}
	}
	return rs
}

func NewFilterSets()*FilterSets{
	return &FilterSets{}
}