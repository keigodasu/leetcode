select
company_id,
employee_id,
employee_name,
round(salary * (100-t.tax_rate) / 100, 0) salary
from
salaries
join
(
select
company_id,
case 
when max(salary) > 10000 then 49
when max(salary) between 1000 and 10000 then 24
else 0 end tax_rate
from
salaries
group by company_id
) t
using(company_id)
