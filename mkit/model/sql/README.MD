# Mkit Dao

mkit dao fork to go-zero。

## 快速开始

* 通过datasource生成

    ```shell script
    mkit dao mysql datasource -url="user:password@tcp(127.0.0.1:3306)/database" -table="user_*"  -dir="./dao"  -gitlab="code.xxx.com" -repo="xxx"
    ```

* 生成代码示例

```go
package dao

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	appzaplog "your.gitlab.com/gomod/utils/ulog"
	utildao "your.gitlab.com/your-repository/server/src/util/dao"
	"your.gitlab.com/your-repository/server/src/util/uredis"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	userXxxxxLogFieldNames          = utildao.FieldNames(&UserXxxxxLog{})
	userXxxxxLogRows                = strings.Join(userXxxxxLogFieldNames, ",")
	userXxxxxLogRowsExpectAutoSet   = strings.Join(utildao.Remove(userXxxxxLogFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userXxxxxLogRowsWithPlaceHolder = strings.Join(utildao.Remove(userXxxxxLogFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	UserXxxxxLogDao interface {
		Insert(data *UserXxxxxLog) (sql.Result, error)
		FindOne(id int64) (*UserXxxxxLog, error)
		Update(data UserXxxxxLog) error
		Delete(id int64) error
	}

	defaultUserXxxxxLogDao struct {
		mizaDB      *gorm.DB
		redisClient *uredis.Client
		table       string
	}

	UserXxxxxLog struct {
		Id           int64         `gorm:"column:id" db:"id"`                         // 主键id
		RelationId   int64         `gorm:"column:relation_id" db:"relation_id"`       // 师徒关系id
		PlayerId     int64         `gorm:"column:player_id" db:"player_id"`           // player_id
		PlayerName   string        `gorm:"column:player_name" db:"player_name"`       // 昵称
		TaskType     int8          `gorm:"column:task_type" db:"task_type"`           // 完成任务类型:1初级任务2进阶任务3高级任务4特殊任务
		AwardType    int8          `gorm:"column:award_type" db:"award_type"`         // 奖励类型:1金币2礼物3特效
		GiftId       sql.NullInt64 `gorm:"column:gift_id" db:"gift_id"`               // 礼物/特效id
		AwardNum     int32         `gorm:"column:award_num" db:"award_num"`           // 奖励数量
		CreateTime   time.Time     `gorm:"column:create_time" db:"create_time"`       // 创建时间
		TaskDetailId int64         `gorm:"column:task_detail_id" db:"task_detail_id"` // 任务id,等级任务为0
	}
)

func (obj UserXxxxxLog) TableName() string {
	return "user_xxxxx_log"
}

func newUserXxxxxLogDao() UserXxxxxLogDao {
	return &defaultUserXxxxxLogDao{
		mizaDB:      mainDB,
		redisClient: RedisClient,
		table:       "user_xxxxx_log",
	}
}

func (m *defaultUserXxxxxLogDao) Insert(data *UserXxxxxLog) (sql.Result, error) {
	appzaplog.Debug("UserXxxxxLogDao Insert start", zap.Any("data", data))
	create := m.mizaDB.Create(data)
	err := create.Error
	if err != nil {
		appzaplog.Error("UserXxxxxLogDao Insert error", zap.Error(err))
		return driver.RowsAffected(0), err
	}
	// todo 如需增加缓存自己增加
	appzaplog.Debug("UserXxxxxLogDao Insert", zap.Any("Rows", create.RowsAffected), zap.Any("data", data))
	return driver.RowsAffected(create.RowsAffected), nil
}

func (m *defaultUserXxxxxLogDao) FindOne(id int64) (*UserXxxxxLog, error) {
	appzaplog.Debug("UserXxxxxLogDao FindOne", zap.Any("id", id))
	// todo 如需查询缓存自己增加
	resp := new(UserXxxxxLog)
	err := m.mizaDB.Where("id = ?", id).First(resp).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	appzaplog.Debug("UserXxxxxLogDao FindOne", zap.Any("resp", resp))
	return resp, err
}

func (m *defaultUserXxxxxLogDao) Update(data UserXxxxxLog) error {
	appzaplog.Debug("UserXxxxxLogDao Update start", zap.Any("data", data))
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userXxxxxLogRowsWithPlaceHolder)
	err := m.mizaDB.Exec(query, data.RelationId, data.PlayerId, data.PlayerName, data.TaskType, data.AwardType, data.GiftId, data.AwardNum, data.TaskDetailId, data.Id).Error
	// todo 如需清理缓存自己增加
	return err
}

func (m *defaultUserXxxxxLogDao) Delete(id int64) error {
	appzaplog.Debug("UserXxxxxLogDao Delete", zap.Any("id", id))
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	err := m.mizaDB.Exec(query, id).Error
	// todo 如需清理缓存自己增加
	return err
}

```

## 用法

```Plain Text
mkit dao mysql -h
```

```Plain Text
NAME:
   mkit dao mysql - generate mysql dao"

USAGE:
   mkit dao mysql command [command options] [arguments...]

COMMANDS:
   ddl         generate mysql dao from ddl"(已去除)
   datasource  generate dao from datasource"

OPTIONS:
   --help, -h  show help
```

## 生成规则

