package models

import (
	"fmt"
	"strings"
)

func CreateGroup(grp *GroupProfile) (*GroupProfile, error) {
	tpl, validErr := validateNewGroup(grp)
	if validErr != nil {
		fmt.Println(validErr)
		return nil, validErr
	}
	if err := db.Model(&Group{}).Create(tpl).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	// fmt.Println(tpl)
	return &tpl.GroupProfile, nil
}

func GetGroup(id string) (*GroupProfile, error) {
	grp := &GroupProfile{}
	if err := db.Model(&Group{}).Where("id = ?", id).First(grp).Error; err != nil {
		return nil, err
	}
	return grp, nil
}

func GetFavoriteGroup(mobile string) ([]GroupProfile, error) {
	fav_grps := []GroupProfile{}
	usr := &User{}
	if err1 := db.Model(&User{}).Where("mobile_number = ?", mobile).First(usr).Error; err1 != nil {
		return fav_grps, err1
	}
	db.Model(&usr).Association("FavoriteGroups").Find(&fav_grps)
	return fav_grps, nil
}

func ListGroups() ([]GroupProfile, error) {
	all_grps := []GroupProfile{}
	if err1 := db.Model(&Group{}).Find(&all_grps).Error; err1 != nil {
		return all_grps, err1
	}
	return all_grps, nil
}

func AddGroupToFavorite(id string, usr_mobile string) error {
	grp := &Group{}
	usr_query := &User{}
	if err1 := db.Model(&User{}).Where("mobile_number = ?", usr_mobile).First(usr_query).Error; err1 != nil {
		return err1
	}
	if err2 := db.Model(&Group{}).Where("id = ?", id).First(grp).Error; err2 != nil {
		return err2
	}
	err3 := db.Model(&usr_query).Association("FavoriteGroups").Append(grp)
	if err3 != nil {
		fmt.Printf("Error while associating FavoriteGroups to user, Reason: %v\n", err3)
	}
	return nil
}

func FindGroups(search *GroupSearch) ([]GroupProfile, error) {
	var grps = []GroupProfile{}
	// var query_string string
	var query_string strings.Builder
	var enable_and_before_diety bool
	var enable_and_before_location bool
	fmt.Println(search)
	if len(search.Name) > 0 {
		query_string.WriteString("name LIKE '%" + search.Name + "%'")
		enable_and_before_diety, enable_and_before_location = true, true
	}
	if len(search.Deity) > 0 {
		if enable_and_before_diety {
			query_string.WriteString(" AND ")
		}
		query_string.WriteString("core_deity LIKE '%" + search.Deity + "%'")
		enable_and_before_location = true
	}
	if len(search.Location) > 0 {
		if enable_and_before_location {
			query_string.WriteString(" AND ")
		}
		query_string.WriteString("location LIKE '%" + search.Location + "%'")
	}
	fmt.Println(query_string.String())
	if err := db.Model(&Group{}).Where(query_string.String()).Find(&grps).Error; err != nil {
		return nil, err
	}
	return grps, nil
}

func UpdateGroupProfile(grp *GroupProfile, grp_id string) error {
	tpl, validErr := validategroup(grp)
	if validErr != nil {
		return validErr
	}
	if err := db.Model(&Group{}).Where("g_id = ?", grp_id).Updates(tpl).Error; err != nil {
		return err
	}
	return nil
}

func validateNewGroup(grp *GroupProfile) (*Group, error) {
	new_grp := &Group{GroupProfile: *grp}
	return new_grp, nil
}

func validategroup(grp *GroupProfile) (interface{}, error) {
	rt_grp := map[string]interface{}{}
	if len(grp.Pincode) > 0 && len(grp.Pincode) == 6 && grp.Pincode > "110000" && grp.Pincode < "850000" {
		rt_grp["pincode"] = grp.Pincode
	}
	if len(grp.GroupContact) > 0 && len(grp.GroupContact) == 10 {
		rt_grp["group_contact"] = grp.GroupContact
	}
	if len(grp.GID) > 0 {
		rt_grp["g_id"] = grp.GID
	}
	if len(grp.Bio) > 0 {
		rt_grp["bio"] = grp.Bio
	}
	if len(grp.Name) > 0 {
		rt_grp["name"] = grp.Name
	}
	if len(grp.GroupAvatar) > 0 {
		rt_grp["group_avatar"] = grp.GroupAvatar
	}
	fmt.Println(rt_grp)
	return rt_grp, nil
}
