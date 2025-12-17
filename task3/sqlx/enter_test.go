package sqlx

import (
	_ "database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func TestMain(m *testing.M) {
	var err error
	dsn := "root:123456@tcp(192.168.3.124:3306)/homework?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("open db err:%v", err)
		return
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(60 * time.Second)

	code := m.Run()
	os.Exit(code)
}

func TestQueryTechDeptRows(t *testing.T) {
	sql := "select * from employees where dept=?"
	rows, err := db.Queryx(sql, "技术部")
	if err != nil {
		fmt.Printf("query err:%v", err)
		return
	}

	var res []Employees
	for rows.Next() {
		var employee Employees
		err := rows.StructScan(&employee)
		if err != nil {
			fmt.Printf("scan err:%v", err)
			return
		}
		res = append(res, employee)
	}
	fmt.Printf("res:%v \n", res)
}

func TestFindMaxSalaryEmployee(t *testing.T) {
	sql := "select * from employees where salary = (select max(salary) from employees) limit 1"
	row := db.QueryRowx(sql)
	var employee Employees
	if err := row.StructScan(&employee); err != nil {
		fmt.Printf("scan err:%v", err)
		return
	}
	fmt.Printf("employee:%v \n", employee)
}

func TestExecuteSimpleSql(t *testing.T) {
	tx, err := db.Begin()
	defer func() {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			t.Errorf("rollbackErr:%v", rollbackErr)
		}
	}()

	if err != nil {
		fmt.Printf("begin tx err:%v", err)
		return
	}
	sql := "insert into employees(name, dept, salary) values (?, ?, ?)"
	_, err = tx.Exec(sql, "wxq", "技术部", "9000.0")
	if err != nil {
		t.Errorf("exec sql err:%v", err)
		return
	}

	var employee Employees
	row := tx.QueryRow("select * from employees where name = ?", "wxq")
	err = row.Scan(&employee.ID, &employee.Name, &employee.Dept, &employee.Salary)
	if err != nil {
		t.Errorf("scan err:%v", err)
		return
	}
	fmt.Printf("employee:%v \n", employee)
}

func TestGet(t *testing.T) {
	var employee Employees
	sql := "select * from employees where id = ?"
	err := db.Get(&employee, sql, 1)
	if err != nil {
		t.Errorf("get err:%v", err)
		return
	}
	fmt.Printf("employee:%v \n", employee)
}

func TestSelect(t *testing.T) {
	var res []Employees
	sql := "select id, name, salary from employees where dept = ?"
	err := db.Select(&res, sql, "技术部")
	if err != nil {
		t.Errorf("select err:%v", err)
		return
	}
	fmt.Printf("res:%v \n", res)
}

func TestPrepareStatement(t *testing.T) {
	sql := "select id, name, salary from employees where dept = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		t.Errorf("prepare statement err:%v", err)
		return
	}
	rows, err := stmt.Query("技术部")
	if err != nil {
		t.Errorf("query err:%v", err)
		return
	}

	var arr []Employees
	for rows.Next() {
		var employee Employees
		err = rows.Scan(&employee.ID, &employee.Name, &employee.Salary)
		if err != nil {
			t.Errorf("scan err:%v", err)
			return
		}
		arr = append(arr, employee)
	}
	fmt.Printf("arr:%v \n", arr)
}

func TestNamedPlaceholder(t *testing.T) {
	sql := "select id, name, salary from employees where dept = :dept"
	params := map[string]interface{}{
		"dept": "技术部",
	}
	rows, err := db.NamedQuery(sql, params)
	if err != nil {
		t.Errorf("named placeholder err:%v", err)
		return
	}
	var arr []Employees
	for rows.Next() {
		var employee Employees
		err = rows.Scan(&employee.ID, &employee.Name, &employee.Salary)
		if err != nil {
			t.Errorf("scan err:%v", err)
			return
		}
		arr = append(arr, employee)
	}
	fmt.Printf("arr:%v \n", arr)
}

func TestRangeQuery(t *testing.T) {
	ids := []int{1, 2, 3}
	sql := "select * from employees where id in (?)"
	query, args, err := sqlx.In(sql, ids)
	//query = db.Rebind(query)
	rows, err := db.Query(query, args...)
	if err != nil {
		t.Errorf("query err:%v", err)
		return
	}
	var arr []Employees
	for rows.Next() {
		var employee Employees
		err = rows.Scan(&employee.ID, &employee.Name, &employee.Dept, &employee.Salary)
		if err != nil {
			t.Errorf("scan err:%v", err)
			return
		}
		arr = append(arr, employee)
	}
	fmt.Printf("arr:%v \n", arr)
}
