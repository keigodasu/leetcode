select
user_id buyer_id,
join_date,
sum(case when left(order_date, 4) = '2019' then 1 else 0 end) orders_in_2019
from
users u
left join
orders o
on u.user_id = o.buyer_id
group by u.user_id;
