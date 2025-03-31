package types

import (
	"strings"

	"github.com/tphan267/common/utils"
)

type Params map[string]any

// type Params interface {
// 	Get(key string) any
// 	Set(key string, val any)
// }

func (p *Params) Set(key string, val any) {
	params := p
	keys := strings.Split(key, ".")
	l := len(keys) - 1
	for i, k := range keys {
		if existParams, ok := (*params)[k]; !ok {
			// key not exists
			if i == l {
				(*params)[k] = val
			} else {
				newParams := Params{}
				(*params)[k] = newParams
				params = &newParams
			}
		} else {
			// key exists
			if i == l {
				(*params)[k] = val
			} else {
				exist := Params(existParams.(map[string]any))
				params = &exist
			}
		}
	}
}

func (p *Params) Get(key string) any {
	var val any
	var ok bool
	params := *p
	keys := strings.Split(key, ".")
	l := len(keys) - 1
	for i, k := range keys {
		if val, ok = params[k]; ok {
			if i == l {
				return val
			} else {
				params = Params(val.(map[string]any))
			}
		}
	}
	return nil
}

func (p *Params) GetParams(key string, defaultVals ...*Params) *Params {
	if val := p.Get(key); val != nil {
		params := Params(val.(map[string]any))
		return &params
	}
	if len(defaultVals) > 0 {
		return defaultVals[0]
	}
	return nil
}

func (p *Params) GetString(key string, defaultVals ...string) string {
	if val := p.Get(key); val != nil {
		return val.(string)
	}
	if len(defaultVals) > 0 {
		return defaultVals[0]
	}
	return ""
}

func (p *Params) GetBool(key string, defaultVals ...bool) bool {
	if val := p.Get(key); val != nil {
		return val.(bool)
	}
	if len(defaultVals) > 0 {
		return defaultVals[0]
	}
	return false
}

func (p *Params) GetInt(key string, defaultVals ...int) int {
	if val := p.Get(key); val != nil {
		return int(utils.ToInt(val))
	}
	if len(defaultVals) > 0 {
		return defaultVals[0]
	}
	return 0
}

func (p *Params) GetUint(key string) uint {
	if val := p.Get(key); val != nil {
		return uint(utils.ToInt(val))
	}
	return 0
}

func (p *Params) GetUint64(key string) uint64 {
	if val := p.Get(key); val != nil {
		return uint64(utils.ToInt(val))
	}
	return 0
}
