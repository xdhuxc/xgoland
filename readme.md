### 参考资料
1、golang 操作 mysql 数据库

https://www.jianshu.com/p/340eb943be2e

```
db, err := sql.Open("mysql", "root:19940423@tcp(127.0.0.1:3306)/xgolang?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	if db != nil {
		fmt.Println(db)
	}
	/**
		调用 sql.Close() 函数释放连接
		使用 defer 语句设置释放连接
	 */
	defer db.Close()
	queryString := "insert into user (username, password, age) values ('xdhuxc', 'Xdhuxc123', 24)"
	result, err := db.Exec(queryString)
	if err != nil {
		log.Fatal(err)
	}
	rowCount, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("插入了 %d 行。", rowCount)
```

2、xorm 资料

http://xorm.io/docs
