package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/shakirovformal/unu_api/models"
)

// Method: get_balance // Возвращает количество доступных средств.
// Входные данные - отсутствуют
// Выходные данные:
// balance (float) – количество средств на балансе в UNU
// freeze (float) – количество замороженных средств текущих задач
func (c *Client) Get_balance(ctx context.Context) (*models.Response, error) {
	var resp *models.Response
	bytesRes := c.post(ctx, "get_balance", nil)

	err := json.Unmarshal([]byte(bytesRes), &resp)
	if err != nil {
		err = fmt.Errorf("ошибка парсинга JSON: %v", err)
		return nil, err
	}

	return resp, nil
}

// Method: get_folders // Возвращает все созданные папки с задачами.
// Входные данные - отсутствуют
// Выходные данные:
// folders (array) – массив с папками. Каждый элемент массива содержит:
// id (int) – уникальный идентификатор папки
// name (text) – имя папки
func (c *Client) Get_folders(ctx context.Context) (*models.Response, error) {
	var resp *models.Response
	bytesRes := c.post(ctx, "get_folders", nil)

	err := json.Unmarshal([]byte(bytesRes), &resp)
	if err != nil {
		err = fmt.Errorf("ошибка парсинга JSON: %v", err)
		return nil, err
	}
	return resp, nil
}

// Method: create_folder // Создаёт новую папку.
// Входные данные:
// name (text) – имя папки
// Выходные данные:
// folder_id (int) – уникальный идентификатор созданной папки
func (c *Client) Create_folder(ctx context.Context, folder_name string) (*models.Response, error) {
	var resp *models.Response
	params := make(map[string]interface{})
	params["name"] = folder_name
	bytesRes := c.post(ctx, "create_folder", params)

	err := json.Unmarshal([]byte(bytesRes), &resp)
	if err != nil {
		err = fmt.Errorf("ошибка парсинга JSON: %v", err)
		return nil, err
	}

	return resp, nil

}

// Method: del_folder // Удаляет папку.
// Входные данные:
// folder_id (id) – идентификатор папки, которую нужно удалить
// Выходные данные - отсутствуют
func (c *Client) Del_folder(ctx context.Context, folder_id int) (*models.Response, error) {
	var resp *models.Response
	params := make(map[string]interface{})
	params["folder_id"] = folder_id
	bytesRes := c.post(ctx, "del_folder", params)
	fmt.Println(bytesRes)

	err := json.Unmarshal([]byte(bytesRes), &resp)
	if err != nil {
		err = fmt.Errorf("ошибка парсинга JSON: %v", err)
		return nil, err
	}
	if !resp.Success {
		err = fmt.Errorf("ошибка парсинга JSON: %v", err)
		return nil, err
	}

	return resp, nil
}

// Method: move_task // Перемещает задачу в указанную папку.
// Входные данные:
// task_id (int) – идентификатор задачи
// folder_id (int) – идентификатор папки, куда нужно переместить задачу
// Выходные данные - отсутствуют
func (c *Client) Move_task(ctx context.Context, task_id, folder_id int) (*models.Response, error) {
	var resp *models.Response
	params := make(map[string]interface{})
	params["task_id"] = task_id
	params["folder_id"] = folder_id
	bytesRes := c.post(ctx, "move_task", params)

	err := json.Unmarshal([]byte(bytesRes), &resp)
	if err != nil {
		err = fmt.Errorf("ошибка парсинга JSON: %v", err)
		return nil, err
	}

	return resp, nil
}

