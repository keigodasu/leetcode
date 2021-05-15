with employee_with_rnk as (
select
Id,
Name,
Salary,
DepartmentId,
rank() over(partition by DepartmentId order by Salary desc) as rnk
from
Employee
)

select
d.Name as Department,
e.Name as Employee,
e.Salary
from
employee_with_rnk e
join
Department d
on e.DepartmentId = d.Id
where e.rnk = 1