* 默认规则

  我们默认用户在建表时会创建createTime、updateTime字段(忽略大小写、下划线命名风格)且默认值均为`CURRENT_TIMESTAMP`
  ，而updateTime支持`ON UPDATE CURRENT_TIMESTAMP`，对于这两个字段生成`insert`、`update`时会被移除，不在赋值范畴内，当然，如果你不需要这两个字段那也无大碍。
* 带缓存模式
    * ddl

      ```shell script
      mkit dao mysql -src={patterns} -dir={dir} -cache
      ```

      help

      ```
      NAME:
         mkit dao mysql ddl - generate mysql dao from ddl
      
      USAGE:
         mkit dao mysql ddl [command options] [arguments...]
      
      OPTIONS:
         --src value, -s value  the path or path globbing patterns of the ddl
         --dir value, -d value  the target dir
         --style value          the file naming format, see [https://github.com/tal-tech/go-zero/tree/master/tools/mkit/config/readme.md]
         --cache, -c            generate code with cache [optional]
         --idea                 for idea plugin [optional]
         --database, -db        the name of database [optional]
      ```

    * datasource

      ```shell script
      mkit dao mysql datasource -url={datasource} -table={patterns}  -dir={dir} -cache=true
      ```

      help

      ```
      NAME:
         mkit dao mysql datasource - generate dao from datasource
      
      USAGE:
         mkit dao mysql datasource [command options] [arguments...]
      
      OPTIONS:
         --url value              the data source of database,like "root:password@tcp(127.0.0.1:3306)/database
         --table value, -t value  the table or table globbing patterns in the database
         --cache, -c              generate code with cache [optional]
         --dir value, -d value    the target dir
         --style value            the file naming format, see [https://github.com/tal-tech/go-zero/tree/master/tools/mkit/config/readme.md]
         --idea                   for idea plugin [optional]
  
  
      ```

      示例用法请参考[用法](./example/generator.sh)

      > NOTE: mkit dao mysql ddl/datasource 均新增了一个`--style`参数，用于标记文件命名风格。

  目前仅支持redis缓存，如果选择带缓存模式，即生成的`FindOne(ByXxx)`&`Delete`
  代码会生成带缓存逻辑的代码，目前仅支持单索引字段（除全文索引外），对于联合索引我们默认认为不需要带缓存，且不属于通用型代码，因此没有放在代码生成行列，如example中user表中的`id`、`name`、`mobile`
  字段均属于单字段索引。

* 不带缓存模式

    * ddl

        ```shell script
          mkit dao -src={patterns} -dir={dir}
        ```

    * datasource

        ```shell script
          mkit dao mysql datasource -url={datasource} -table={patterns}  -dir={dir}
        ```

  or
    * ddl

        ```shell script
          mkit dao -src={patterns} -dir={dir}
        ```

    * datasource

        ```shell script
          mkit dao mysql datasource -url={datasource} -table={patterns}  -dir={dir}
        ```

生成代码仅基本的CURD结构。

## 缓存

对于缓存这一块我选择用一问一答的形式进行罗列。我想这样能够更清晰的描述dao中缓存的功能。

* 缓存会缓存哪些信息？

  对于主键字段缓存，会缓存整个结构体信息，而对于单索引字段（除全文索引）则缓存主键字段值。

* 数据有更新（`update`）操作会清空缓存吗？

  会，但仅清空主键缓存的信息，why？这里就不做详细赘述了。

* 为什么不按照单索引字段生成`updateByXxx`和`deleteByXxx`的代码？

  理论上是没任何问题，但是我们认为，对于dao层的数据操作均是以整个结构体为单位，包括查询，我不建议只查询某部分字段（不反对），否则我们的缓存就没有意义了。

* 为什么不支持`findPageLimit`、`findAll`这么模式代码生层？

  目前，我认为除了基本的CURD外，其他的代码均属于<i>业务型</i>代码，这个我觉得开发人员根据业务需要进行编写更好。

# 类型转换规则

| mysql dataType | golang dataType | golang dataType(if null&&default null) |
|----------------|-----------------|----------------------------------------|
| bool           | uint8           | sql.NullInt64                          |
| boolean        | uint8           | sql.NullInt64                          |
| tinyint        | int8            | sql.NullInt64                          |
| smallint       | int16           | sql.NullInt64                          |
| mediumint      | int32           | sql.NullInt64                          |
| int            | int32           | sql.NullInt64                          |
| integer        | int32           | sql.NullInt64                          |
| bigint         | int64           | sql.NullInt64                          |
| float          | float64         | sql.NullFloat64                        |
| double         | float64         | sql.NullFloat64                        |
| decimal        | float64         | sql.NullFloat64                        |
| date           | time.Time       | sql.NullTime                           |
| datetime       | time.Time       | sql.NullTime                           |
| timestamp      | time.Time       | sql.NullTime                           |
| time           | string          | sql.NullString                         |
| year           | time.Time       | sql.NullInt64                          |
| char           | string          | sql.NullString                         |
| varchar        | string          | sql.NullString                         |
| binary         | string          | sql.NullString                         |
| varbinary      | string          | sql.NullString                         |
| tinytext       | string          | sql.NullString                         |
| text           | string          | sql.NullString                         |
| mediumtext     | string          | sql.NullString                         |
| longtext       | string          | sql.NullString                         |
| enum           | string          | sql.NullString                         |
| set            | string          | sql.NullString                         |
| json           | string          | sql.NullString                         |