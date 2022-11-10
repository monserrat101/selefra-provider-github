# Table: github_installations

## Primary Keys 

```
org, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| app_slug | string | X | √ |  | 
| repositories_url | string | X | √ |  | 
| single_file_name | string | X | √ |  | 
| has_multiple_single_files | bool | X | √ |  | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| updated_at | timestamp | X | √ |  | 
| account | json | X | √ |  | 
| node_id | string | X | √ |  | 
| app_id | int | X | √ |  | 
| access_tokens_url | string | X | √ |  | 
| target_type | string | X | √ |  | 
| single_file_paths | string_array | X | √ |  | 
| permissions | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| id | int | X | √ |  | 
| suspended_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| suspended_by | json | X | √ |  | 
| html_url | string | X | √ |  | 
| repository_selection | string | X | √ |  | 
| events | string_array | X | √ |  | 
| target_id | int | X | √ |  | 


