// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"wktline-server/internal/dao/internal"
)

// internalTaskStatusDao is internal type for wrapping internal DAO implements.
type internalTaskStatusDao = *internal.TaskStatusDao

// taskStatusDao is the data access object for table task_status.
// You can define custom methods on it to extend its functionality as you wish.
type taskStatusDao struct {
	internalTaskStatusDao
}

var (
	// TaskStatus is globally public accessible object for table task_status operations.
	TaskStatus = taskStatusDao{
		internal.NewTaskStatusDao(),
	}
)

// Fill with you ideas below.
