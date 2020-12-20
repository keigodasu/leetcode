SELECT
MIN(ABS(p1.x - p2.x)) AS shortest
FROM
point p1 join point p2 on p1.x != p2.x;
