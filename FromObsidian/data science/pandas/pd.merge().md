1. 여러 테이블 join하기
```python
merged_df = df1.merge(df2, how = 'left', on = 'key')
			   .merge(df3, how = 'left', on = 'key')
```

2. 두 테이블 & 여러 개의 key로 join하기
```python
df1.merge(df2, on = '공통된 열')

# 공통된 열이 여러개라면
df1.merge(df2, on = '공통열1, 공통열2, 공통열3')
```

