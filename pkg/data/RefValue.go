package data

type RefValue struct {
	Symbolic bool
	Value    string
}

func InitRefValue(value string) RefValue {
	return RefValue{
		Value: value,
	}
}
