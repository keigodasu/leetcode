with recursive cte as (
select 1 as month
union all
select month + 1 from cte where month < 12
),
ride_history as (
select
r.ride_id,
r.requested_at,
a.ride_distance,
a.ride_duration
from
rides r
join
AcceptedRides a
using(ride_id) 
where r.requested_at between '2020-1-1' and '2020-12-31'
)

select
c.month as month,
ifnull(round(sum(r.ride_distance) / 3, 2), 0) as average_ride_distance,
ifnull(round(sum(r.ride_duration) / 3, 2), 0) as average_ride_duration
from
cte c
left join
ride_history r
on date_format(r.requested_at, '%m') between c.month and c.month + 2
group by c.month
having month < 11
