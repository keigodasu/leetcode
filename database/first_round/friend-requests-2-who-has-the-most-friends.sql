select
requester_id as id,
count(*) as num
from
(
select
requester_id
from
request_accepted
union all
select
accepter_id as requester_id
from
request_accepted
) temp
group by requester_id
order by num desc limit 1
