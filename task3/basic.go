package task3

type Student struct {
	ID    int
	Name  string
	Age   int
	Grade string
}

//func (stu *Student) AfterCreate(tx *gorm.DB) (err error) {
//	if tx.Error != nil {
//		fmt.Println(tx.Error)
//	} else {
//		fmt.Printf("affected rows: %d\n", tx.RowsAffected)
//	}
//	return
//}
