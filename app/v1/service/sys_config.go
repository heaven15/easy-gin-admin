package service

import (
	"errors"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/dto"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/model"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/vo"
	"github.com/develop-kevin/easy-gin-vue-admin/global"
	"gorm.io/gorm"
	"time"
)

type SysConfigService struct{}

func (s *SysConfigService) generateData(c *model.SysConfig, data map[string]interface{}) error {
	data["real_name"] = c.RealName
	data["app_key"] = c.AppKey
	data["app_val"] = c.AppVal
	data["remark"] = c.Remark
	data["create_user"] = c.CreateUser
	data["permission_code"] = c.PermissionCode
	return nil
}

func (s *SysConfigService) Create(c *model.SysConfig) error {
	var m model.SysConfig
	if !errors.Is(global.EGVA_DB.Model(&m).Where("real_name = ?", c.RealName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysConfigNameAlreadyExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("app_key = ?", c.AppKey).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysConfigAppKeyAlreadyExists")
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

func (s *SysConfigService) Update(c *model.SysConfig) error {
	var m model.SysConfig
	if errors.Is(global.EGVA_DB.Model(&m).First(&m, c.ID).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysConfigInfoDoesNotExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("id != ? and real_name = ? ", c.ID, c.RealName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysConfigNameAlreadyExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("id != ? and app_key = ? ", c.ID, c.AppKey).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysConfigAppKeyAlreadyExists")
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

func (s *SysConfigService) QueryKey(key string) (*model.SysConfig, error) {
	var m model.SysConfig
	if errors.Is(global.EGVA_DB.Model(&m).Where("app_key = ?", key).First(&m).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("SysConfigInfoDoesNotExists")
	}
	return &m, nil
}

func (s *SysConfigService) Delete(id []int) error {
	var m model.SysConfig
	if err := global.EGVA_DB.Model(&m).Delete(&m, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysConfigService) List(p *dto.PageReq) (*vo.PageDataVo, error) {
	var m model.SysConfig
	var ms []model.SysConfig
	localDB := global.EGVA_DB.Model(&m)
	if p.Keyword != "" {
		localDB = localDB.Where("real_name LIKE ? OR app_key LIKE ? ", p.Keyword+"%", p.Keyword+"%")
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
