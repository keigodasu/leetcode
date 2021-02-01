with cte as (
	select
	row_number() over(partition by u.user_id order by o.order_date asc) as rownum,
	u.user_id,
	o.order_date,
	i.item_brand
	from
	users u
	left join
	orders o
	on u.user_id = o.seller_id
	left join
	items i
	using(item_id)
)

select
user_id as seller_id,
case when u.favorite_brand = (select item_brand from cte where u.user_id = user_id and rownum = 2 and u.favorite_brand = item_brand) then 'yes' else 'no' end as 2nd_item_fav_brand 
from
users u
