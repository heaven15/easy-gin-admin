package service

import (
	"fmt"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/dto"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/model"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/vo"
	"github.com/develop-kevin/easy-gin-vue-admin/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"time"
)

type SysPageService struct{}

func (s *SysPageService) generateData(p *model.SysPage, data map[string]interface{}) error {
	if p.PermissionID > 0 {
		var mp model.SysPermission
		if errors.Is(global.EGVA_DB.Model(&mp).First(&mp, p.PermissionID).Error, gorm.ErrRecordNotFound) {
			return errors.New("SysPermissionInfoDoesNotExists")
		}
		data["permission_id"] = p.PermissionID
	}
	data["real_name"] = p.RealName
	data["url"] = p.Url
	data["sort"] = p.Sort
	data["is_type"] = p.IsType
	data["remark"] = p.Remark
	data["status"] = p.Status
	return nil
}

func (s *SysPageService) Create(p *model.SysPage) error {
	var mp model.SysPage
	if !errors.Is(global.EGVA_DB.Model(&mp).Where("real_name = ?", p.RealName).First(&mp).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysPageNameAlreadyExists")
	}
	data := map[string]interface{}{}
	if err := s.generateData(p, data); err != nil {
		return err
	}
	data["created_at"] = time.Now()
	if err := global.EGVA_DB.Model(&mp).Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysPageService) Update(p *model.SysPage) error {
	var mp model.SysPage
	if errors.Is(global.EGVA_DB.Model(&mp).First(&mp, p.ID).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysPageInfoDoesNotExists")
	}
	if errors.Is(global.EGVA_DB.Model(&mp).Where("id != ? and real_name = ?", p.ID, p.RealName).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysPageNameAlreadyExists")
	}
	data := map[string]interface{}{}
	if err := s.generateData(p, data); err != nil {
		return err
	}
	if err := global.EGVA_DB.Model(&mp).Where("id=?", p.ID).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysPageService) Delete(id []int) error {
	if len(id) <= 0 {
		return errors.New("DeleteDataIsNull")
	}
	for _, v := range id {
		var m model.SysMenu
		if !errors.Is(global.EGVA_DB.Model(&m).Where("page_id = ?", v).First(&m).Error, gorm.ErrRecordNotFound) {
			return errors.New(fmt.Sprintf("SysMenuIsInUseBySubClasses_%v", v))
		}
	}
	var p model.SysPage
	if err := global.EGVA_DB.Model(&p).Delete(&p, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysPageService) List(p *dto.PageReq) (*vo.PageDataVo, error) {
	var m model.SysPage
	var ms []model.SysPage
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
		Page:     p.Page,
		PageSize: p.PageSize,
		Total:    total,
		Data:     ms,
	}, nil
}
