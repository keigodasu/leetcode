select
  s.name
from
  salesperson s
where
  not exists (
    select
      1
    from
      orders
      join salesperson on orders.sales_id = salesperson.sales_id
      join company on orders.com_id = company.com_id
    where
      company.name = 'RED'
      and s.sales_id = orders.sales_id
  )

