package setting

import "wailstemplate/application/db/entity"

const Version = "v1.00"
const GinPort = "9874"

// ServeDir 服务端路径
const ServeDir = "D:\\data\\xiaod-admin\\server"

//const ServeDir = "D:\\project\\geeker-admin-go\\server"

//const ServeDir = "/Users/xiaod/project/geeker-admin-go/server"

// FrontDir 前端路径
const FrontDir = "D:\\data\\xiaod-admin\\front"

//const FrontDir = "/Users/xiaod/project/geeker-admin-go/front"
//
//const FrontDir = "D:\\project\\geeker-admin-go\\front"

// UseSoftDelete 是否使用软删除功能如果true 表必须存在deleted_at字段
const UseSoftDelete = false

const UpdateExeName = "QrPubClient"

// EntityAutoMigrateList 自动迁移的实体列表
// var EntityAutoMigrateList = []any{new(entity.User)}
var EntityAutoMigrateList = []any{new(entity.Article)}
