CREATE DATABASE sql_p1;

use sql_p1;
drop table if exists retail_sales;

CREATE TABLE retail_sales(
  transaction_id int primary key,
  sale_date date,
  sale_time time,
  customer_id int,
  gender varchar(15),
  age int,
  category varchar(25),
  quantity int,
  price_per_unit float,
  cogs float,
  total_sale float
);

select * from retail_sales limit 10;

select * from retail_sales where transaction_id  is null;
select * from retail_sales where sale_date  is null;
select * from retail_sales where sale_time  is null;

select * from retail_sales
where 
	transaction_id is null
	or sale_date is NULL 
	or sale_time is null 
	or gender is null 
	or category is null 
	or cogs is null 
	or total_sale is null 
	;

# 数据清理
delete from retail_sales 
where 
	transaction_id is null
	or sale_date is NULL 
	or sale_time is null 
	or gender is null 
	or category is null 
	or cogs is null 
	or total_sale is null;

# 检查有多少销售记录
select count(*) as total_sale_count from retail_sales;


# 检查有多少客户
select count(customer_id) as customer_count from retail_sales; -- 会被重复计算

select COUNT(distinct customer_id) as customer_count from retail_sales; -- 使用distinct去重

select count(distinct category) as category_count from retail_sales; -- 查看有多少分类

select distinct category from retail_sales; -- 查看所有的分类

# 查询某一天的销售
select * from retail_sales where sale_date = '2022-11-05';


SELECT * FROM retail_sales 
WHERE category  = 'Clothing'
AND sale_date >= '2022-11-01'
AND sale_date <= '2022-12-01'
AND quantity >= 3;


# 从2022年11月的Clothing类销售中,按每天分组,汇总出每天卖了多少件衣服
select category, sale_date, sum(quantity) as num from retail_sales
	where category = 'Clothing' 
	and sale_date >= '2022-11-01'
	and sale_date <= '2022-12-01'
	group by category, sale_date;


# NOTE: WHERE 是在分组前筛选，HAVING 是在分组后筛选
# SQL执行顺序 FROM → WHERE → GROUP BY → HAVING → SELECT → ORDER BY

select category, sale_date, sum(quantity) as num from retail_sales
	where category = 'Clothing' 
	and sale_date >= '2022-11-01'
	and sale_date <= '2022-12-01'
	group by category, sale_date
	HAVING num > 5
	;

# 查询每一个分类的销售额和销售数量
SELECT 
	category,
	SUM(total_sale) as sale,
	COUNT(*) as num
FROM retail_sales
GROUP BY category;

# 查询买美妆产品用户平均年龄
SELECT AVG(age)as avg_age from retail_sales WHERE category = 'Beauty';

# 编写 SQL 查询来查找每个类别中每个性别进行的交易总数
SELECT 
	category, gender, SUM(1) as num 
from retail_sales 
group by
	category, gender
ORDER by category;

-- 或者

SELECT 
	category, gender, count(*) as num 
from retail_sales 
group by
	category, gender
ORDER BY category;

# 编写一个 SQL 查询来计算每日的平均销售额。找出每年最畅销的月份
SELECT
	sale_date, 
	AVG(total_sale) as total_sale 
FROM 
	retail_sales
group by sale_date 
ORDER BY total_sale DESC 
; 

# 编写一个 SQL 查询来计算每个月的平均销售额。找出每年最畅销的月份
SELECT
	DATE_FORMAT(sale_date, "%Y-%m") as ymoth,
	ROUND(AVG(total_sale), 2) as avg_sale -- 保留2位小数
FROM 
	retail_sales
group by
	ymoth
order by 
	avg_sale desc; 
;









