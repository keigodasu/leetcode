select
*
from(
select
product_id,
'store1' as store,
store1 as price
from
(select product_id, store1 from products) t
union all
select
product_id,
'store2' as store,
store2 as price
from
(select product_id, store2 from products) t
union all
select
product_id,
'store3' as store,
store3 as price
from
(select product_id, store3 from products) t
) t
where price is not null

