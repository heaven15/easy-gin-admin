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

type SysPostService struct{}

func (s *SysPostService) generateData(c *model.SysPost, data map[string]interface{}) error {
	data["real_name"] = c.RealName
	data["sort"] = c.Sort
	data["status"] = c.Status
	data["remark"] = c.Remark
	data["create_user"] = c.CreateUser
	return nil
}

func (s *SysPostService) Create(c *model.SysPost) error {
	var m model.SysPost
	if !errors.Is(global.EGVA_DB.Model(&m).Where("real_name = ?", c.RealName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysPostNameAlreadyExist")
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

func (s *SysPostService) Update(c *model.SysPost) error {
	var m model.SysPost
	if errors.Is(global.EGVA_DB.Model(&m).First(&m, c.ID).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysPostInfoDoesNotExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("id != ? and real_name = ? ", c.ID, c.RealName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysPostNameAlreadyExists")
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

func (s *SysPostService) Delete(id []int) error {
	if len(id) <= 0 {
		return errors.New("DeleteDataIsNull")
	}
	var m model.SysPost
	for _, v := range id {
		var u model.SysUser
		if !errors.Is(global.EGVA_DB.Model(&u).Where("post_id = ? ", v).First(&u).Error, gorm.ErrRecordNotFound) {
			return errors.New(fmt.Sprintf("SysPostIsInUseByUser_%v", v))
		}
	}
	if err := global.EGVA_DB.Model(&m).Delete(&m, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysPostService) List(p *dto.PageReq) (*vo.PageDataVo, error) {
	var m model.SysPost
	var ms []model.SysPost
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
