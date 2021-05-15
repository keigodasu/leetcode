WITH TEMP AS (SELECT E.ID,
	                     E.COMPANY,
								 E.SALARY,
													 ROW_NUMBER() OVER (PARTITION BY E.COMPANY ORDER BY E.SALARY) AS RNO,
																		 COUNT(SALARY) OVER (PARTITION BY E.COMPANY) AS CNT
																					  FROM EMPLOYEE E)
																				SELECT ID,
																				       COMPANY,
																					   SALARY
																					FROM TEMP
																					WHERE RNO = CEIL(CNT/2) OR RNO = CNT/2 +1
