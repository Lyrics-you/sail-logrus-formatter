# **sail**-logrus-formatter
Logurs' Formatter Project : Provides "yyyy-MM-dd" date formatting, log level colors, and positioning.

![shadow-zoom](https://image-taragrade.oss-cn-hangzhou.aliyuncs.com/imagehub/image-20220213145846973.png)

## 介绍

参考了：[nested-logrus-formatter](https://github.com/antonfisher/nested-logrus-formatter)

区别：

1. 增加了文件定位功能，提供[file:line]格式。总体格式更符合SailAnchor
2. 日期格式化支持"yyyy-MM-dd HH:mm:ss.SSS zzz"格式，(yyyy:year,MM:month,dd:day,:HH:hour,mm:minute,ss:second,SSS:microsecond,zzz:time zone)
3. 颜色等级选择共：Black\Red\Green\Yellow\Blue\Magenta\Cyan\Gray等
4. 属性更符合认知：NoColor->Color等。

## 说明

## 

Formatter struct

| 字段                  | 属性                        | 说明                                          |
| --------------------- | --------------------------- | --------------------------------------------- |
| FieldsOrder           | string                      | 按字母顺序排列的字段                          |
| TimeStampFormat       | string                      | "2006-01-02 15:04:05.000 MST"                 |
| CharStampFormat       | string                      | "yyyy-MM-dd hh:mm:ss.SSS zzz"                 |
| HideKeys              | bool                        | 显示[fieldValue]而不是[fieldKey:fieldValue]。 |
| Position              | bool                        | 显示[文件:行数]                               |
| Colors                | bool                        | 显示颜色                                      |
| FieldsColors          | bool                        | 显示字段颜色                                  |
| FieldsSpace           | bool                        | 字段中间是否添加空格                          |
| ShowFullLevel         | bool                        | 显示等级的全称：[WARNING]取代[WARN]           |
| LowerCaseLevel        | bool                        | 等级小写                                      |
| TrimMessages          | bool                        | 修剪信息中的空白处                            |
| CallerFirst           | bool                        | 先打印调用信息                                |
| CustomCallerFormatter | func(*runtime.Frame) string | 为调用信息设置自定义格式                      |

提供time的时间格式化供参考

| 名称        | 格式                                  |
| ----------- | ------------------------------------- |
| ANSIC       | "Mon Jan _2 15:04:05 2006"            |
| UnixDate    | "Mon Jan _2 15:04:05 MST 2006"        |
| RubyDate    | "Mon Jan 02 15:04:05 -0700 2006"      |
| RFC822      | "02 Jan 06 15:04 MST"                 |
| RFC822Z     | "02 Jan 06 15:04 -0700"               |
| RFC850      | "Monday, 02-Jan-06 15:04:05 MST"      |
| RFC112      | "Mon, 02 Jan 2006 15:04:05 MST"       |
| RFC1123Z    | "Mon, 02 Jan 2006 15:04:05 -0700"     |
| RFC3339     | "2006-01-02T15:04:05Z07:00"           |
| RFC3339Nano | "2006-01-02T15:04:05.999999999Z07:00" |
| Kitchen     | "3:04PM"                              |
| Stamp       | Jan _2 15:04:05"                      |
| StampMilli  | "Jan _2 15:04:05.000"                 |
| StampMicro  | "Jan _2 15:04:05.000000"              |
| StampNano   | "Jan _2 15:04:05.000000000"           |
| StampNormal | "2006-01-02 15:04:05.000 MST"         |

## 使用

## 

```go
import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

log := logrus.New()
log.SetFormatter(&nested.Formatter{
	HideKeys:    true,
	FieldsOrder: []string{"component", "category"},
})

log.Info("just info message")
log.WithField("component", "rest").Warn("warn message")
```

