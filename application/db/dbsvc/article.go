package dbsvc

import (
	"fmt"
	"gorm.io/gorm"
	"wailstemplate/application/db/dbdto"
	"wailstemplate/application/db/entity"
	"wailstemplate/application/db/repository"
	"wailstemplate/application/pkg/utils"
	"wailstemplate/application/pkg/where"
)

type ArticleSvc struct {
	dbWrite *gorm.DB
	dbRead  *gorm.DB
}

func NewArticleSvc(dbWrite_, dbRead_ *gorm.DB) *ArticleSvc {
	return &ArticleSvc{
		dbWrite: dbWrite_,
		dbRead:  dbRead_,
	}
}

// Create 创建数据
func (s *ArticleSvc) Create(dtoCreate *dbdto.ArticleCreate) (*dbdto.ArticleInfo, error) {
	info := new(dbdto.ArticleInfo)
	err := repository.CreateWithDb(dtoCreate, new(entity.Article), info, s.dbWrite)
	if err != nil {
		return nil, err
	}
	if info == nil || info.ID <= 0 {
		return nil, fmt.Errorf("创建失败，info数据为空")
	}
	return info, err
}

// FindOne 查询一条数据
func (s *ArticleSvc) FindOne(wheres []*where.Field) (*dbdto.ArticleInfo, error) {
	entityInfo := new(entity.Article)
	err := repository.FindOneWithDb(wheres, entityInfo, false, s.dbRead)
	if err != nil {
		return nil, err
	}
	//转成dto
	dtoInfo := new(dbdto.ArticleInfo)
	utils.DtoToEntity(entityInfo, dtoInfo)
	if dtoInfo == nil || dtoInfo.ID <= 0 {
		return nil, fmt.Errorf("查询到的数据为空")
	}
	return dtoInfo, nil
}

// FindList 查询列表数据
func (s *ArticleSvc) FindList(wheres []*where.Field, extra *where.Extra) ([]*dbdto.ArticleInfo, error) {
	var entityList []*entity.Article
	//传入的entityList必须要加 &取地址符号，切片本身并不是指针，向函数传递一个切片时，实际上是复制了该切片的结构体
	err := repository.FindListWithDb(wheres, extra, &entityList, false, s.dbRead)
	if err != nil {
		return nil, err
	}
	//开始转成dto  list
	var dtoInfoList []*dbdto.ArticleInfo
	for _, entityInfo := range entityList {
		dtoInfo := new(dbdto.ArticleInfo)
		utils.DtoToEntity(entityInfo, dtoInfo)
		dtoInfoList = append(dtoInfoList, dtoInfo)
	}

	return dtoInfoList, nil
}

// Update 更新数据
func (s *ArticleSvc) Update(wheres []*where.Field, dtoUpdate *dbdto.ArticleUpdate, columnsCfg ...string) error {
	err := repository.UpdateWithDb(wheres, new(entity.Article), dtoUpdate, s.dbWrite, columnsCfg...)
	return err
}

// Delete 删除数据
func (s *ArticleSvc) Delete(wheres []*where.Field) error {
	err := repository.DeleteWithDb(wheres, new(entity.Article), false, s.dbWrite)
	return err
}
