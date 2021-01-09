select
Score,
dense_rank() over(order by score desc) `Rank`
from
scores
order by score desc;
