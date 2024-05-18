CREATE TABLE departments (
    department_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE employees (
    employee_id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    department_id INT,
    email VARCHAR(255),
    FOREIGN KEY (department_id) REFERENCES departments(department_id)
);