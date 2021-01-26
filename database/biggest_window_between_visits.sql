select
user_id,
max(t.window01) biggest_window
from
(
select
user_id,
abs(datediff(visit_date, ifnull(lead(visit_date) over(partition by user_id order by visit_date), '2021-1-1'))) window01
from
uservisits
) t
group by t.user_id

