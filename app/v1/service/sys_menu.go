package service

import (
	"errors"
	"fmt"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/dto"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/model"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/vo"
	"github.com/develop-kevin/easy-gin-vue-admin/global"
	"github.com/develop-kevin/easy-gin-vue-admin/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"time"
)

type SysMenuService struct{}

func (s *SysMenuService) generateData(m *model.SysMenu, data map[string]interface{}) error {
	if m.ParentID > 0 {
		var sm model.SysMenu
		if errors.Is(global.EGVA_DB.Model(&sm).First(&sm, m.ParentID).Error, gorm.ErrRecordNotFound) {
			return errors.New("SysMenuInfoOfTheParentClassDoesNotExists")
		}
		data["parent_id"] = m.ParentID
	}
	data["real_name"] = m.RealName
	data["redirect"] = m.Redirect
	data["meta"] = m.Meta
	data["sort"] = m.Sort
	data["create_user"] = m.CreateUser
	return nil
}

func (s *SysMenuService) Create(m *model.SysMenu) error {
	var sm model.SysMenu
	if !errors.Is(global.EGVA_DB.Model(&sm).Where("real_name = ?", m.RealName).First(&sm).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysMenuNameAlreadyExist")
	}
	data := map[string]interface{}{}
	data["created_at"] = time.Now()
	if err := s.generateData(m, data); err != nil {
		return err
	}
	if err := global.EGVA_DB.Model(&sm).Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysMenuService) Update(m *model.SysMenu) error {
	var sm model.SysMenu
	if errors.Is(global.EGVA_DB.Model(&sm).First(&sm, m.ID).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysMenuInfoDoesNotExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&sm).Where("id != ? and real_name = ? ", m.ID, m.RealName).First(&sm).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysMenuNameAlreadyExists")
	}
	data := map[string]interface{}{}
	if err := s.generateData(m, data); err != nil {
		return err
	}
	if err := global.EGVA_DB.Model(&sm).Where("id = ?", m.ID).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysMenuService) Delete(id []int) error {
	if len(id) <= 0 {
		return errors.New("DeleteDataIsNull")
	}
	for _, v := range id {
		var sm model.SysMenu
		if !errors.Is(global.EGVA_DB.Model(&sm).Where("parent_id = ?", v).First(&sm).Error, gorm.ErrRecordNotFound) {
			return errors.New(fmt.Sprintf("SysMenuIsInUseBySubClasses_%v", v))
		}
	}
	var m model.SysMenu
	if err := global.EGVA_DB.Model(&m).Delete(&m, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysMenuService) Tree() []model.SysMenu {
	var ms []model.SysMenu
	global.EGVA_DB.Model(
		&model.SysMenu{}).Where("parent_id is NULL ").Where("status = ? ", utils.SUCCESS).Order(
		clause.OrderByColumn{
			Column: clause.Column{Name: "sort"},
			Desc:   true},
	).Preload("Children.Children").Find(&ms)
	return ms
}

func (s *SysMenuService) List(p *dto.PageReq) (*vo.PageDataVo, error) {
	var m model.SysMenu
	var ms []model.SysMenu
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
