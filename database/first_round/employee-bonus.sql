select
e.name,
b.bonus
from
Employee e
left outer join
Bonus b
on e.empId = b.empId
where b.bonus < 1000 || b.bonus is null
