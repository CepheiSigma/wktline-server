// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"wktline-server/internal/dao/internal"
)

// internalTaskDao is internal type for wrapping internal DAO implements.
type internalTaskDao = *internal.TaskDao

// taskDao is the data access object for table task.
// You can define custom methods on it to extend its functionality as you wish.
type taskDao struct {
	internalTaskDao
}

var (
	// Task is globally public accessible object for table task operations.
	Task = taskDao{
		internal.NewTaskDao(),
	}
)

// Fill with you ideas below.
