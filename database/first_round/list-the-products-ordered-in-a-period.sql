select
p.product_name PRODUCT_NAME,
sum(o.unit) UNIT
from
Orders o
join
Products p
using(product_id)
where order_date like '2020-02-%'
group by o.product_id
having sum(o.unit) >= 100;
