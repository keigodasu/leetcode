select 
    seller_id
from
    (
        select
            seller_id,
            dense_rank() over (order by sum(price) desc) as rnk
        from
            Sales
        group by seller_id
    ) t
where
    t.rnk = 1;
    
