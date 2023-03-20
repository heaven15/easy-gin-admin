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

type SysElementService struct{}

func (s *SysElementService) generateData(e *model.SysElement, data map[string]interface{}) error {
	data["real_name"] = e.RealName
	data["code"] = e.Code
	data["status"] = e.Status
	data["remark"] = e.Remark
	data["create_user"] = e.CreateUser
	return nil
}

func (s *SysElementService) Create(e *model.SysElement) error {
	var m model.SysElement
	if !errors.Is(global.EGVA_DB.Model(&m).Where("code = ?", e.Code).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysElementCodeAlreadyExists")
	}
	data := map[string]interface{}{}
	data["created_at"] = time.Now()
	if err := s.generateData(e, data); err != nil {
		return err
	}
	if err := global.EGVA_DB.Model(&m).Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysElementService) Update(e *model.SysElement) error {
	var m model.SysElement
	if errors.Is(global.EGVA_DB.Model(&m).First(&m, e.ID).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysElementInfoDoesNotExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("id != ? and code = ? ", e.ID, e.Code).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysElementCodeAlreadyExists")
	}
	data := map[string]interface{}{}
	if err := s.generateData(e, data); err != nil {
		return err
	}
	if err := global.EGVA_DB.Model(&m).Where("id = ?", e.ID).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysElementService) Delete(id []int) error {
	var m model.SysElement
	if err := global.EGVA_DB.Model(&m).Delete(&m, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysElementService) List(p *dto.PageReq) (*vo.PageDataVo, error) {
	var m model.SysElement
	var ms []model.SysElement
	localDB := global.EGVA_DB.Model(&m)
	if p.Keyword != "" {
		localDB = localDB.Where("code LIKE ? ", p.Keyword+"%")
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
