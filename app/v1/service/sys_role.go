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

type SysRoleService struct{}

func (s *SysRoleService) generateData(r *model.SysRole, data map[string]interface{}) error {
	data["real_name"] = r.RealName
	data["code"] = r.Code
	data["status"] = r.Status
	data["remark"] = r.Remark
	data["sort"] = r.Sort
	data["create_user"] = r.CreateUser
	return nil
}

func (s *SysRoleService) Create(r *model.SysRole) error {
	var m model.SysRole
	if !errors.Is(global.EGVA_DB.Model(&m).Where("real_name = ?", r.RealName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysRoleNameAlreadyExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("code = ?", r.Code).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysRoleCodeAlreadyExists")
	}
	data := map[string]interface{}{}
	data["created_at"] = time.Now()
	if err := s.generateData(r, data); err != nil {
		return err
	}
	if err := global.EGVA_DB.Model(&m).Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysRoleService) Update(r *model.SysRole) error {
	var m model.SysRole
	if errors.Is(global.EGVA_DB.Model(&m).First(&m, r.ID).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysRoleInfoDoesNotExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("id != ? and real_name = ? ", r.ID, r.RealName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysRoleNameAlreadyExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("id != ? and code = ?", r.ID, r.Code).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysRoleCodeAlreadyExists")
	}
	data := map[string]interface{}{}
	if err := s.generateData(r, data); err != nil {
		return err
	}
	if err := global.EGVA_DB.Model(&m).Where("id = ?", r.ID).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysRoleService) Delete(id []int) error {
	var m model.SysRole
	for _, v := range id {
		var us model.SysUserRole
		if errors.Is(global.EGVA_DB.Model(&us).Where("sys_role_id =? ", v).First(&us).Error, gorm.ErrRecordNotFound) {
			return errors.New(fmt.Sprintf("SysRoleIsInUseByUser_%v", v))
		}
	}
	if err := global.EGVA_DB.Model(&m).Delete(&m, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysRoleService) AssignAuthority(r *dto.SysRolePermissionReq) error {
	var m model.SysRole
	if errors.Is(global.EGVA_DB.Model(&m).First(&m, r.RoleID).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysRoleInfoDoesNotExists")
	}
	//开启锁机制
	mutexes := fmt.Sprintf("assign-authority-role-mutex-%s", utils.GenerateCode(10))
	mutex := global.EGVA_REDISSYNC.NewMutex(mutexes)
	if err := mutex.Lock(); err != nil {
		return err
	}
	m.ID = r.RoleID
	if len(r.PermissionList) > 0 {
		var replacePermission []model.SysPermission
		for _, v := range r.PermissionList {
			var p model.SysPermission
			p.ID = int64(v)
			replacePermission = append(replacePermission, p)
		}
		if err := global.EGVA_DB.Model(&m).Omit("Permission.*").Association("Permission").Replace(replacePermission); err != nil {
			return err
		}
	} else {
		if err := global.EGVA_DB.Model(&m).Association("Permission").Clear(); err != nil {
			return err
		}
	}
	//关闭锁机制
	if ok, err := mutex.Unlock(); !ok || err != nil {
		return err
	}
	return nil
}

func (s *SysRoleService) GetPermission(ids []int) ([]*model.SysRole, error) {
	var m model.SysRole
	var ms []*model.SysRole
	if err := global.EGVA_DB.Model(&m).Find(&ms, ids).Error; err != nil {
		return nil, err
	}
	return ms, nil
}

func (s *SysRoleService) List(p *dto.PageReq) (*vo.PageDataVo, error) {
	var m model.SysRole
	var ms []model.SysRole
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
