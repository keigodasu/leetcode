with userstrips as (
select
client_id as users_id,
status,
request_at
from
trips
union all
select
client_id as users_id,
status,
request_at
from
trips
)

select
request_at as day,
round(avg(status = 'cancelled_by_client' or status = 'cancelled_by_driver'), 2) as 'Cancellation Rate'
from
userstrips
join
users
using(users_id)
where banned = 'No' and request_at between '2013-10-01' and '2013-10-03'
group by request_at
