select
user_id,
user_name,
credit,
case when credit < 0 then "Yes" else "No" end as credit_limit_breached
from
(
select
u.user_id as USER_ID,
u.user_name as USER_NAME,
ifnull(sum(case when u.user_id = t.paid_by then -amount else amount end) + u.credit, u.credit) as CREDIT
from
users u
left join
transactions t
on u.user_id = t.paid_by or u.user_id = t.paid_to
group by u.user_id
) t

