/*
* @desc:部门model
* @company:舒可到家健康管理有限公司
* @Author: pmf<291641454@qq.com>
* @Date:   2022/4/11 9:07
 */

package model

import "github.com/tiger1103/gfast/v3/internal/app/system/model/entity"

type SysDeptTreeRes struct {
	*entity.SysDept
	Children []*SysDeptTreeRes `json:"children"`
}
