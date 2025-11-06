package main

import (
	"fmt"
	"reflect"
	"sort"
)

func OrdenarPorKey(lista []map[string]interface{}, key string, reverse bool) []map[string]interface{} {
	out := make([]map[string]interface{}, len(lista))
	copy(out, lista)

	// sort con lambda (función anónima) como comparador
	sort.SliceStable(out, func(i, j int) bool {
		ai, ahi := out[i][key]
		bj, bhj := out[j][key]

		// Quien no tiene la key va al final (en asc); al inicio si reverse=true
		if ahi != bhj {
			if reverse {
				return !ahi && bhj // faltantes primero en descendente
			}
			return ahi && !bhj // faltantes al final en ascendente
		}
		// Ambos faltan: mantener orden estable
		if !ahi && !bhj {
			return false
		}

		c := cmpAny(ai, bj) // -1 si ai<bj, 0 si ==, 1 si ai>bj
		if reverse {
			return c > 0
		}
		return c < 0
	})

	return out
}

// ---- utilidades de comparación ----
func asFloat(v interface{}) (float64, bool) {
	switch x := v.(type) {
	case int:
		return float64(x), true
	case int8:
		return float64(x), true
	case int16:
		return float64(x), true
	case int32:
		return float64(x), true
	case int64:
		return float64(x), true
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(v).Uint()), true
	case float32:
		return float64(x), true
	case float64:
		return x, true
	default:
		return 0, false
	}
}

func cmpAny(a, b interface{}) int {
	// números
	if fa, okA := asFloat(a); okA {
		if fb, okB := asFloat(b); okB {
			switch {
			case fa < fb:
				return -1
			case fa > fb:
				return 1
			default:
				return 0
			}
		}
	}
	// strings
	if sa, ok := a.(string); ok {
		if sb, ok2 := b.(string); ok2 {
			switch {
			case sa < sb:
				return -1
			case sa > sb:
				return 1
			default:
				return 0
			}
		}
	}
	// bool
	if ba, ok := a.(bool); ok {
		if bb, ok2 := b.(bool); ok2 {
			// false < true
			switch {
			case !ba && bb:
				return -1
			case ba && !bb:
				return 1
			default:
				return 0
			}
		}
	}
	// fallback: comparar por texto
	sa := fmt.Sprint(a)
	sb := fmt.Sprint(b)
	switch {
	case sa < sb:
		return -1
	case sa > sb:
		return 1
	default:
		return 0
	}
}

func main() {
	// Ejemplo del enunciado
	datos := []map[string]interface{}{
		{"make": "Nokia", "model": 216, "color": "Black"},
		{"make": "Mi Max", "model": 2, "color": "Gold"},
		{"make": "Samsung", "model": 7, "color": "Blue"},
	}

	res := OrdenarPorKey(datos, "model", false) // ascendente
	for _, m := range res {
		fmt.Println(m)
	}
}
