select
ad_id,
ifnull(round(sum(if(action='Clicked',1,0))/sum(if(action = 'Clicked' or action = 'Viewed', 1, 0))*100, 2), 0) ctr
from
ads
group by ad_id
order by ctr desc, ad_id asc;
