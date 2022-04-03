package tests

import (
	models "meetup/models"
	"testing"
)

func TestCleanUp(t *testing.T) {
	var err error
	user := &models.User{}
	models.GetDB().Where("id = ?", usr_id).First(&user)
	if err = models.GetDB().Where("group_id = ? AND user_id = ?", grp_id, usr_id).Delete(&models.Story{}).Error; err != nil {
		t.Error(err)
	}
	models.GetDB().Model(&user).Association("RSVPEvents").Clear()
	if err = models.GetDB().Where("group_id = ?", grp_id).Delete(&models.GroupEvent{}).Error; err != nil {
		t.Error(err)
	}
	models.GetDB().Model(&user).Association("FavoriteGroups").Clear()
	if err = models.GetDB().Where("group_id = ?", grp_id).Delete(&models.GroupEvent{}).Error; err != nil {
		t.Error(err)
	}
	if err = models.GetDB().Where("mobile_number = ?", testNumber).Delete(&models.User{}).Error; err != nil {
		t.Error(err)
	}
	if err = models.GetDB().Where("group_contact = ?", grp_group_contact).Delete(&models.Group{}).Error; err != nil {
		t.Error(err)
	}
}
