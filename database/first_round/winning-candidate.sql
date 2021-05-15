select
name 
from
(
select
candidateid,
name,
count(*) votes
from
vote v
join
candidate c
on c.id = v.candidateid
group by candidateid
order by votes desc limit 1
) t
