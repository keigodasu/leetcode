select round(min(sqrt(power(abs(p1.x - p2.x), 2) + power(abs(p1.y - p2.y), 2))), 2) as shortest 
from 
point_2d p1
join
point_2d p2
on p1.x != p2.x or p1.y != p2.y

