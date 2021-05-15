select
employee_id,
t.team_size
from
Employee
join
(
select
team_id,
count(*) team_size
from
Employee
group by team_id
) t
using(team_id);

select 
employee_id, 
count(*) over(partition by team_id) as team_size
from employee;
