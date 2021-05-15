select
e1.name
from
Employee e1
join
Employee e2
on e1.id = e2.managerid
group by e2.managerid
having count(*) >= 5
