# Table: github_external_groups

## Primary Keys 

```
org, group_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| group_id | int | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| group_name | string | X | √ |  | 
| teams | json | X | √ |  | 
| members | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


