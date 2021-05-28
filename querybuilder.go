/*
 * Echotron
 * Copyright (C) 2018-2021  Nicolò Santamaria, Michele Dimaggio
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
		return strconv.FormatFloat(v.Interface().(float64), 'f', -1, 64)

	case reflect.Int:
		i := v.Interface().(int)
		return strconv.FormatInt(int64(i), 10)

	case reflect.Int64:
		return strconv.FormatInt(v.Interface().(int64), 10)

	case reflect.Bool:
		return strconv.FormatBool(v.Interface().(bool))

	case reflect.Struct, reflect.Interface, reflect.Slice, reflect.Array:
		b, _ := json.Marshal(v.Interface())
		return string(b)

	default:
		return ""
	}
}

func scan(i interface{}, v url.Values) url.Values {
	e := reflect.ValueOf(i).Elem()

	if e.Kind() == reflect.Invalid {
		return url.Values{}
	}

	for i := 0; i < e.NumField(); i++ {
		fTag := e.Type().Field(i).Tag

		switch name := fTag.Get("query"); name {
		case "recursive":
			scan(e.Field(i).Interface(), v)

		case "":
			continue

		default:
			if !e.Field(i).IsZero() {
				v.Set(name, toString(e.Field(i)))
			}
		}
	}
	return v
}

func querify(i interface{}) string {
	return scan(i, url.Values{}).Encode()
}
