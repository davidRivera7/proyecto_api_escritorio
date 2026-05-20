CREATE DATABASE IF NOT EXISTS employees;
USE employees;

DROP TABLE IF EXISTS employees, 
                     departments,
                     dept_manager;
                     
CREATE TABLE employees (
    emp_no      INT             NOT NULL,
    birth_date  DATE            NOT NULL,
    first_name  VARCHAR(14)     NOT NULL,
    last_name   VARCHAR(16)     NOT NULL,
    gender      ENUM ('M','F')  NOT NULL,    
    hire_date   DATE            NOT NULL,
    PRIMARY KEY (emp_no)
);

CREATE TABLE departments (
    dept_no     CHAR(4)         NOT NULL,
    dept_name   VARCHAR(40)     NOT NULL,
    PRIMARY KEY (dept_no),
    UNIQUE  KEY (dept_name)
);

CREATE TABLE dept_manager (
   emp_no       INT             NOT NULL,
   dept_no      CHAR(4)         NOT NULL,
   from_date    DATE            NOT NULL,
   to_date      DATE            NOT NULL,
   FOREIGN KEY (emp_no)  REFERENCES employees (emp_no)    ON DELETE CASCADE,
   FOREIGN KEY (dept_no) REFERENCES departments (dept_no) ON DELETE CASCADE,
   PRIMARY KEY (emp_no,dept_no)
); 

INSERT INTO `employees` VALUES 
(10001,'2003-09-27','David','Rivera','M','2025-06-22'),
(10002,'1964-06-02','Eduardo','Ponce','M','1985-11-21'),
(10003,'1959-12-03','Selene','Castro','F','1986-08-28');

INSERT INTO `departments` VALUES 
('d001','Marketing'),
('d002','Finance'),
('d003','Human Resources'),
('d004','Production'),
('d005','Development');

INSERT INTO `dept_manager` VALUES 
(10001,'d001','1985-01-01','1991-10-01'),

(10002,'d005','1991-10-01','9999-01-01'),

(10001,'d002','1985-01-01','1989-12-17');