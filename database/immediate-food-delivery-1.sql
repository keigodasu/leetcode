select
round((select count(1) from delivery where order_date = customer_pref_delivery_date)/count(1)*100, 2) immediate_percentage
from
delivery;


select
    round(sum(order_date = customer_pref_delivery_date)/count(1) * 100, 2) immediate_percentage
from
    Delivery;

