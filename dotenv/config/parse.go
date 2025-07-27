package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// ParseFiles 解析多个配置问题，如果存在相同的key，
// 后面会覆盖前面的配置
func ParseFiles(filenames ...string) *Config {
	confmap := make(map[string]any)
	for _, file := range filenames {
		kv, err := readFile(file)
		if err != nil {
			panic(fmt.Sprintf("读取配置%s,错误:%q", file, err))
		}
		for k, v := range kv {
			confmap[k] = v
		}
	}
	var conf = new(Config)
	err := mapToStruct(confmap, conf, "env")
	if err != nil {
		panic("mapToStruct error:" + err.Error())
	}
	return conf
}

// readFile 读取文件，交给godotenv解析
func readFile(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return godotenv.Parse(file)
}

// mapToStruct map 对象转struct，支持嵌套
func mapToStruct(obj map[string]any, conf any, tagName string) error {
	sType := reflect.TypeOf(conf).Elem()
	sValue := reflect.ValueOf(conf).Elem()

	for i := 0; i < sType.NumField(); i++ {
		ft := sType.Field(i)
		fv := sValue.Field(i)

		// 判断是否是嵌套结构体，进行递归处理
		if fv.Kind() == reflect.Struct && fv.CanSet() {
			if err := mapToStruct(obj, fv.Addr().Interface(), tagName); err != nil {
				return err
			}
			continue
		}

		// 获取key
		mkey := ft.Name
		if key, ok := ft.Tag.Lookup(tagName); ok && strings.Trim(key, " ") != "" {
			mkey = key
		}
		mval, ok := obj[mkey]
		if !ok {
			continue
		}

		// 给字段赋值
		rVal := reflect.ValueOf(mval)
		if rVal.Type().ConvertibleTo(ft.Type) && fv.CanSet() {
			fv.Set(rVal.Convert(ft.Type))
		} else {
			switch ft.Type.String() {
			case "time.Duration":
				if err := setDuration(fv, rVal); err != nil {
					return err
				}
				continue // 跳过当次循环
			case "[]string":
				if err := setStringSlice(fv, rVal); err != nil {
					return err
				}
				continue // 跳过当次循环
			case "int":
				val, err := strconv.ParseInt(rVal.String(), 10, 64)
				if err == nil {
					fv.Set(reflect.ValueOf(int(val)))
					continue // 跳过当次循环
				}
			case "bool":
				val, err := strconv.ParseBool(rVal.String())
				if err == nil {
					fv.Set(reflect.ValueOf(bool(val)))
					continue // 跳过当次循环
				}
			}

			return fmt.Errorf("无法将 %v 转换为 %v", rVal.Type(), ft.Type)
		}
	}

	return nil
}

func setDuration(fv reflect.Value, rVal reflect.Value) error {
	dur, err := time.ParseDuration(rVal.String())
	if err != nil {
		return fmt.Errorf("time.ParseDuration error: %v", err)
	}
	fv.Set(reflect.ValueOf(dur))
	return nil
}

func setStringSlice(fv reflect.Value, rVal reflect.Value) error {
	// strings.Trim(s string, cutset string) string 其中 cutset 是集合，而不是单个字符串
	str := strings.Trim(rVal.String(), "[]") // 去除 "[" 和 "]"
	arr := strings.Split(str, ",")
	fv.Set(reflect.ValueOf(arr))
	return nil
}
