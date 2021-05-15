with recursive months as
(
select 202001 as month
union
select month + 1 from months where month < 202012
)
,
drivers_with_year as (
select
driver_id,
month
from
months m
left join
drivers d
on
m.month >= date_format(d.join_date, '%Y%m')
)


select
(month - 202000) as month, 
ifnull(round(count(distinct a.driver_id) / count(distinct d.driver_id) * 100, 2), 0) as working_percentage 
from
rides r
join
AcceptedRides a
using(ride_id)
right join
drivers_with_year d
on date_format(r.requested_at, '%Y%m') = d.month
group by month
