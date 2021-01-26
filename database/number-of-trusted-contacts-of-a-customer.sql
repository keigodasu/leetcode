select
invoice_id,
cu.customer_name,
price,
sum(case when cn.user_id is not null then 1 else 0 end) as contacts_cnt,
count(cu2.customer_name) as trusted_contacts_cnt
from
invoices i
join
customers cu
on 
i.user_id = cu.customer_id
left join
contacts cn
on i.user_id = cn.user_id
left join
customers cu2
on cn.contact_name= cu2.customer_name
group by invoice_id
order by 1
