DROP TABLE IF EXISTS employees;
CREATE TABLE employees(
    id int primary key AUTO_INCREMENT,
    name varchar(20),
    dept varchar(20),
    salary decimal(8,2)
);

INSERT INTO employees(name, dept, salary) VALUES
('小明', '技术部', 8000.0),
('李四', '项目部', 3200.0),
('王五', '财务部', 4400.0),
('赵六', '人力部', 5300.0),
('钱七', '技术部', 8300.0),
('孙八', '项目部', 4600.0),
('周九', '财务部', 5000.0),
('吴十', '技术部', 7200.0),
('郑十一', '人力部', 4000.0),
('王十二', '技术部', 9000.0);