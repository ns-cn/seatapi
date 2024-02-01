package seatapi

import "time"

// RowRecord 行记录
type RowRecord struct {
	Id    string    `json:"_id"`
	Ctime time.Time `json:"_ctime"`
	Mtime time.Time `json:"_mtime"`
}

// BaseInfo 基础信息
type BaseInfo struct {
	Version       int                `json:"version"`
	FormatVersion int                `json:"format_version"`
	Statistics    []interface{}      `json:"statistics"`
	Links         []interface{}      `json:"links"`
	Tables        []BaseTable        `json:"tables"`
	Collaborators []BaseCollaborator `json:"collaborators"`
}

// BaseTableRow 行信息数据
type BaseTableRow struct {
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
}

// BaseTableColumn 列信息数据
type BaseTableColumn struct {
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
}

// BaseTableView 表视图信息数据
type BaseTableView struct {
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
}

// BaseTable 表信息数据
type BaseTable struct {
	Id            string            `json:"_id"`
	Name          string            `json:"name"`
	Rows          []BaseTableRow    `json:"rows"`
	Columns       []BaseTableColumn `json:"columns"`
	ViewStructure struct {
		Folders []interface{} `json:"folders"`
		ViewIds []string      `json:"view_ids"`
	} `json:"view_structure"`
	Views    []BaseTableView `json:"views"`
	IdRowMap struct {
	} `json:"id_row_map"`
	SummaryConfigs struct {
	} `json:"summary_configs,omitempty"`
	IsHeaderLocked bool `json:"is_header_locked,omitempty"`
	HeaderSettings struct {
	} `json:"header_settings,omitempty"`
}

// BaseCollaborator 用户信息数据
type BaseCollaborator struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	AvatarUrl    string `json:"avatar_url"`
	ContactEmail string `json:"contact_email"`
	IdInOrg      string `json:"id_in_org"`
	NamePinyin   string `json:"name_pinyin"`
}
