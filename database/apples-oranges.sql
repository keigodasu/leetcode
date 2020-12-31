select
sale_date,
sum(case when fruit = 'apples' then sold_num end) - sum(case when fruit = 'oranges' then sold_num end) diff
from
sales
group by sale_date;

select
sale_date,
sum(case when fruit = 'apples' then sold_num else -sold_num end) diff
from
sales
group by sale_date;
