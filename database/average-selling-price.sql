select
u.product_id,
round(sum(u.units * p.price)/sum(u.units), 2) average_price
from
UnitsSold u
join
Prices p
on u.product_id = p.product_id and (u.purchase_date between p.start_date and p.end_date)
group by u.product_id
