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

type SysGroupService struct{}

func (s *SysGroupService) generateData(g *model.SysGroup, data map[string]interface{}) error {
	if g.ParentID > 0 {
		var m model.SysGroup
		if errors.Is(global.EGVA_DB.Model(&m).First(&m, g.ParentID).Error, gorm.ErrRecordNotFound) {
			return errors.New("SysGroupInfoOfTheParentClassDoesNotExists")
		}
	}
	data["parent_id"] = g.ParentID
	data["real_name"] = g.RealName
	data["remark"] = g.Remark
	data["sort"] = g.Sort
	data["status"] = g.Status
	data["remark"] = g.Remark
	data["create_user"] = g.CreateUser
	return nil
}

func (s *SysGroupService) Create(g *model.SysGroup) error {
	var m model.SysGroup
	if !errors.Is(global.EGVA_DB.Model(&m).Where("real_name = ?", g.RealName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysGroupNameAlreadyExist")
	}
	data := map[string]interface{}{}
	data["created_at"] = time.Now()
	if err := s.generateData(g, data); err != nil {
		return err
	}
	if err := global.EGVA_DB.Model(&m).Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysGroupService) Update(g *model.SysGroup) error {
	var m model.SysGroup
	if errors.Is(global.EGVA_DB.Model(&m).First(&m, g.ID).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysGroupInfoDoesNotExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("id != ? and real_name = ? ", g.ID, g.RealName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysGroupNameAlreadyExists")
	}
	data := map[string]interface{}{}
	if err := s.generateData(g, data); err != nil {
		return err
	}
	if err := global.EGVA_DB.Model(&m).Where("id = ?", g.ID).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysGroupService) Delete(id []int) error {
	if len(id) <= 0 {
		return errors.New("DeleteDataIsNull")
	}
	for _, v := range id {
		var sm model.SysGroup
		if !errors.Is(global.EGVA_DB.Model(&sm).Where("parent_id = ?", v).First(&sm).Error, gorm.ErrRecordNotFound) {
			return errors.New(fmt.Sprintf("SysGroupIsInUseBySubClasses_%v", v))
		}
	}
	var m model.SysGroup
	if err := global.EGVA_DB.Model(&m).Delete(&m, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysGroupService) List(p *dto.PageReq) (*vo.PageDataVo, error) {
	var m model.SysGroup
	var ms []model.SysGroup
	localDB := global.EGVA_DB.Model(&m)
	if p.Keyword != "" {
		localDB = localDB.Where("real_name LIKE ? ", p.Keyword+"%")
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
