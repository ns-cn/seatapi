package seatapi

import (
	"fmt"
	"net/http"
	"time"
)

// GetBaseInfo 获取基础信息
func (api SeaTableApi) GetBaseInfo(context BaseContext) SeaTableAction[BaseInfo] {
	url := api.wholeUrl(fmt.Sprintf("/dtable-server/dtables/68fa7023-aa53-467d-b000-b2edbe844a34"))
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE3MDU2NDUzMDUsImR0YWJsZV91dWlkIjoiNjhmYTcwMjMtYWE1My00NjdkLWIwMDAtYjJlZGJlODQ0YTM0IiwidXNlcm5hbWUiOiIiLCJwZXJtaXNzaW9uIjoicnciLCJhcHBfbmFtZSI6ImFwaVx1NmQ0Ylx1OGJkNSJ9.O0ymWG_wsr0SX-o3U9monFK-zl14C4Or20CwjPWbpBc")
	action := SeaTableAction[BaseInfo]{}
	action.Request = req
	return action
}

// BaseInfo Base基础信息
type BaseInfo struct {
	Version       int           `json:"version"`
	FormatVersion int           `json:"format_version"`
	Statistics    []interface{} `json:"statistics"`
	Links         []interface{} `json:"links"`
	Tables        []struct {
		Id   string `json:"_id"`
		Name string `json:"name"`
		Rows []struct {
			Id           string `json:"_id"`
			Participants []struct {
				AvatarUrl    string `json:"avatar_url"`
				ContactEmail string `json:"contact_email"`
				Email        string `json:"email"`
				Name         string `json:"name"`
				NamePinyin   string `json:"name_pinyin"`
				Id           string `json:"id"`
			} `json:"_participants,omitempty"`
			Creator      string    `json:"_creator,omitempty"`
			Ctime        time.Time `json:"_ctime"`
			LastModifier string    `json:"_last_modifier,omitempty"`
			Mtime        time.Time `json:"_mtime"`
			Field7       string    `json:"0000,omitempty"`
			XH2J         string    `json:"xH2J,omitempty"`
			MHcY         string    `json:"MHcY,omitempty"`
		} `json:"rows"`
		Columns []struct {
			Key       string `json:"key"`
			Name      string `json:"name"`
			Type      string `json:"type"`
			Width     int    `json:"width"`
			Editable  bool   `json:"editable"`
			Resizable bool   `json:"resizable"`
			Draggable bool   `json:"draggable,omitempty"`
			Data      *struct {
				EnableFillDefaultValue bool        `json:"enable_fill_default_value,omitempty"`
				DefaultValue           interface{} `json:"default_value"`
				Options                []struct {
					Name      string `json:"name"`
					Color     string `json:"color"`
					TextColor string `json:"textColor"`
					Id        string `json:"id"`
				} `json:"options"`
			} `json:"data,omitempty"`
			PermissionType             string        `json:"permission_type,omitempty"`
			PermittedUsers             []interface{} `json:"permitted_users,omitempty"`
			EditMetadataPermissionType string        `json:"edit_metadata_permission_type,omitempty"`
			EditMetadataPermittedUsers []interface{} `json:"edit_metadata_permitted_users,omitempty"`
			Description                interface{}   `json:"description"`
		} `json:"columns"`
		ViewStructure struct {
			Folders []interface{} `json:"folders"`
			ViewIds []string      `json:"view_ids"`
		} `json:"view_structure"`
		Views []struct {
			Id          string        `json:"_id"`
			Name        string        `json:"name"`
			Type        string        `json:"type"`
			IsLocked    bool          `json:"is_locked"`
			Rows        []interface{} `json:"rows"`
			FormulaRows struct {
			} `json:"formula_rows"`
			Summaries struct {
			} `json:"summaries"`
			FilterConjunction string        `json:"filter_conjunction"`
			Filters           []interface{} `json:"filters"`
			Sorts             []interface{} `json:"sorts"`
			HiddenColumns     []interface{} `json:"hidden_columns"`
			GroupBys          []interface{} `json:"groupbys"`
			Groups            []interface{} `json:"groups"`
			Colors            struct {
			} `json:"colors"`
			ColumnColors struct {
			} `json:"column_colors"`
			LinkRows struct {
			} `json:"link_rows"`
			PrivateFor interface{} `json:"private_for"`
			RowHeight  string      `json:"row_height,omitempty"`
			Colorbys   struct {
			} `json:"colorbys,omitempty"`
		} `json:"views"`
		IdRowMap struct {
		} `json:"id_row_map"`
		SummaryConfigs struct {
		} `json:"summary_configs,omitempty"`
		IsHeaderLocked bool `json:"is_header_locked,omitempty"`
		HeaderSettings struct {
		} `json:"header_settings,omitempty"`
	} `json:"tables"`
	Collaborators []struct {
		Email        string `json:"email"`
		Name         string `json:"name"`
		AvatarUrl    string `json:"avatar_url"`
		ContactEmail string `json:"contact_email"`
		IdInOrg      string `json:"id_in_org"`
		NamePinyin   string `json:"name_pinyin"`
	} `json:"collaborators"`
}
