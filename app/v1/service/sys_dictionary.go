package service

import (
	"errors"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/dto"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/model"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/vo"
	"github.com/develop-kevin/easy-gin-vue-admin/global"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"time"
)

type SysDictionaryService struct{}

func (s *SysDictionaryService) generateData(c *model.SysDictionary, data map[string]interface{}) error {
	data["real_name"] = c.RealName
	data["code"] = c.Code
	data["is_type"] = c.IsType
	data["remark"] = c.Remark
	data["status"] = c.Status
	data["create_user"] = c.CreateUser
	return nil
}

func (s *SysDictionaryService) Create(c *model.SysDictionary) error {
	var m model.SysDictionary
	if !errors.Is(global.EGVA_DB.Model(&m).Where("real_name = ?", c.RealName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysDictionaryNameAlreadyExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("code = ? ", c.Code).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysDictionaryCodeAlreadyExists")
	}
	data := map[string]interface{}{}
	data["created_at"] = time.Now()
	if err := s.generateData(c, data); err != nil {
		return err
	}
	if err := global.EGVA_DB.Model(&m).Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysDictionaryService) Update(c *model.SysDictionary) error {
	var m model.SysDictionary
	if errors.Is(global.EGVA_DB.Model(&m).First(&m, c.ID).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysDepartmentInfoDoesNotExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("id != ? and real_name = ? ", c.ID, c.RealName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysDictionaryNameAlreadyExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("id != ? and code = ? ", c.ID, c.Code).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysDictionaryCodeAlreadyExists")
	}
	data := map[string]interface{}{}
	if err := s.generateData(c, data); err != nil {
		return err
	}
	if err := global.EGVA_DB.Model(&m).Where("id = ?", c.ID).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysDictionaryService) Delete(id []int) error {
	var m model.SysDictionary
	if err := global.EGVA_DB.Model(&m).Delete(&m, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysDictionaryService) List(p *dto.PageReq) (*vo.PageDataVo, error) {
	var m model.SysDictionary
	var ms []model.SysDictionary
	localDB := global.EGVA_DB.Model(&m)
	if p.Keyword != "" {
		localDB = localDB.Where("real_name LIKE ? OR code LIKE ? ", p.Keyword+"%", p.Keyword+"%")
	}
	if p.Status > 0 {
		localDB = localDB.Where("status = ? ", p.Status)
	}
	if p.Order != "" {
		parseBool, _ := strconv.ParseBool(p.Sort)
		localDB = localDB.Order(clause.OrderByColumn{Column: clause.Column{Name: p.Order}, Desc: parseBool})
	}
	var total int64
	global.EGVA_DB.Model(&m).Count(&total)
	localDB.Scopes(model.Paginate(p.Page, p.PageSize)).Find(&ms)
	return &vo.PageDataVo{
		Total:    total,
		Data:     ms,
		Page:     p.Page,
		PageSize: p.PageSize,
	}, nil
}
