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

type SysDepartmentService struct{}

func (s *SysDepartmentService) generateData(d *model.SysDepartment, data map[string]interface{}) error {
	if d.ParentID > 0 {
		var m model.SysDepartment
		if errors.Is(global.EGVA_DB.Model(&m).First(&m, d.ParentID).Error, gorm.ErrRecordNotFound) {
			return errors.New("SysDepartmentToParentClassDoesNotExists")
		}
	}
	data["parent_id"] = d.ParentID
	data["real_name"] = d.RealName
	data["full_name"] = d.FullName
	data["is_type"] = d.IsType
	data["remark"] = d.Remark
	data["sort"] = d.Sort
	data["status"] = d.Status
	data["create_user"] = d.CreateUser
	return nil
}

func (s *SysDepartmentService) Create(d *model.SysDepartment) error {
	var m model.SysDepartment
	if !errors.Is(global.EGVA_DB.Model(&m).Where("real_name = ?", d.RealName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysDepartmentNameAlreadyExists")
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

func (s *SysDepartmentService) Update(d *model.SysDepartment) error {
	var m model.SysDepartment
	if errors.Is(global.EGVA_DB.Model(&m).First(&m, d.ID).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysDepartmentInfoDoesNotExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("id != ? and real_name = ? ", d.ID, d.RealName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysDepartmentNameAlreadyExists")
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

func (s *SysDepartmentService) Delete(id []int) error {
	if len(id) <= 0 {
		return errors.New("DeleteDataIsNull")
	}
	for _, v := range id {
		var sm model.SysDepartment
		if !errors.Is(global.EGVA_DB.Model(&sm).Where("parent_id = ? ", v).Error, gorm.ErrRecordNotFound) {
			return errors.New(fmt.Sprintf("SysDepartmentIsInUseBySubClasses_%v", v))
		}
	}
	var m model.SysDepartment
	if err := global.EGVA_DB.Model(&m).Delete(&m, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysDepartmentService) Join(d *dto.SysUserJoinDepartmentReq) error {
	var m model.SysDepartment
	if errors.Is(global.EGVA_DB.Model(&m).First(&m, d.DepartmentID).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysDepartmentInfoDoesNotExists")
	}
	//开启锁机制
	mutexes := fmt.Sprintf("join-departmen-mutex-%s", utils.GenerateCode(10))
	mutex := global.EGVA_REDISSYNC.NewMutex(mutexes)
	if err := mutex.Lock(); err != nil {
		return err
	}
	localDB := global.EGVA_DB.Begin()
	m.ID = d.DepartmentID
	if len(d.UserList) > 0 {
		var replaceUser []model.SysUser
		for _, v := range d.UserList {
			var u model.SysUser
			u.ID = int64(v)
			replaceUser = append(replaceUser, u)
		}
		if err := localDB.Model(&m).Omit("UserList.*").Association("UserList").Replace(replaceUser); err != nil {
			localDB.Rollback()
			return err
		}
	} else {
		if err := localDB.Model(&m).Association("UserList").Clear(); err != nil {
			localDB.Rollback()
			return err
		}
	}
	//关闭锁机制
	if ok, err := mutex.Unlock(); !ok || err != nil {
		return err
	}
	localDB.Commit()
	return nil
}

func (s *SysDepartmentService) List(p *dto.PageReq) (*vo.PageDataVo, error) {
	var m model.SysDepartment
	var ms []model.SysDepartment
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
