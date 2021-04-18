select transaction_id
from 
(select transaction_id, rank() over (partition by date(day) order by amount desc) rnk
from Transactions) t
where rnk = 1
order by 1
