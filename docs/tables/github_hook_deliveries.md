# Table: github_hook_deliveries

## Primary Keys 

```
org, id, hook_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| action | string | X | √ |  | 
| installation_id | int | X | √ |  | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| redelivery | bool | X | √ |  | 
| status | string | X | √ |  | 
| status_code | int | X | √ |  | 
| request | string | X | √ |  | 
| github_hooks_selefra_id | string | X | X | fk to github_hooks.selefra_id | 
| duration | float | X | √ |  | 
| repository_id | int | X | √ |  | 
| hook_id | int | X | √ | `Hook ID` | 
| response | string | X | √ |  | 
| delivered_at | timestamp | X | √ |  | 
| guid | string | X | √ |  | 
| id | int | X | √ |  | 
| event | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


