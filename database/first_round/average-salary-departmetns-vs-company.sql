with avg_per_month as(
select
left(pay_date, 7) as pay_month,
avg(amount) as avg_amount
from
salary
group by 1
),
avg_per_month_by_department as (
select
left(pay_date, 7) as pay_month,
avg(amount) as avg_amount,
department_id,
employee_id
from
salary
join
employee
using(employee_id)
group by 1, department_id
)

select
ad.pay_month,
ad.department_id,
case
when ad.avg_amount > am.avg_amount then 'higher' 
when ad.avg_amount < am.avg_amount then 'lower' 
else 'same' end as comparison
from avg_per_month_by_department ad
join
avg_per_month am
on
ad.pay_month = am.pay_month
