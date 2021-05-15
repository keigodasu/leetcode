ve cte as (
select 1 as month, concat('2020-', lpad(1, 2, 0)) as yearmonth
union all
select month + 1, concat('2020-', lpad(month + 1, 2, 0)) from cte where month < 12
),
active_drivers as (
select 
        c.yearmonth, c.month, sum(case when driver_id is not null then 1 else 0 end) as active_drivers
from
cte c
left join
drivers d
on c.yearmonth >= date_format(d.join_date, '%Y-%m')
group by yearmonth
),
accepted_rides as (
select
date_format(requested_at, '%Y-%m') as yearmonth,
count(*) as accepted_rides
from
acceptedrides a
join
rides r
using(ride_id)
group by 1
)

select
d.month as month,
active_drivers,
ifnull(accepted_rides, 0) as accepted_rides
from
active_drivers d
left join
accepted_rides r
on d.yearmonth = r.yearmonth
