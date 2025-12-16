-- 创建学生表
DROP TABLE IF EXISTS student;
CREATE TABLE student (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    age INT NOT NULL,
    grade VARCHAR(20) NOT NULL
);

-- 插入10条学生数据
INSERT INTO student (name, age, grade) VALUES
('小明', 18, '高三'),
('李四', 17, '高二'),
('王五', 16, '高一'),
('赵六', 15, '初三'),
('钱七', 14, '初二'),
('孙八', 13, '初一'),
('周九', 12, '小学六年级'),
('吴十', 19, '小学五年级'),
('郑十一', 10, '小学四年级'),
('王十二', 9, '小学三年级');