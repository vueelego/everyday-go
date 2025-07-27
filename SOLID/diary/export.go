package diary

import "errors"

// NOTE: 在diary分别实现日记和存储功能；现在有个新的需求，
// 希望导出日志

// NOTE: ExportDiary 导出日志方式，看起来也没什么问题；
// 但是，如果要增加一个类型，就要在原方法上增加一个case；
// 再者，如果要修改百度网盘导出方法，也需要在同一个方法上修改；
// 这样很有可能影响到其他导出方式；
// 也有可能引入新的问题
// 因此可以使用 开闭原则 来解决这个问题
func ExportDiary(d *Diary, dst string) error {
	switch dst {
	case "硬盘":
		return nil
	case "百度网盘":
		return nil
	case "有道云笔记":
		return nil
	default:
		return errors.New("没有此导出方式")
	}
}

// ======================================

// NOTE: ExportDiaryV2 这里其实也就是 里氏替换原则，选中需要的导出类型，替换 exporter 即可；
// 将每种导出方式拆分，就是开闭原则
func ExportDiaryV2(d *Diary, exporter Exporter) error {
	return exporter.Export(d)
}

// Exporter 提取导出接口方法
type Exporter interface {
	Export(d *Diary) error
}

// HardDiskExport 硬盘导出结构体
type HardDiskExport struct {
}

func (e *HardDiskExport) Export(d *Diary) error {
	return nil
}

// BaiDuExport 百度网盘导出
type BaiDuExport struct {
}

func (e *BaiDuExport) Export(d *Diary) error {
	return nil
}

// 省略有道云导出...

// NOTE: 接口隔离模式，在Go语言中，几乎任何地方都可以使用，
// Go 的接口定义都是很小的，通常都是只有1-2个方法。
// Go 鼓励创建更小，更集中的接口，而不是大型接口。
// 通过接口组合需要的功能
// 同时Go语言，interface 是可以嵌套的，嵌套之后还可以定义其他方法，
// 这样组合起来就很灵活了
