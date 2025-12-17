DROP TABLE IF EXISTS t_account;
CREATE TABLE t_account(
    id int primary key AUTO_INCREMENT,
    account varchar(50),
    balance decimal(8, 2)
);

DROP TABLE IF EXISTS t_transaction;
CREATE TABLE t_transaction(
    id int primary key AUTO_INCREMENT,
    from_account varchar(50),
    to_account varchar(50),
    amount decimal(8, 2)
);

INSERT INTO t_account(account, balance) values
('wxq', 500),
('weq', 500);