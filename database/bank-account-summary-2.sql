select
name,
sum(amount) balance
from
transactions
join
users
using(account)
group by account
having sum(amount) > 10000
