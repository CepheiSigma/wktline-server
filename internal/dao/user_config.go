// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"wktline-server/internal/dao/internal"
)

// internalUserConfigDao is internal type for wrapping internal DAO implements.
type internalUserConfigDao = *internal.UserConfigDao

// userConfigDao is the data access object for table user_config.
// You can define custom methods on it to extend its functionality as you wish.
type userConfigDao struct {
	internalUserConfigDao
}

var (
	// UserConfig is globally public accessible object for table user_config operations.
	UserConfig = userConfigDao{
		internal.NewUserConfigDao(),
	}
)

// Fill with you ideas below.
