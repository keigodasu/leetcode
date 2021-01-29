select
department,
employee,
salary
from
(
select
d.name department,
e.name employee,
dense_rank() over(partition by d.name order by salary desc) rnk,
salary
from
employee e
join
department d
on 
e.departmentid = d.id
) t
where rnk between 1 and 3