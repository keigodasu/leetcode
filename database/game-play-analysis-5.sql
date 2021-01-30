with installedDate as (
select
player_id,
min(event_date) as install_dt
from
activity
group by player_id
)

select
install_dt,
count(*) installs,
round(sum(case when event_date is not null then 1 else 0 end)/count(*), 2) as day1_retention
from
installedDate i
left join
activity a
on 
i.player_id = a.player_id
and
a.event_date = date_add(i.install_dt, interval 1 day)
group by i.install_dt;