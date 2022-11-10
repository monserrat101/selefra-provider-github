# Table: github_organization_members

## Primary Keys 

```
org, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| received_events_url | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| gravatar_id | string | X | √ |  | 
| public_repos | int | X | √ |  | 
| following | int | X | √ |  | 
| two_factor_authentication | bool | X | √ |  | 
| repos_url | string | X | √ |  | 
| role_name | string | X | √ |  | 
| login | string | X | √ |  | 
| name | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| private_gists | int | X | √ |  | 
| events_url | string | X | √ |  | 
| subscriptions_url | string | X | √ |  | 
| membership | json | X | √ |  | 
| company | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| owned_private_repos | int | X | √ |  | 
| node_id | string | X | √ |  | 
| public_gists | int | X | √ |  | 
| github_organizations_selefra_id | string | X | X | fk to github_organizations.selefra_id | 
| text_matches | json | X | √ |  | 
| permissions | json | X | √ |  | 
| email | string | X | √ |  | 
| site_admin | bool | X | √ |  | 
| total_private_repos | int | X | √ |  | 
| starred_url | string | X | √ |  | 
| location | string | X | √ |  | 
| bio | string | X | √ |  | 
| ldap_dn | string | X | √ |  | 
| plan | json | X | √ |  | 
| url | string | X | √ |  | 
| following_url | string | X | √ |  | 
| gists_url | string | X | √ |  | 
| id | int | X | √ |  | 
| hireable | bool | X | √ |  | 
| twitter_username | string | X | √ |  | 
| suspended_at | timestamp | X | √ |  | 
| type | string | X | √ |  | 
| disk_usage | int | X | √ |  | 
| collaborators | int | X | √ |  | 
| followers_url | string | X | √ |  | 
| avatar_url | string | X | √ |  | 
| html_url | string | X | √ |  | 
| blog | string | X | √ |  | 
| followers | int | X | √ |  | 
| organizations_url | string | X | √ |  | 


