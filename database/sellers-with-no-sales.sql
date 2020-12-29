select
seller_name
from
seller
where
seller_id not in (
select
s.seller_id
from
seller s
join
orders o
using(seller_id)
where
left(sale_date, 4) = '2020'
)
order by seller_name asc
