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
	"time"
)

type SysUserService struct{}

func (s *SysUserService) generateData(u *model.SysUser, data map[string]interface{}) error {
	data["account_name"] = u.AccountName
	data["username"] = u.UserName
	data["nickname"] = u.NickName
	if u.PassWord != "" {
		salt := utils.GenerateCode(6)
		data["salt"] = salt
		data["password"] = utils.BcryptHash(fmt.Sprintf("%s%s", u.PassWord, salt))
	}
	fmt.Println("data[\"password\"]", data["password"])
	data["gender"] = u.Gender
	data["birthday"] = u.Birthday
	data["email"] = u.Email
	data["mobile"] = u.Mobile
	data["ip"] = u.Ip
	data["avatar"] = u.Avatar
	data["city"] = u.City
	data["address"] = u.Address
	data["status"] = u.Status
	data["remark"] = u.Remark
	data["create_user"] = u.CreateUser
	return nil
}

func (s *SysUserService) Create(u *model.SysUser) error {
	var m model.SysUser
	if !errors.Is(global.EGVA_DB.Model(&m).Where("account_name = ?", u.AccountName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysUserAccountNameAlreadyExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("username = ?", u.UserName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysUserNameAlreadyExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("mobile = ?", u.Mobile).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysUserMobileAlreadyExists")
	}
	data := map[string]interface{}{}
	data["created_at"] = time.Now()
	if err := s.generateData(u, data); err != nil {
		return err
	}
	//开启锁机制
	mutexes := fmt.Sprintf("create-user-mutex-%s", utils.GenerateCode(10))
	mutex := global.EGVA_REDISSYNC.NewMutex(mutexes)
	if err := mutex.Lock(); err != nil {
		return err
	}
	localDB := global.EGVA_DB.Begin()
	if err := localDB.Model(&m).Create(&data).Error; err != nil {
		localDB.Rollback()
		return err
	}
	m.ID = u.ID
	if len(u.RoleList) > 0 {
		var replaceRole []model.SysRole
		for _, v := range u.RoleList {
			var r model.SysRole
			r.ID = v.ID
			replaceRole = append(replaceRole, r)
		}
		if err := localDB.Model(&m).Omit("RoleList.*").Association("RoleList").Replace(replaceRole); err != nil {
			localDB.Rollback()
			return err
		}
	} else {
		if err := localDB.Model(&m).Association("RoleList").Clear(); err != nil {
			localDB.Rollback()
			return err
		}
	}
	//关闭锁机制
	if ok, err := mutex.Unlock(); !ok || err != nil {
		localDB.Rollback()
		return err
	}
	localDB.Commit()
	return nil
}

func (s *SysUserService) Update(u *model.SysUser) error {
	var m model.SysUser
	if errors.Is(global.EGVA_DB.Model(&m).First(&m, u.ID).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysUserInfoDoesNotExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("account_name = ?", u.AccountName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysUserAccountNameAlreadyExists")
	}
	if !errors.Is(global.EGVA_DB.Model(&m).Where("username = ?", u.UserName).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysUserNameAlreadyExists")
	}
	data := map[string]interface{}{}
	if err := s.generateData(u, data); err != nil {
		return err
	}
	//开启锁机制
	mutexes := fmt.Sprintf("update-user-mutex-%s", utils.GenerateCode(10))
	mutex := global.EGVA_REDISSYNC.NewMutex(mutexes)
	if err := mutex.Lock(); err != nil {
		return err
	}
	localDB := global.EGVA_DB.Begin()
	if err := localDB.Model(&m).Updates(&data).Error; err != nil {
		localDB.Rollback()
		return err
	}
	m.ID = u.ID
	if len(u.RoleList) > 0 {
		var replaceRole []model.SysRole
		for _, v := range u.RoleList {
			var r model.SysRole
			r.ID = v.ID
			replaceRole = append(replaceRole, r)
		}
		if err := localDB.Model(&m).Omit("RoleList.*").Association("RoleList").Replace(replaceRole); err != nil {
			localDB.Rollback()
			return err
		}
	} else {
		if err := localDB.Model(&m).Association("RoleList").Clear(); err != nil {
			localDB.Rollback()
			return err
		}
	}
	//关闭锁机制
	if ok, err := mutex.Unlock(); !ok || err != nil {
		localDB.Rollback()
		return err
	}
	localDB.Commit()
	return nil
}

func (s *SysUserService) AssignPostRank(u *model.SysUser) error {
	var m model.SysUser
	if u.PostID > 0 {
		var p model.SysPost
		if errors.Is(global.EGVA_DB.Model(&p).First(&p, u.PostID).Error, gorm.ErrRecordNotFound) {
			return errors.New("SysPostInfoDoesNotExists")
		}
	}
	if u.RankID > 0 {
		var r model.SysRank
		if errors.Is(global.EGVA_DB.Model(&r).First(&r, u.RankID).Error, gorm.ErrRecordNotFound) {
			return errors.New("SysRankInfoDoesNotExists")
		}
	}
	if err := global.EGVA_DB.Model(&m).Updates(u).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysUserService) Delete(id []int) error {
	var m model.SysUser
	if err := global.EGVA_DB.Model(&m).Delete(&m, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysUserService) Detail(u *model.SysUser) (*model.SysUser, error) {
	var m model.SysUser
	if errors.Is(global.EGVA_DB.Model(&m).Preload("PostInfo").Preload("RankInfo").First(&m, u.ID).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("SysUserInfoDoesNotExists")
	}
	return &m, nil
}

func (s *SysUserService) ResetPassWord(u *model.SysUser) error {
	var m model.SysUser
	if errors.Is(global.EGVA_DB.Model(&m).First(&m, u.ID).Error, gorm.ErrRecordNotFound) {
		return errors.New("SysUserInfoDoesNotExists")
	}
	data := map[string]interface{}{}
	salt := utils.GenerateCode(6)
	data["salt"] = salt
	data["password"] = utils.BcryptHash(fmt.Sprintf("%s%s", u.PassWord, salt))
	if err := global.EGVA_DB.Model(&m).Where("id = ?", u.ID).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysUserService) QueryName(u *model.SysUser) (*model.SysUser, error) {
	var m model.SysUser
	if errors.Is(global.EGVA_DB.Model(&m).Where("username = ? ", u.UserName).First(&m).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("SysUserInfoDoesNotExists")
	}
	return &m, nil
}

func (s *SysUserService) QueryMobile(u *model.SysUser) (*model.SysUser, error) {
	var m model.SysUser
	if errors.Is(global.EGVA_DB.Model(&m).Where("mobile = ? ", u.Mobile).First(&m).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("SysUserMobileDoesNotExists")
	}
	return &m, nil
}

func (s *SysUserService) JoinGroup(u *model.SysUser) {

}

func (s *SysUserService) List(p *dto.PageReq) (*vo.PageDataVo, error) {
	var m model.SysUser
	var ms []model.SysUser
	localDB := global.EGVA_DB.Model(&m)
	if p.Keyword != "" {
		localDB = localDB.Where("account_name LIKE ? OR username LIKE ? OR nickname LIKE ? OR mobile LIKE ? OR email LIKE ? OR mobile LIKE ?",
			p.Keyword+"%", p.Keyword+"%", p.Keyword+"%", p.Keyword+"%", p.Keyword+"%", p.Keyword+"%")
	}
	if p.Status > 0 {
		localDB = localDB.Where("status = ? ", p.Status)
	}
	var total int64
	global.EGVA_DB.Model(&m).Count(&total)
	localDB.Scopes(model.Paginate(p.Page, p.PageSize)).Preload("PostInfo").Preload("RankInfo").Find(&ms)
	return &vo.PageDataVo{
		Total:    total,
		Data:     ms,
		Page:     p.Page,
		PageSize: p.PageSize,
	}, nil
}
