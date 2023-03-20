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

type SysRankService struct{}

func (s *SysRankService) generateData(c *model.SysRank, data map[string]interface{}) error {
	data["real_name"] = c.RealName
	data["sort"] = c.Sort
	data["status"] = c.Status
	data["remark"] = c.Remark
	data["create_user"] = c.CreateUser
	return nil
}

func (s *SysRankService) Create(c *model.SysRank) error {
	var m model.SysRank
	if !errors.Is(global.EGVA_DB.Model(&m).Where("real_name = ?", c.RealName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysRankNameAlreadyExists")
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

func (s *SysRankService) Update(c *model.SysRank) error {
	var m model.SysRank
	if errors.Is(global.EGVA_DB.Model(&m).First(&m, c.ID).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysRankInfoDoesNotExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("id != ? and real_name = ? ", c.ID, c.RealName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysRankNameAlreadyExists")
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

func (s *SysRankService) Delete(id []int) error {
	if len(id) <= 0 {
		return errors.New("DeleteDataIsNull")
	}
	for _, v := range id {
		var u model.SysUser
		if !errors.Is(global.EGVA_DB.Model(&u).Where("rank_id = ? ", v).First(&u).Error, gorm.ErrRecordNotFound) {
			return errors.New(fmt.Sprintf("SysRankIsInUseByUser_%v", v))
		}
	}
	var m model.SysRank
	if err := global.EGVA_DB.Model(&m).Delete(&m, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysRankService) List(p *dto.PageReq) (*vo.PageDataVo, error) {
	var m model.SysRank
	var ms []model.SysRank
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
