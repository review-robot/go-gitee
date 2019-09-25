/*
 * 码云 Open API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.3.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gitee

// 获取某个Pull Request的操作日志
type OperateLog struct {
	Id        string `json:"id,omitempty"`
	Icon      string `json:"icon,omitempty"`
	User      string `json:"user,omitempty"`
	Content   string `json:"content,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
