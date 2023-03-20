package service

import (
	"errors"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/dto"
	model "github.com/develop-kevin/easy-gin-vue-admin/app/v1/model"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/vo"
	"github.com/develop-kevin/easy-gin-vue-admin/global"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"time"
)

type SysFileService struct{}

func (s *SysFileService) generateData(f *model.SysFile, data map[string]interface{}) error {
	data["real_name"] = f.RealName
	data["path"] = f.Path
	data["status"] = f.Status
	data["remark"] = f.Remark
	data["create_user"] = f.CreateUser
	return nil
}

func (s *SysFileService) Create(f *model.SysFile) error {
	var m model.SysFile
	if !errors.Is(global.EGVA_DB.Model(&m).Where("real_name = ?", f.RealName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysFileNameAlreadyExists")
	}
	data := map[string]interface{}{}
	data["created_at"] = time.Now()
	if err := s.generateData(f, data); err != nil {
		return err
	}
	if err := global.EGVA_DB.Model(&m).Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysFileService) Update(f *model.SysFile) error {
	var m model.SysFile
	if errors.Is(global.EGVA_DB.Model(&m).First(&m, f.ID).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysFileInfoDoesNotExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("id != ? and real_name = ? ", f.ID, f.RealName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysFileNameAlreadyExists")
	}
	data := map[string]interface{}{}
	if err := s.generateData(f, data); err != nil {
		return err
	}
	if err := global.EGVA_DB.Model(&m).Where("id = ?", f.ID).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysFileService) Delete(id []int) error {
	var m model.SysFile
	if err := global.EGVA_DB.Model(&m).Delete(&m, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysFileService) List(p *dto.PageReq) (*vo.PageDataVo, error) {
	var m model.SysFile
	var ms []model.SysFile
	localDB := global.EGVA_DB.Model(&m)
	if p.Keyword != "" {
		localDB = localDB.Where("real_name LIKE ? ", p.Keyword+"%", p.Keyword+"%")
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
