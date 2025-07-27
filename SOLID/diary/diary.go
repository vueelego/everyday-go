package diary

import "time"

// Diary 定义一个日记结构体，存在基础数据字段
type Diary struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

func (d *Diary) GetTitle() string {
	return d.Title
}

func (d *Diary) Validate() bool {
	return len(d.Title) > 0 && len(d.Content) > 0
}

// NOTE: 咋一看，定义一个Save没有什么问题，其实是违反了单一职责原则；
// 混合了日记逻辑和数据库逻辑，需要拆分
// func (s *Diary) Save() error {
// 	return nil
// }

// =======================================================

type Repository interface {
	Save(d *Diary) error
}

type InMemoryRepository struct {
	data []*Diary
}

func (r *InMemoryRepository) Save(d *Diary) error {
	if r.data == nil {
		r.data = make([]*Diary, 1)
	}
	r.data = append(r.data, d)
	return nil
}

func SaveDiary(d *Diary, r Repository) error {
	return r.Save(d)
}

// =======================================================
// 依赖反转原则

// NOTE: 在这里，InMemoryRepository 是高模块，DiaryManager 是一个低模块
// 为什么这么说，因为 InMemoryRepository 从业务角度来讲，它没有依赖其他模块，
// InMemoryRepository 也符合单一职责设计。
// 这样写，非常不利于测试
type DiaryManager struct {
	store InMemoryRepository
}

func NewDiaryManager() DiaryManager {
	return DiaryManager{
		store: InMemoryRepository{},
	}
}

// 根据定义“高层模块不应该依赖低层模块，二者都应依赖于抽象”修改
// 修改后，如果我们要测试，就可以传递模拟存储库或者内存库来测试了

type DiaryManagerV2 struct {
	store Repository
}

func NewDiaryManageV2(store Repository) DiaryManagerV2 {
	return DiaryManagerV2{
		store: store,
	}
}
