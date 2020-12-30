select
s.id,
s.name
from
students s
left join
departments d
on s.department_id = d.id
where d.id is null;

select
id,
name
from students
where
department_id not in 
(select id from Departments);
