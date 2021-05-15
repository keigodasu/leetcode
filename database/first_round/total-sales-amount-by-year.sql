with recursive cte as
(
select min(period_start) as date from sales
union all
select date_add(date, interval 1 day) from cte
where date <= (select max(period_end) from sales)
)

select
s.product_id,
p.product_name,
left(c.date, 4) as report_year,
sum(s.average_daily_sales) as total_amount
from
sales s
join
product p
using(product_id)
join
cte c
on s.period_start <= c.date and s.period_end >= c.date 
group by 1,3
order by 1,3
 