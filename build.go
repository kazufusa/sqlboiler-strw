package main

var stringConverters map[string]string

func init() {
	stringConverters := map[string]string{}
	stringConverters["int"] = `
	var _v int64
	_v, err = strconv.ParseInt(value, 10, 32)
	if err != nil {
		return nil, err
	}
	v := int(_v)
	{.Return}
	`

	stringConverters["int32"] = `
	var _v int64
	_v, err = strconv.ParseInt(value, 10, 32)
	if err != nil {
		return nil, err
	}
	v := int32(_v)
	{.Return}
	`

	stringConverters["int64"] = `
	var _v int64
	v, err = strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil, err
	}
	{.Return}
	`

	stringConverters["float64"] = `
	v, err = strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil, err
	}
	{.Return}
	`

	stringConverters["float32"] = `
	var _v float64
	var v float32
	_v, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return nil, err
	}
	v = float32(_v)
	{.Return}
	`

}

// return {.Name}Where.{.WhereName}.{.OP}(v), nil
