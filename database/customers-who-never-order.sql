select
Name Customers
from
Customers
where
not exists
(select 1 from Orders where Customers.Id = Orders.CustomerId)
