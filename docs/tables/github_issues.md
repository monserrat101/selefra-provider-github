# Table: github_issues

## Primary Keys 

```
org, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| milestone | json | X | √ |  | 
| repository_url | string | X | √ |  | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| author_association | string | X | √ |  | 
| labels | json | X | √ |  | 
| comments | int | X | √ |  | 
| locked | bool | X | √ |  | 
| closed_by | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| url | string | X | √ |  | 
| comments_url | string | X | √ |  | 
| node_id | string | X | √ |  | 
| number | int | X | √ |  | 
| state | string | X | √ |  | 
| body | string | X | √ |  | 
| closed_at | timestamp | X | √ |  | 
| title | string | X | √ |  | 
| events_url | string | X | √ |  | 
| text_matches | json | X | √ |  | 
| reactions | json | X | √ |  | 
| active_lock_reason | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| assignees | json | X | √ |  | 
| id | int | X | √ |  | 
| assignee | json | X | √ |  | 
| labels_url | string | X | √ |  | 
| pull_request | json | X | √ |  | 
| user | json | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| html_url | string | X | √ |  | 
| repository | json | X | √ |  | 


