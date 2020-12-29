select
country_name,
case 
 when avg(weather_state) <= 15 then 'Cold' 
 when avg(weather_state) >= 15 and avg(weather_state) < 25 then 'Warm' 
 when avg(weather_state) >= 25 then 'Hot' 
 end weather_type
from
weather
join
Countries
using(country_id)
where
day like '2019-11-%'
group by country_id;
