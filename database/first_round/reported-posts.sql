select
    extra report_reason,
    count(distinct(post_id)) report_count
from
Actions
where
    action = 'report'
    and
    action_date = '2019-07-04'
group by extra
