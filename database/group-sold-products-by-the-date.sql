select
sell_date,
count(distinct product) num_sold,
GROUP_CONCAT(distinct product order by product asc) products
from
Activities 
group by sell_date;
