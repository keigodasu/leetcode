select
activity_date day,
count(distinct user_id) active_users
from
Activity
where
activity_date between DATE_SUB('2019-07-27', INTERVAL 29 DAY) and '2019-07-27'
group by activity_date
