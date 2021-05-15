select
customer_id,
name
from
orders
join
customers
using(customer_id)
join
product
using(product_id)
group by customer_id
having
sum(if(left(order_date, 7) = '2020-06', quantity, 0) * price) >= 100
and
sum(if(left(order_date, 7) = '2020-07', quantity, 0) * price) >= 100