// Method: get_tasks // Возвращает существующие задачи.
// Входные данные:
// folder_id (int) – идентификатор папки, из которой нужно показать задачи (необязательный параметр)
// status (int) – статус задачи, один или несколько статусов через запятую (необязательный параметр)
// task_id (int) – идентификатор задачи, один или несколько id через запятую (необязательный параметр)
// offset (int) – данный параметр устанавливает смещение для выборки, по-умолчанию метод возвращает не более 50 тыс. записей (необязательный параметр)
// Выходные данные
// id (int) – идентификатор задачи
// name (text) – название задачи
// price_rub (float) – стоимость выполнения в рублях
// tarif_id (int) – ID тарифа задачи
// status (int) – текущий статус задачи
//
//	1 – новое задание, нужно оплатить (увеличить лимит)
//	2 – достигло лимита
//	3 – остановлено
//	4 – активно
//	5 – отклонено модератором
//	6 – на модерации
//
// folder_id (int) – идентификатор папки
// limit_total (int) – количество заказанных выполнений
func (c *Client) Get_tasks(ctx context.Context, folder_id, status, task_id, offset int) (*models.Response, error) {
	var resp *models.Response

	params := map[string]interface{}{
		"folder_id": folder_id,
		"status":    status,
		"task_id":   task_id,
		"offset":    offset,
	}
	if status == 0 {
		delete(params, "status")
	}
	if task_id == 0 {
		delete(params, "task_id")
	}
	if offset == 0 {
		delete(params, "offset")
	}

	bytesRes := c.post(ctx, "get_tasks", params)
	if err := json.Unmarshal([]byte(bytesRes), &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// Method: get_reports //Возвращает отчёты по определённой задаче или все существующие отчёты.
// Входные данные:
// task_id (int) – идентификатор задачи, по которой нужно вернуть отчёты,
// один или несколько id через запятую (обязательный параметр)
// offset (int) – данный параметр устанавливает смещение для выборки,
// по-умолчанию метод возвращает не более 1000 записей (необязательный параметр)
// Выходные данные:
// id (int) – идентификатор отчёта
// task_id (int) – идентификатор отчёта
// worker_id (int) – идентификатор пользователя, выполняющего работу
// price_rub (float) – стоимость выполнения в рублях
// status (int) – текущий статус отчёта
//
//	1 – в работе
//	2 – на проверке
//	3 – на доработке
//	6 – оплачено
//
// IP (string) – IP адрес исполнителя
// messages (array) – массив сообщений по отчету
// from_id (int) – ID отправителя
// to_id (int) – ID получателя
// date – дата
// text – сообщение
// files (array) – массив ссылок на файлы
func (c *Client) Get_reports(ctx context.Context, task_id, offset int) (*models.Response, error) {
	var resp *models.Response
	params := map[string]interface{}{
		"task_id": task_id,
		"offset":  offset,
	}
	if offset == 0 {
		delete(params, "offset")
	}
	bytesRes := c.post(ctx, "get_reports", params)
	if err := json.Unmarshal([]byte(bytesRes), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Method: approve_report // принимает (оплачивает) отчёт по заданию.
// Входные данные:
// report_id (int) – идентификатор отчёта, который нужно одобрить
// Выходные данные - отсутствуют
func (c *Client) Approve_report(ctx context.Context, report_id int) (*models.Response, error) {
	var resp *models.Response
	params := map[string]interface{}{
		"report_id": report_id,
	}
	bytesRes := c.post(ctx, "approve_report", params)
	if err := json.Unmarshal([]byte(bytesRes), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Method: reject_report // Отклоняет отчёт по заданию.
// Входные данные
// report_id (int) – идентификатор отчёта, который нужно одобрить
// comment (text) – причина отказа
// reject_type (int) – какое именно действие нужно выполнить
// 1 – отправить на доработку
// 2 – отказать
// Выходные данные отсутствуют
func (c *Client) Reject_report(ctx context.Context, report_id int, comment string, reject_type int) (*models.Response, error) {
	var resp *models.Response
	params := map[string]interface{}{
		"report_id":   report_id,
		"comment":     comment,
		"reject_type": reject_type,
	}
	bytesRes := c.post(ctx, "reject_report", params)
	if err := json.Unmarshal([]byte(bytesRes), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Method: get_expenses // Возврашает сумму израсходованных средств
// Входные данные
// task_id (int) – идентификатор задачи по которой нужно получить расходы (необязательный параметр)
// folder_id (int) – идентификатор папки по которой нужно получить расходы (необязательный параметр)
// date_from (datetime) – начало периода (с какой даты нужно вернуть расходы), пример 2019-11-01 13:00:00 (необязательный параметр)
// date_to (datetime) – конец периода (по какую дату нужно вернуть расходы), пример 2019-11-05 13:00:00 (необязательный параметр)
// Выходные данные
// expenses (float) – сумма расходов в UNU
// expenses_in_rub (float) – сумма расходов в рублях
// group_by_days (array) – расходы, сгруппированные по дням
func (c *Client) Get_expenses(ctx context.Context, task_id int, folder_id int, date_from string, date_to string) (*models.Response, error) {
	var resp *models.Response
	params := map[string]interface{}{
		"task_id":   task_id,
		"folder_id": folder_id,
		"date_from": date_from,
		"date_to":   date_to,
	}
	if task_id == 0 {
		delete(params, "task_id")
	}
	if folder_id == 0 {
		delete(params, "folder_id")
	}
	if date_from == "" {
		delete(params, "date_from")
	}
	if date_to == "" {
		delete(params, "date_to")
	}
	bytesRes := c.post(ctx, "get_expenses", params)
	if err := json.Unmarshal([]byte(bytesRes), &resp); err != nil {
		return nil, err
	}
	return resp, nil

}

// Method: add_task // Создаёт новую задачу
// Входные данные
// name (text) – название задачи
// descr (text) – текст задания
// link (text) – URL, необходимый для выполнения задания (необязательный параметр)
// need_for_report (text) – что должен предоставить исполнитель для отчёта по задаче
// price (float) – стоимость одного выполнения задачи в рублях
// tarif_id (int) – идентификатор тарифа
// folder_id (int) – идентификатор папки, в которую нужно поместить задачу
// need_screen (boolean) – если в задании исполнителю нужно прикерпить скриншот, нужно передать 1 (необязательный параметр)
// anonym_task (boolean) – делает задание анонимным (необязательный параметр)
// time_for_work (int) – сколько часов дать исполнителю для работы, от 2 до 168 (необязательный параметр)
// time_for_check (int) – сколько часов вам нужно для проверки задания, от 10 до 168 (необязательный параметр)
// limit_per_day (int) – лимит выполнений в сутки (необязательный параметр)
// limit_per_hour (int) – лимит выполнений в час (необязательный параметр)
// limit_per_user (int) – лимит выполнений для одного исполнителя (необязательный параметр)
// limit_per_user_folder (int) – лимит от исполнителя на папку (необязательный параметр)
// limit_per_ip (int) – лимит выполнений с одного IP (необязательный параметр)
// limit_only_for_level_id (int) – минимально необходимый уровень исполнителя: 1 - не ниже базового, 2 – не ниже продвинутого, 3 – не ниже выского, 4 – не ниже профи (необязательный параметр)
// limit_date_from (datetime) – время старта задания
// limit_date_to (datetime) – время остановки задания
// delay_from (int) – задержка между выполнениями в минутах, от (необязательный параметр)
// delay_to (int) – задержка между выполнениями в минутах, до (необязательный параметр)
// targeting_gender (int) – параметр таргетинга: пол. 1 – женский, 2 – мужской (необязательный параметр)
// targeting_age_from (int) – параметр таргетинга: возраст от (необязательный параметр)
// targeting_age_to (int) – параметр таргетинга: возраст до (необязательный параметр)
// targeting_geo_country_id (int) – параметр геотаргетинга: ID страны (необязательный параметр)
// targeting_geo_region_id (int) – параметр геотаргетинга: ID региона (необязательный параметр)
// targeting_geo_city_id (int) – параметр геотаргетинга: ID города (необязательный параметр)
// task_only_for_list_id (int) – ID белого списка исполнителей, которому будет доступно задание (необязательный параметр)
// list_of_pages (text) – данные для равномерного распределения информации среди исполнителей (необязательный параметр)
// Выходные данные
// task_id (int) – идентификатор созданной задачи
func (c *Client) Add_task(ctx context.Context, name, descr, link, need_for_report string,
	price float64,
	tarif_id, folder_id int,
	need_screen, anonym_task bool,
	time_for_work, time_for_check, limit_per_day, limit_per_hour, limit_per_user int,
	limit_per_user_folder, limit_per_ip, limit_only_for_level_id int,
	limit_date_from, limit_date_to string,
	delay_from, delay_to int,
	targeting_gender, targeting_age_from, targeting_age_to, targeting_geo_country_id int,
	targeting_geo_region_id, targeting_geo_city_id int,
	task_only_for_list_id int,
	list_of_pages string) (*models.Response, error) {
	// TODO: реализовать метод
	var resp *models.Response
	params := map[string]interface{}{
		"name":                     name,
		"descr":                    descr,
		"link":                     link,
		"need_for_report":          need_for_report,
		"price":                    price,
		"tarif_id":                 tarif_id,
		"folder_id":                folder_id,
		"need_screen":              need_screen,
		"anonym_task":              anonym_task,
		"time_for_work":            time_for_work,
		"time_for_check":           time_for_check,
		"limit_per_day":            limit_per_day,
		"limit_per_hour":           limit_per_hour,
		"limit_per_user":           limit_per_user,
		"limit_per_user_folder":    limit_per_user_folder,
		"limit_per_ip":             limit_per_ip,
		"limit_only_for_level_id":  limit_only_for_level_id,
		"limit_date_from":          limit_date_from,
		"limit_date_to":            limit_date_to,
		"targeting_gender":         targeting_gender,
		"targeting_age_from":       targeting_age_from,
		"targeting_age_to":         targeting_age_to,
		"targeting_geo_country_id": targeting_geo_country_id,
		"targeting_geo_region_id":  targeting_geo_region_id,
		"targeting_geo_city_id":    targeting_geo_city_id,
		"task_only_for_list_id":    task_only_for_list_id,
		"list_of_pages":            list_of_pages,
	}
	// TODO: Найти решение как обойти этот длинный кусок кода
	if time_for_work == 0 {
		delete(params, "time_for_work")
	}
	if time_for_check == 0 {
		delete(params, "time_for_check")
	}
	if limit_per_day == 0 {
		delete(params, "limit_per_day")
	}
	if limit_per_hour == 0 {
		delete(params, "limit_per_hour")
	}
	if limit_per_user == 0 {
		delete(params, "limit_per_user")
	}
	if limit_per_user_folder == 0 {
		delete(params, "limit_per_user_folder")
	}
	if limit_per_ip == 0 {
		delete(params, "limit_per_ip")
	}
	if limit_only_for_level_id == 0 {
		delete(params, "limit_only_for_level_id")
	}
	if limit_date_from == "" {
		delete(params, "limit_date_from")
	}
	if limit_date_to == "" {
		delete(params, "limit_date_to")
	}
	if delay_from == 0 {
		delete(params, "delay_from")
	}
	if delay_to == 0 {
		delete(params, "delay_to")
	}
	if targeting_gender == 0 {
		delete(params, "targeting_gender")
	}
	if targeting_age_from == 0 {
		delete(params, "targeting_age_from")
	}
	if targeting_age_to == 0 {
		delete(params, "targeting_age_to")
	}
	if targeting_geo_country_id == 0 {
		delete(params, "targeting_geo_country_id")
	}
	if targeting_geo_region_id == 0 {
		delete(params, "targeting_geo_region_id")
	}
	if targeting_geo_city_id == 0 {
		delete(params, "targeting_geo_city_id")
	}
	if task_only_for_list_id == 0 {
		delete(params, "task_only_for_list_id")
	}
	if list_of_pages == "" {
		delete(params, "list_of_pages")
	}

	bytesRes := c.post(ctx, "add_task", params)
	if err := json.Unmarshal([]byte(bytesRes), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Method: task_limit_add // Устанавливает лимит (добавляет выполнения) определённой задачи.
// После создания любой задачи обязательно нужно задать лимит выполнений по ней.
// Входные данные:
// task_id (int) – идентификатор задачи
// add_to_limit (int) – сколько раз нужно выполнить задание
// Выходные данные отсутствуют
func (c *Client) Task_limit_add(ctx context.Context, task_id, add_to_limit int) (*models.Response, error) {
	var resp *models.Response
	params := map[string]interface{}{
		"task_id":      task_id,
		"add_to_limit": add_to_limit,
	}
	bytesRes := c.post(ctx, "task_limit_add", params)
	if err := json.Unmarshal([]byte(bytesRes), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Method: task_limit_sub // Устанавливает лимит (убирает выполнения) определённой задачи. Изменяет лимит выполнений по задаче.

// Входные данные
// task_id (int) – идентификатор задачи
// sub_to_limit (int) – сколько выполнений нужно убрать у задания
// Выходные данные отсутствуют
func (c *Client) Task_limit_sub(ctx context.Context, task_id, sub_to_limit int) (*models.Response, error) {
	var resp *models.Response
	params := map[string]interface{}{
		"task_id":      task_id,
		"add_to_limit": sub_to_limit,
	}
	bytesRes := c.post(ctx, "task_limit_sub", params)
	if err := json.Unmarshal([]byte(bytesRes), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Method: edit_task // Редактирует существующую задачу
// Входные данные:
// task_id (int) – идентификатор задачи для редактирования
// name (text) – название задачи
// descr (text) – текст задания
// need_for_report (text) – что должен предоставить исполнитель для отчёта по задаче
// price (float) – стоимость одного выполнения задачи в рублях
// tarif_id (int) – идентификатор тарифа
// folder_id (int) – идентификатор папки, в которую нужно поместить задачу
// need_screen (boolean) – если в задании исполнителю нужно прикерпить скриншот, нужно передать 1 (необязательный параметр)
// anonym_task (boolean) – делает задание анонимным (необязательный параметр)
// time_for_work (int) – сколько часов дать исполнителю для работы, от 2 до 168 (необязательный параметр)
// time_for_check (int) – сколько часов вам нужно для проверки задания, от 10 до 168 (необязательный параметр)
// limit_per_day (int) – лимит выполнений в сутки (необязательный параметр)
// limit_per_hour (int) – лимит выполнений в час (необязательный параметр)
// limit_per_user (int) – лимит выполнений для одного исполнителя (необязательный параметр)
// limit_per_user_folder (int) – лимит от исполнителя на папку (необязательный параметр)
// limit_per_ip (int) – лимит выполнений с одного IP (необязательный параметр)
// limit_date_from (datetime) – время старта задания
// limit_date_to (datetime) – время остановки задания
// limit_only_for_level_id (int) – минимально необходимый уровень исполнителя (1 - не ниже базового, 2 – не ниже продвинутого, 3 – не ниже выского, 4 – не ниже профи)
// delay_from (int) – задержка между выполнениями в минутах, от (необязательный параметр)
// delay_to (int) – задержка между выполнениями в минутах, до (необязательный параметр)
// targeting_gender (int) – параметр таргетинга: пол. 1 – женский, 2 – мужской (необязательный параметр)
// targeting_age_from (int) – параметр таргетинга: возраст от (необязательный параметр)
// targeting_age_to (int) – параметр таргетинга: возраст до (необязательный параметр)
// targeting_geo_country_id (int) – параметр геотаргетинга: ID страны (необязательный параметр)
// targeting_geo_region_id (int) – параметр геотаргетинга: ID региона (необязательный параметр)
// targeting_geo_city_id (int) – параметр геотаргетинга: ID города (необязательный параметр)
// list_of_pages (text) – данные для равномерного распределения информации среди исполнителей (необязательный параметр)
// Выходные данные отсутствуют
func (c *Client) Edit_task(ctx context.Context, name, descr, link, need_for_report string,
	price float64,
	tarif_id, folder_id int,
	need_screen, anonym_task bool,
	time_for_work, time_for_check, limit_per_day, limit_per_hour, limit_per_user int,
	limit_per_user_folder, limit_per_ip, limit_only_for_level_id int,
	limit_date_from, limit_date_to string,
	delay_from, delay_to int,
	targeting_gender, targeting_age_from, targeting_age_to, targeting_geo_country_id int,
	targeting_geo_region_id, targeting_geo_city_id int,
	task_only_for_list_id int,
	list_of_pages string) (*models.Response, error) {
	var resp *models.Response
	params := map[string]interface{}{
		"name":                     name,
		"descr":                    descr,
		"link":                     link,
		"need_for_report":          need_for_report,
		"price":                    price,
		"tarif_id":                 tarif_id,
		"folder_id":                folder_id,
		"need_screen":              need_screen,
		"anonym_task":              anonym_task,
		"time_for_work":            time_for_work,
		"time_for_check":           time_for_check,
		"limit_per_day":            limit_per_day,
		"limit_per_hour":           limit_per_hour,
		"limit_per_user":           limit_per_user,
		"limit_per_user_folder":    limit_per_user_folder,
		"limit_per_ip":             limit_per_ip,
		"limit_only_for_level_id":  limit_only_for_level_id,
		"limit_date_from":          limit_date_from,
		"limit_date_to":            limit_date_to,
		"targeting_gender":         targeting_gender,
		"targeting_age_from":       targeting_age_from,
		"targeting_age_to":         targeting_age_to,
		"targeting_geo_country_id": targeting_geo_country_id,
		"targeting_geo_region_id":  targeting_geo_region_id,
		"targeting_geo_city_id":    targeting_geo_city_id,
		"task_only_for_list_id":    task_only_for_list_id,
		"list_of_pages":            list_of_pages,
	}
	if time_for_work == 0 {
		delete(params, "time_for_work")
	}
	if time_for_check == 0 {
		delete(params, "time_for_check")
	}
	if limit_per_day == 0 {
		delete(params, "limit_per_day")
	}
	if limit_per_hour == 0 {
		delete(params, "limit_per_hour")
	}
	if limit_per_user == 0 {
		delete(params, "limit_per_user")
	}
	if limit_per_user_folder == 0 {
		delete(params, "limit_per_user_folder")
	}
	if limit_per_ip == 0 {
		delete(params, "limit_per_ip")
	}
	if limit_only_for_level_id == 0 {
		delete(params, "limit_only_for_level_id")
	}
	if limit_date_from == "" {
		delete(params, "limit_date_from")
	}
	if limit_date_to == "" {
		delete(params, "limit_date_to")
	}
	if delay_from == 0 {
		delete(params, "delay_from")
	}
	if delay_to == 0 {
		delete(params, "delay_to")
	}
	if targeting_gender == 0 {
		delete(params, "targeting_gender")
	}
	if targeting_age_from == 0 {
		delete(params, "targeting_age_from")
	}
	if targeting_age_to == 0 {
		delete(params, "targeting_age_to")
	}
	if targeting_geo_country_id == 0 {
		delete(params, "targeting_geo_country_id")
	}
	if targeting_geo_region_id == 0 {
		delete(params, "targeting_geo_region_id")
	}
	if targeting_geo_city_id == 0 {
		delete(params, "targeting_geo_city_id")
	}
	if task_only_for_list_id == 0 {
		delete(params, "task_only_for_list_id")
	}
	if list_of_pages == "" {
		delete(params, "list_of_pages")
	}
	bytesRes := c.post(ctx, "edit_task", params)
	if err := json.Unmarshal([]byte(bytesRes), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Method: del_task // Удаляет задачу
// Входные данные:
// task_id (int) – идентификатор задачи
// Выходные данные отсутствуют
func (c *Client) Del_task(ctx context.Context, task_id int) (*models.Response, error) {
	var resp *models.Response
	params := map[string]interface{}{
		"task_id": task_id,
	}
	bytesRes := c.post(ctx, "del_task", params)
	if err := json.Unmarshal([]byte(bytesRes), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Method: get_tariffs // Возвращает все доступные тарифы.
// Входные данные отсутствует
// Выходные данные:
// tariffs (array) – массив с тарифами
// Каждый элемент массива содержит:
//
//	id (int) – уникальный идентификатор тарифа
//	name (text) – названия тарифа
//	min_price_rub (float) – минимальная стоимость в рублях
//	group_id (int) – идентификатор группы тарифов
func (c *Client) Get_tariffs(ctx context.Context) (*models.Response, error) {
	var resp *models.Response

	bytesRes := c.post(ctx, "get_tariffs", nil)
	if err := json.Unmarshal([]byte(bytesRes), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Method: get_countries // Возвращает список стран для таргетинга.
// Входные данные отсутствуют
// Выходные данные:
// countries (array) – массив со странами
// Каждый элемент массива содержит:
//
//	id (int) – уникальный идентификатор страны
//	name (text) – название страны
//	В настройках ГЕО-таргетинга задания имеется возможность выбрать параметр "СНГ и ближнее зарубежье" (ID 236).
//	Он включает в себя следующие страны: Россия, Азербайджан, Армения, Беларусь, Казахстан, Киргизия (Кыргызстан),
//	Молдова, Таджикистан, Узбекистан, Украина и Туркменистан.
func (c *Client) Get_countries(ctx context.Context) (*models.Response, error) {
	var resp *models.Response

	bytesRes := c.post(ctx, "get_countries", nil)
	if err := json.Unmarshal([]byte(bytesRes), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Method: task_pause // Приостанавливает выполнение задачи
// Входные данные:
// task_id (int) – идентификатор задачи
// Выходные данные отсутствуют
func (c *Client) Task_pause(ctx context.Context, task_id int) (*models.Response, error) {
	var resp *models.Response
	params := map[string]interface{}{
		"task_id": task_id,
	}
	bytesRes := c.post(ctx, "task_pause", params)
	if err := json.Unmarshal([]byte(bytesRes), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Method: task_play // Активирует выполнение задачи
// Входные данные:
// task_id (int) – идентификатор задачи
// Выходные данные отсутствуют
func (c *Client) Task_play(ctx context.Context, task_id int) (*models.Response, error) {
	var resp *models.Response
	params := map[string]interface{}{
		"task_id": task_id,
	}
	bytesRes := c.post(ctx, "task_play", params)
	if err := json.Unmarshal([]byte(bytesRes), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Method: task_to_top // Разово поднимает задачу в поиске (платная услуга)
// Входные данные:
// task_id (int) – идентификатор задачи
// Выходные данные отсутствуют
func (c *Client) Task_to_top(ctx context.Context, task_id int) (*models.Response, error) {
	var resp *models.Response
	params := map[string]interface{}{
		"task_id": task_id,
	}
	bytesRes := c.post(ctx, "task_to_top", params)
	if err := json.Unmarshal([]byte(bytesRes), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Method: add_blacklist // Добавляет пользователя в Чёрный список
// Входные данные:
// add_blacklist_id (int) – ID пользователя в системе
// Выходные данные отсутствуют
func (c *Client) Add_blacklist(ctx context.Context, add_blacklist_id int) (*models.Response, error) {
	var resp *models.Response
	params := map[string]interface{}{
		"add_blacklist_id": add_blacklist_id,
	}
	bytesRes := c.post(ctx, "add_blacklist", params)
	if err := json.Unmarshal([]byte(bytesRes), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Method: add_whitelist // Добавляет пользователя в Белый список
// Входные данные:
// add_whitelist_id (int) – ID пользователя в системе
// Выходные данные отсутствуют
func (c *Client) Add_whitelist(ctx context.Context, add_whitelist int) (*models.Response, error) {
	var resp *models.Response
	params := map[string]interface{}{
		"add_whitelist": add_whitelist,
	}
	bytesRes := c.post(ctx, "add_whitelist", params)
	if err := json.Unmarshal([]byte(bytesRes), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Method: get_blacklist // Возвращает ID пользователей из Чёрного списка
// Входные данные отсутствуют
// Выходные данные:
// users (array) – массив с ID пользователей, находящихся в Чёрном списке
func (c *Client) Get_blacklist(ctx context.Context) (*models.Response, error) {
	var resp *models.Response

	bytesRes := c.post(ctx, "get_blacklist", nil)
	if err := json.Unmarshal([]byte(bytesRes), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Method: delete_user_blacklist // Удаляет пользователя из Чёрного списка.
// Входные данные:
// id_user_blacklist (int) - идентификатор пользователя, находящегося в Чёрном списке
// Выходные данные отсутствуют
func (c *Client) Delete_user_blacklist(ctx context.Context, id_user_blacklist int) (*models.Response, error) {
	var resp *models.Response
	params := map[string]interface{}{
		"id_user_blacklist": id_user_blacklist,
	}
	bytesRes := c.post(ctx, "delete_user_blacklist", params)
	if err := json.Unmarshal([]byte(bytesRes), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
