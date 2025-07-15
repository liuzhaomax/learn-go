# Gorm

## Tips

### 1. 外键
  + 在一对多关系中，foreignKey指定的键永远指向多方
  + 在一对多关系中，foreignKey的指定要放在多方
    + Players []User \`gorm:"foreignKey:ID"\`
    + 上述ID是User的ID
  + 在一对多关系中，如果foreignKey的指定放在一方
    + RoleID uint \`gorm:"not null"\`
    + Role   Role \`gorm:"foreignKey:RoleID"\` // 与Role的多对一关系
    + 上述foreignKey的RoleID是本表RoleID
  + 一对多不需要references
  + 在一对一关系中，foreignKey指定的键指向本表，references指定的键指向关联表
    + UserID uint \`gorm:"unique;not null"\`
    + User   User \`gorm:"foreignKey:UserID;references:ID"\` // 与User的1:1关系
    + 上述foreignKey是本表UserID，references是关联表ID
  + 在多对多关系中，使用`gorm:"many2many:中间表;"`，两个表都要指定
    + Groups []Group \`gorm:"many2many:user_group;"\`

---

### 2. 错误处理

| 操作 | 错误                                                       | Model层处理                                                                                                           | Service层处理                               |
|:---:|:---------------------------------------------------------|:-------------------------------------------------------------------------------------------------------------------|:-----------------------------------------|
|查询| 错误包含`ErrRecordNotFound`，`RowsAffected`在DB宕机和查询不到的时候都是`0` | `if result.Error != nil {return result.Error}`                                                                     | `errors.Is(err, gorm.ErrRecordNotFound)` |
|创建| 错误包含`ErrDuplicatedKey`，`RowsAffected`在DB宕机和键重复的时候都是`0`   | `if result.Error != nil {return result.Error}`                                                                     | `errors.Is(err, gorm.ErrDuplicatedKey)`   |
|更新| 没有匹配的记录，或新值与旧值相同，或乐观锁冲突，会导致`RowsAffected`为`0`，但不会有错误 | `if result.Error != nil {return result.Error}` + `if result.RowsAffected == 0 { return gorm.ErrRecordNotFound }`   | `errors.Is(err, gorm.ErrRecordNotFound)`  |
|删除| 没有匹配的记录，或乐观锁冲突，会导致`RowsAffected`为`0`，但不会有错误 | `if result.Error != nil {return result.Error}` + `if result.RowsAffected == 0 { return gorm.ErrRecordNotFound }`   | `errors.Is(err, gorm.ErrRecordNotFound)` |

---

### 3. 更新的方法

**1. `Save`**
- **作用**：保存整个模型，更新所有字段（包括零值）。
- **特点**：
    - 若记录存在（主键有值），则更新所有字段；
    - 若记录不存在，则插入新记录。
- **示例**：
  ```go
  user := User{ID: 1, Name: "Alice"}  
  db.Save(&user)  // 更新所有字段（包括未修改的）  
  ```

**2. `Update`**
- **作用**：更新单个字段（需指定列名和值）。
- **特点**：
    - 仅更新指定字段，忽略其他字段；
    - 必须搭配条件（如 `Where`）使用。
- **示例**：
  ```go
  db.Model(&User{}).Where("id = ?", 1).Update("name", "Bob")  
  ```

**3. `Updates`**
- **作用**：批量更新多个字段（支持结构体或 `map`）。
- **特点**：
    - 仅更新非零值字段（若用 `map` 则更新所有指定字段）；
    - 可搭配条件使用。
- **示例**：
  ```go
  db.Model(&User{}).Where("id = ?", 1).Updates(User{Name: "Bob"}) // 忽略零值  
  db.Model(&User{}).Where("id = ?", 1).Updates(map[string]any{"name": "Bob", "age": 0}) // 强制更新零值  
  ```

### **核心区别总结**
| 方法       | 适用场景               | 字段更新规则           | 零值处理       |  
|------------|-----------------------|-----------------------|---------------|  
| `Save`     | 全量保存记录           | 更新所有字段           | 零值覆盖       |  
| `Update`   | 更新单个字段           | 仅更新指定字段         | 显式指定       |  
| `Updates`  | 批量更新部分字段       | 默认忽略零值（结构体） | `map` 可更新零值 |  

**选择建议**：
- 用 `Save` 谨慎（可能覆盖未修改字段）；
- 用 `Updates` 更安全（推荐结构体更新）。

### 4. 删除的方法

**注意**
```
删除是软删除，除非用 Unscoped()
```

### 5. 创建的方法

**注意**
```
使用Create，不要使用FirstOrCreate，它包含查询会让错误处理混乱
```
