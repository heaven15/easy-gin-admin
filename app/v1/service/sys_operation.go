package service

import (
	"errors"
	"fmt"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/dto"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/model"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/vo"
	"github.com/develop-kevin/easy-gin-vue-admin/global"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"time"
)

type SysOperationService struct{}

func (s *SysOperationService) generateData(d *model.SysOperation, data map[string]interface{}) error {
	if d.ParentID > 0 {
		var m model.SysOperation
		if errors.Is(global.EGVA_DB.Model(&m).First(&m, d.ParentID).Error, gorm.ErrRecordNotFound) {
			return errors.New("SysOperationOfTheParentClassDoesNotExists")
		}
		data["parent_id"] = d.ParentID
	}
	data["real_name"] = d.RealName
	data["code"] = d.Code
	data["url"] = d.Url
	data["status"] = d.Status
	data["remark"] = d.Remark
	data["create_user"] = d.CreateUser
	return nil
}

func (s *SysOperationService) Create(d *model.SysOperation) error {
	var m model.SysOperation
	if !errors.Is(global.EGVA_DB.Model(&m).Where("real_name = ?", d.RealName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysOperationNameAlreadyExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("code = ?", d.Code).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysOperationCodeAlreadyExists")
	}
	data := map[string]interface{}{}
	data["created_at"] = time.Now()
	if err := s.generateData(d, data); err != nil {
		return err
	}
	if err := global.EGVA_DB.Model(&m).Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysOperationService) Update(d *model.SysOperation) error {
	var m model.SysOperation
	if errors.Is(global.EGVA_DB.Model(&m).First(&m, d.ID).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysOperationInfoDoesNotExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("id != ? and real_name = ? ", d.ID, d.RealName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysOperationNameAlreadyExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("code = ?", d.Code).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysOperationCodeAlreadyExists")
	}
	data := map[string]interface{}{}
	if err := s.generateData(d, data); err != nil {
		return err
	}
	if err := global.EGVA_DB.Model(&m).Where("id = ?", d.ID).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysOperationService) Delete(id []int) error {
	if len(id) <= 0 {
		return errors.New("DeleteDataIsNull")
	}
	for _, v := range id {
		var sm model.SysOperation
		if !errors.Is(global.EGVA_DB.Model(&sm).Where("parent_id = ?", v).First(&sm).Error, gorm.ErrRecordNotFound) {
			return errors.New(fmt.Sprintf("SysOperationIsInUseBySubClasses_%v", v))
		}
	}
	var m model.SysOperation
	if err := global.EGVA_DB.Model(&m).Delete(&m, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysOperationService) List(p *dto.PageReq) (*vo.PageDataVo, error) {
	var m model.SysOperation
	var ms []model.SysOperation
	localDB := global.EGVA_DB.Model(&m)
	if p.Keyword != "" {
		localDB = localDB.Where("real_name LIKE ? OR full_name LIKE ? ", p.Keyword+"%", p.Keyword+"%")
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
