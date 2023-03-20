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

type SysPermissionService struct{}

func (s *SysPermissionService) generateData(p *model.SysPermission, data map[string]interface{}) error {
	data["real_name"] = p.RealName
	data["code"] = p.Code
	data["status"] = p.Status
	data["remark"] = p.Remark
	data["create_user"] = p.CreateUser
	return nil
}

func (s *SysPermissionService) Create(p *model.SysPermission) error {
	var m model.SysPermission
	if !errors.Is(global.EGVA_DB.Model(&m).Where("real_name = ?", p.RealName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysPermissionNameAlreadyExists")
	}
	data := map[string]interface{}{}
	data["created_at"] = time.Now()
	if err := s.generateData(p, data); err != nil {
		return err
	}
	if err := global.EGVA_DB.Model(&m).Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysPermissionService) Update(p *model.SysPermission) error {
	var m model.SysPermission
	if errors.Is(global.EGVA_DB.Model(&m).First(&m, p.ID).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysPermissionInfoDoesNotExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("id != ? and real_name = ? ", p.ID, p.RealName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysPermissionNameAlreadyExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("code = ?", p.Code).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysOperationCodeAlreadyExists")
	}
	data := map[string]interface{}{}
	if err := s.generateData(p, data); err != nil {
		return err
	}
	if err := global.EGVA_DB.Model(&m).Where("id = ?", p.ID).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysPermissionService) Delete(id []int) error {
	if len(id) <= 0 {
		return errors.New("DeleteDataIsNull")
	}
	for _, v := range id {
		var p model.SysPage
		if !errors.Is(global.EGVA_DB.Model(&p).Where("permission_id = ?", v).First(&p).Error, gorm.ErrRecordNotFound) {
			return errors.New(fmt.Sprintf("SysPermissionIsInUsePage_%v", v))
		}
	}
	var m model.SysPermission
	if err := global.EGVA_DB.Model(&m).Delete(&m, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysPermissionService) List(p *dto.PageReq) (*vo.PageDataVo, error) {
	var m model.SysPermission
	var ms []model.SysPermission
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
