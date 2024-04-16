package setting

import "wailstemplate/application/db/entity"

const Version = "v1.0-beta.1"

const AppName = "VitePressSimple"
const GitRepo = "VitePressSimple"
const GitAuthor = "zhangdi168"
const WindowZipContainStr = "windows_"
const MacosZipContainStr = "mac_"

// EntityAutoMigrateList 自动迁移的实体列表
// var EntityAutoMigrateList = []any{new(entity.User)}
var EntityAutoMigrateList = []any{new(entity.Article)}
