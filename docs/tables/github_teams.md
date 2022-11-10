# Table: github_teams

## Primary Keys 

```
org, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| html_url | string | X | √ |  | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| id | int | X | √ |  | 
| permissions | json | X | √ |  | 
| privacy | string | X | √ |  | 
| organization | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| slug | string | X | √ |  | 
| members_count | int | X | √ |  | 
| members_url | string | X | √ |  | 
| parent | json | X | √ |  | 
| ldap_dn | string | X | √ |  | 
| node_id | string | X | √ |  | 
| name | string | X | √ |  | 
| description | string | X | √ |  | 
| permission | string | X | √ |  | 
| repositories_url | string | X | √ |  | 
| url | string | X | √ |  | 
| repos_count | int | X | √ |  | 


