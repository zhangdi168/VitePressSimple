package setting

import "wailstemplate/application/db/entity"

const Version = "v1.00"
const GinPort = "9874"
const AppName = "VitePressSimple"

// EntityAutoMigrateList 自动迁移的实体列表
// var EntityAutoMigrateList = []any{new(entity.User)}
var EntityAutoMigrateList = []any{new(entity.Article)}
