select
project_id,
employee_id
from
project p
join
employee e
using(employee_id)
where (project_id, experience_years) in 
(select project_id, max(e.experience_years) from project p join employee e using(employee_id) group by p.project_id)
