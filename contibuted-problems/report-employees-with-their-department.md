# Question Title
report employees with their department

# Category
database

# Reason
* Where have you seen this question?
- I often see this type of question in practice.

* Was it in a coding challenge, phone screen, or an on-site interview?
- not yet.

* How difficult do you think the question is?
- the middle between easy and middle.

* Is there anything special about this question that motivates you to contirubte?
-  I believe this type of question could make SQL user flexible. Especially, this problem teach users effective use case of SQL case statement.

# Description
Table: Employee

+--------------+---------+
| Column Name  |  Type   |
+--------------+---------+
| employee_id  | varchar |
| deprtment_id | varchar |
| primary_flag | varchar |
+--------------+---------+

employee_od and department_id is primary key for this table.(composite key)
employee_id is the id of employee.
department_id is the id of department employee belongs to.
primary_flag is enum which are 'yes' or 'no'. if 'yes', the department is primary for the employee. If 'no', the department is not primary.
Employees can belong to multiple departments. When the employee join other departments, they need to decide which department is primary. When an employee belongs to only one department,  primary column is 'no'.

Write an SQL query to report all employees with their primary department. For employees who belong to one department, report their department. 
Return the result table in any order.
The query result format is in the following example.


Employee table:
+-------------+---------------+--------------+
| employee_id | department_id | primary_flag |
+-------------+---------------+--------------+
|           1 |             1 | no           |
|           2 |             1 | yes          |
|           2 |             2 | no           |
|           3 |             3 | no           |
|           4 |             2 | no           |
|           4 |             3 | yes          |
|           4 |             4 | no           |
+-------------+---------------+--------------+

Result table:
+-------------+---------------+
| employee_id | department_id |
+-------------+---------------+
|           1 |             1 |
|           2 |             1 |
|           3 |             3 |
|           4 |             3 |
+-------------+---------------+

# Solutions
Approach: Using CASE WHEN

SELECT employee_id,
       CASE
         WHEN Count(*) = 1 THEN Max(department_id)
         ELSE Max(CASE
                    WHEN primary_flag = 'yes' THEN department_id
                  end)
       end AS department_id
FROM   empoloyee
GROUP  BY employee_id;

# Test Cases
{"headers":{"Employee":["employee_id","deprtment_id","primary_flag"]},"rows":{"Employee":[["1","1","no"],["2","1","yes"],["2","2","no"],["3","3","no"],["4","2","no"],["4","3","yes"],["4","4","no"]]}}

# Tags

