# Table: github_hooks

## Primary Keys 

```
org, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | int | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| test_url | string | X | √ |  | 
| type | string | X | √ |  | 
| events | string_array | X | √ |  | 
| active | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| url | string | X | √ |  | 
| ping_url | string | X | √ |  | 
| last_response | json | X | √ |  | 
| config | json | X | √ |  | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| updated_at | timestamp | X | √ |  | 
| name | string | X | √ |  | 


