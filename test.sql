ダミーデータの生成は以下の通りです：
```
INSERT INTO departments (name)
VALUES ('Department1'), ('Department2'), ('Department3'), ('Department4'), ('Department5');

SET @count = 0;
INSERT INTO employees (first_name, last_name, department_id, email)
SELECT 
  CONCAT('First', @count := @count + 1),
  CONCAT('Last', @count),
  MOD(@count, 5) + 1,
  CONCAT('email', @count, '@example.com')
FROM
  (
    SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4 UNION ALL 
    SELECT 5 UNION ALL SELECT 6 UNION ALL SELECT 7 UNION ALL SELECT 8 UNION ALL SELECT 9
  ) a,
  (
    SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4 UNION ALL 
    SELECT 5 UNION ALL SELECT 6 UNION ALL SELECT 7 UNION ALL SELECT 8 UNION ALL SELECT 9
  ) b
LIMIT 100;
```
上記のクエリはまず初めに5つの部門をdepartmentsテーブルに挿入しています。その後、employeesテーブルに100行のデータを挿入します。- employeesの行には、一意の名前とメールアドレスとともにランダムなdepartment_idが含まれます。このクエリは最大100行のダミーデータの生成に対応しており、より多くの行が必要な場合は、必要に応じてクエリを調整することができます。
