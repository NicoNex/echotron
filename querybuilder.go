/*
 * Echotron
 * Copyright (C) 2018-2022 The Echotron Devs
 *
 * Echotron is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Echotron is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package echotron

import (
	"encoding/json"
	"net/url"
	"reflect"
	"strconv"
)

func toString(v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return v.String()

	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)

	case reflect.Int, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)

	case reflect.Bool:
		return strconv.FormatBool(v.Bool())

	case reflect.Struct, reflect.Interface, reflect.Slice, reflect.Array:
		b, _ := json.Marshal(v.Interface())
		return string(b)

	default:
		return ""
	}
}

func scan(i any, v url.Values) url.Values {
	e := reflect.ValueOf(i)

	if e.Kind() == reflect.Pointer {
		e = e.Elem()
	}

	if e.Kind() == reflect.Invalid {
		return v
	}

	for i := 0; i < e.NumField(); i++ {
		fTag := e.Type().Field(i).Tag

		if name := fTag.Get("query"); name != "" && !e.Field(i).IsZero() {
			v.Set(name, toString(e.Field(i)))
		}
	}

	return v
}

func urlValues(i any) url.Values {
	if i == nil {
		return nil
	}
	return scan(i, url.Values{})
}

func addValues(vals url.Values, i any) url.Values {
	if i == nil {
		return vals
	}
	if vals == nil {
		vals = make(url.Values)
	}
	return scan(i, vals)
}
