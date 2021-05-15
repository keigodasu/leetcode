select
project_id
from
project
group by project_id
having count(*) = (
select 
    count(*) c
from
    project
group by project_id
order by c desc limit 1
)
