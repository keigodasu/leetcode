select
distinct customer_number
from
orders
where customer_number =
(
select
customer_number
from
orders
group by customer_number
order by count(*) desc limit 1
);

select 
customer_number 
from orders
group by customer_number
order by count(*) desc limit 1;
