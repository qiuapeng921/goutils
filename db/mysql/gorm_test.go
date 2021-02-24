package mysql

import "testing"

func TestConnect(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	if err := Connect(dsn, 10, 10, 10); err != nil {
		t.Errorf("Connect() error = %v", err)
	}
}

type Test struct {
	Id int
	Name string
}

func BenchmarkDB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var test Test
		test.Name = "test"
		Client().Create(&test)
	}
}