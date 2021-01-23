select
e2.employee_id,
e2.name,
count(*) as reports_count,
round(avg(e1.age)) as average_age
from
employees e1
join
employees e2
on e1.reports_to = e2.employee_id
group by e1.reports_to
order by 1
