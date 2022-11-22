package md

func initSqlTempReg() {
	loglist := "select * from `log` as a left join `task` as b on a.`tid`=b.`id` where a.`id`!=''" +
		"{% if name %}" +
		" and a.`name` like '%{{name}}%' " +
		"{% endif %}" +
		"{% if taskname %}" +
		" and b.`name` like '%{{taskname}}%' " +
		"{% endif %}" +
		"{% if dbname %}" +
		" and b.`dbname` like '%{{dbname}}%' " +
		"{% endif %}" +
		" group by a.`id` order by a.`created` desc " +
		"{% if pageSize && currentPage%}" +
		" limit {{pageSize*(currentPage-1)}},{{pageSize}} " +
		"{% endif %}"
	if err := localdb.SqlTemplate.AddSqlTemplate("loglist", loglist); err != nil {
		panic(err)
	}

	tasklist := "select * from `task` where `id`!=''" +
		"{% if name %}" +
		" and `name` like '%{{name}}%' " +
		"{% endif %}" +
		"{% if dbname %}" +
		" and `dbname` like '%{{dbname}}%' " +
		"{% endif %}" +
		"{% if selectTypes==\"sql\" %}" +
		" and (`dbtype`='mysql' or `dbtype`='mssql' or `dbtype`='sqlite' or `dbtype`='postgres' or `dbtype`='mongodb') " +
		"{% else %}" +
		" and `dbtype`=='file' " +
		"{% endif %}" +
		" order by `created` desc " +
		"{% if pageSize && currentPage%}" +
		" limit {{pageSize*(currentPage-1)}},{{pageSize}} " +
		"{% endif %}"
	if err := localdb.SqlTemplate.AddSqlTemplate("tasklist", tasklist); err != nil {
		panic(err)
	}
}
