package common

import (
	"errors"
	"strconv"
)

func String(args ...interface{}) string {
	value := args[0]
	var precision int = 12

	switch value.(type) {
	case string:
		v, _ := value.(string)
		return v
	case int:
		v, _ := value.(int)
		return strconv.Itoa(v)
	case int32:
		v, _ := value.(int32)
		return strconv.FormatInt(int64(v), 10)
	case int64:
		v, _ := value.(int64)
		return strconv.FormatInt(v, 10)
	case uint32:
		v, _ := value.(uint32)
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		v, _ := value.(uint64)
		return strconv.FormatUint(v, 10)
	case float32:
		v, _ := value.(float32)
		if len(args) >= 2 {
			precision = args[1].(int)
		}
		return strconv.FormatFloat(float64(v), 'f', precision, 64)
	case float64:
		v, _ := value.(float64)
		if len(args) >= 2 {
			precision = args[1].(int)
		}
		return strconv.FormatFloat(v, 'f', precision, 64)
	default:
		return ""
	}
}

func Int(value interface{}) (int, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		return strconv.Atoi(v)
	case int:
		v, _ := value.(int)
		return v, nil
	case int32:
		v, _ := value.(int32)
		return int(v), nil
	case int64:
		v, _ := value.(int64)
		return int(v), nil
	case uint32:
		v, _ := value.(uint32)
		return int(v), nil
	case uint64:
		v, _ := value.(uint64)
		return int(v), nil
	case float32:
		v, _ := value.(float32)
		return int(v), nil
	case float64:
		v, _ := value.(float64)
		return int(v), nil
	default:
		return int(0), errors.New("unknown type")
	}
}

func Int32(value interface{}) (int32, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		result, err := strconv.ParseInt(v, 10, 32)
		return int32(result), err
	case int:
		v, _ := value.(int)
		return int32(v), nil
	case int32:
		v, _ := value.(int32)
		return int32(v), nil
	case int64:
		v, _ := value.(int64)
		return int32(v), nil
	case uint32:
		v, _ := value.(uint32)
		return int32(v), nil
	case uint64:
		v, _ := value.(uint64)
		return int32(v), nil
	case float32:
		v, _ := value.(float32)
		return int32(v), nil
	case float64:
		v, _ := value.(float64)
		return int32(v), nil
	default:
		return int32(0), errors.New("unknown type")
	}
}

func Uint32(value interface{}) (uint32, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		result, err := strconv.ParseUint(v, 10, 32)
		return uint32(result), err
	case int:
		v, _ := value.(int)
		return uint32(v), nil
	case int32:
		v, _ := value.(int32)
		return uint32(v), nil
	case int64:
		v, _ := value.(int64)
		return uint32(v), nil
	case uint32:
		v, _ := value.(uint32)
		return v, nil
	case uint64:
		v, _ := value.(uint64)
		return uint32(v), nil
	case float32:
		v, _ := value.(float32)
		return uint32(v), nil
	case float64:
		v, _ := value.(float64)
		return uint32(v), nil
	default:
		return uint32(0), errors.New("unknown type")
	}
}

func Uint64(value interface{}) (uint64, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		result, err := strconv.ParseUint(v, 10, 64)
		return result, err
	case int:
		v, _ := value.(int)
		return uint64(v), nil
	case int32:
		v, _ := value.(int32)
		return uint64(v), nil
	case int64:
		v, _ := value.(int64)
		return uint64(v), nil
	case uint32:
		v, _ := value.(uint32)
		return uint64(v), nil
	case uint64:
		v, _ := value.(uint64)
		return v, nil
	case float32:
		v, _ := value.(float32)
		return uint64(v), nil
	case float64:
		v, _ := value.(float64)
		return uint64(v), nil
	default:
		return uint64(0), errors.New("unknown type")
	}
}

func Int64(value interface{}) (int64, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		return strconv.ParseInt(v, 10, 64)
	case int:
		v, _ := value.(int)
		return int64(v), nil
	case int32:
		v, _ := value.(int32)
		return int64(v), nil
	case int64:
		v, _ := value.(int64)
		return v, nil
	case uint32:
		v, _ := value.(uint32)
		return int64(v), nil
	case uint64:
		v, _ := value.(uint64)
		return int64(v), nil
	case float32:
		v, _ := value.(float32)
		return int64(v), nil
	case float64:
		v, _ := value.(float64)
		return int64(v), nil
	default:
		return int64(0), errors.New("unknown type")
	}
}

func Float32(value interface{}) (float32, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		result, err := strconv.ParseFloat(v, 32)
		return float32(result), err
	case int:
		v, _ := value.(int)
		return float32(v), nil
	case int32:
		v, _ := value.(int32)
		return float32(v), nil
	case int64:
		v, _ := value.(int64)
		return float32(v), nil
	case uint32:
		v, _ := value.(uint32)
		return float32(v), nil
	case uint64:
		v, _ := value.(uint64)
		return float32(v), nil
	case float32:
		v, _ := value.(float32)
		return v, nil
	case float64:
		v, _ := value.(float64)
		return float32(v), nil
	default:
		return float32(0), errors.New("unknown type")
	}
}

func Float64(value interface{}) (float64, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		return strconv.ParseFloat(v, 64)
	case int:
		v, _ := value.(int)
		return float64(v), nil
	case int32:
		v, _ := value.(int32)
		return float64(v), nil
	case int64:
		v, _ := value.(int64)
		return float64(v), nil
	case uint32:
		v, _ := value.(uint32)
		return float64(v), nil
	case uint64:
		v, _ := value.(uint64)
		return float64(v), nil
	case float32:
		v, _ := value.(float32)
		return float64(v), nil
	case float64:
		v, _ := value.(float64)
		return v, nil
	default:
		return float64(0), errors.New("unknown type")
	}
}
