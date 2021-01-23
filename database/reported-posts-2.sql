select
round(avg(average) * 100 , 2) as average_daily_percent
from
(
select
count(distinct r.post_id) / count(distinct a.post_id) as average
from
actions a
left join
removals r
using(post_id)
where a.extra = 'spam'
group by a.action_date
) t

