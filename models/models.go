package models

import "encoding/json"

type Response struct {
	// Базовые поля ответа
	Success bool   `json:"success"`
	Errors  string `json:"errors,omitempty"`

	// Поля для get_balance
	Balance      float64 `json:"balance,omitempty"`
	BlockedMoney float64 `json:"blocked_money,omitempty"`

	// Поля для get_folders и create_folder
	Folders []struct {
		ID   json.Number `json:"id"`
		Name string      `json:"name"`
	} `json:"folders"`

	// Поля для get_tasks
	Tasks []struct {
		ID         json.Number `json:"id"`
		Name       string      `json:"name"`
		PriceRub   float64     `json:"price_rub"`
		TarifID    json.Number `json:"tarif_id"`
		Status     json.Number `json:"status"`
		FolderID   json.Number `json:"folder_id"`
		LimitTotal json.Number `json:"limit_total"`
	} `json:"tasks,omitempty"`

	// Поля для get_reports
	Reports []struct {
		ID       json.Number `json:"id"`
		TaskID   json.Number `json:"task_id"`
		WorkerID json.Number `json:"worker_id"`
		PriceRub float64     `json:"price_rub"`
		Status   json.Number `json:"status"`
		IP       string      `json:"IP"`
		Messages []struct {
			FromID json.Number `json:"from_id"`
			ToID   json.Number `json:"to_id"`
			Date   string      `json:"date"`
			Text   string      `json:"text"`
		} `json:"messages"`
		Files []string `json:"files"`
	} `json:"reports,omitempty"`

	// Поля для get_expenses
	Expenses      float64 `json:"expenses,omitempty"`
	ExpensesInRub float64 `json:"expenses_in_rub,omitempty"`
	GroupByDays   []struct {
		Date          string  `json:"date"`
		Expenses      float64 `json:"expenses"`
		ExpensesInRub float64 `json:"expenses_in_rub"`
	} `json:"group_by_days,omitempty"`

	// Поля для add_task

	// Поля для get_tariffs
	Tariffs []struct {
		ID          json.Number `json:"id"`
		Name        string      `json:"name"`
		MinPriceRub float64     `json:"min_price_rub"`
		GroupID     json.Number `json:"group_id"`
	} `json:"tariffs,omitempty"`

	// Поля для get_countries
	Countries []struct {
		ID   json.Number `json:"id"`
		Name string      `json:"name"`
	} `json:"countries,omitempty"`

	// Поля для get_blacklist
	Users []json.Number `json:"users,omitempty"`

	// Входные параметры для всех методов
	Name                  string      `json:"name,omitempty"`
	FolderID              json.Number `json:"folder_id,omitempty"`
	TaskID                json.Number `json:"task_id,omitempty"`
	Status                json.Number `json:"status,omitempty"`
	Offset                json.Number `json:"offset,omitempty"`
	ReportID              json.Number `json:"report_id,omitempty"`
	Comment               string      `json:"comment,omitempty"`
	RejectType            json.Number `json:"reject_type,omitempty"`
	DateFrom              string      `json:"date_from,omitempty"`
	DateTo                string      `json:"date_to,omitempty"`
	Descr                 string      `json:"descr,omitempty"`
	Link                  string      `json:"link,omitempty"`
	NeedForReport         string      `json:"need_for_report,omitempty"`
	Price                 float64     `json:"price,omitempty"`
	TarifID               json.Number `json:"tarif_id,omitempty"`
	NeedScreen            bool        `json:"need_screen,omitempty"`
	AnonymTask            bool        `json:"anonym_task,omitempty"`
	TimeForWork           json.Number `json:"time_for_work,omitempty"`
	TimeForCheck          json.Number `json:"time_for_check,omitempty"`
	LimitPerDay           json.Number `json:"limit_per_day,omitempty"`
	LimitPerHour          json.Number `json:"limit_per_hour,omitempty"`
	LimitPerUser          json.Number `json:"limit_per_user,omitempty"`
	LimitPerUserFolder    json.Number `json:"limit_per_user_folder,omitempty"`
	LimitPerIP            json.Number `json:"limit_per_ip,omitempty"`
	LimitOnlyForLevelID   json.Number `json:"limit_only_for_level_id,omitempty"`
	LimitDateFrom         string      `json:"limit_date_from,omitempty"`
	LimitDateTo           string      `json:"limit_date_to,omitempty"`
	DelayFrom             json.Number `json:"delay_from,omitempty"`
	DelayTo               json.Number `json:"delay_to,omitempty"`
	TargetingGender       json.Number `json:"targeting_gender,omitempty"`
	TargetingAgeFrom      json.Number `json:"targeting_age_from,omitempty"`
	TargetingAgeTo        json.Number `json:"targeting_age_to,omitempty"`
	TargetingGeoCountryID json.Number `json:"targeting_geo_country_id,omitempty"`
	TargetingGeoRegionID  json.Number `json:"targeting_geo_region_id,omitempty"`
	TargetingGeoCityID    json.Number `json:"targeting_geo_city_id,omitempty"`
	TaskOnlyForListID     json.Number `json:"task_only_for_list_id,omitempty"`
	ListOfPages           string      `json:"list_of_pages,omitempty"`
	AddToLimit            json.Number `json:"add_to_limit,omitempty"`
	SubToLimit            json.Number `json:"sub_to_limit,omitempty"`
	AddBlacklistID        json.Number `json:"add_blacklist_id,omitempty"`
	AddWhitelistID        json.Number `json:"add_whitelist_id,omitempty"`
	IDUserBlacklist       json.Number `json:"id_user_blacklist,omitempty"`
}
