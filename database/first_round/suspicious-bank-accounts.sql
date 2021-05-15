with temp as (
select
t.account_id,
date_format(day, '%Y%m') as date,
sum(amount) as income,
a.max_income
from
transactions t
left join accounts a
using(account_id)
where t.type='Creditor'
group by t.account_id, date_format(t.day, '%Y%m')
having sum(t.amount) > a.max_income
)

select
t1.account_id
from
temp t1, temp t2
where t1.account_id = t2.account_id
and
period_diff(t1.date, t2.date) = 1
group by t1.account_id

