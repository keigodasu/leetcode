select
query_name,
round(avg((rating/position)), 2) quality,
round(sum(if(rating < 3, 1, 0)) / count(*) * 100, 2) poor_query_percentage
from
Queries
group by query_name;
