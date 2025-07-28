# visualgo

好记性不如烂键盘，取名来自visualgo.net，这是一个学习数据结构的好网站。

### 更改go.mod依赖go的版本

```bash
go mod edit -go=version

# 例如:
go mod edit -go=1.22.5
```

### Trie 数据结构

Trie（读作「try」，又称前缀树或字典树）是一种非常高效的字符串检索数据结构，在很多搜索、自动补全、前缀匹配等应用中广泛使用。

**常规操作**
- 插入字符串（Insert）
- 查询字符串（Search）
- 前缀查询（StartsWith）

