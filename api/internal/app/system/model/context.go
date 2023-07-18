/*
* @desc:context-model
* @company:舒可到家健康管理有限公司
* @Author: pmf
* @Date:   2022/3/16 14:45
 */

package model

type Context struct {
	User *ContextUser // User in context.
}

type ContextUser struct {
	*LoginUserRes
}
