# Table: github_team_members

## Primary Keys 

```
org, id, team_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| events_url | string | X | √ |  | 
| gravatar_id | string | X | √ |  | 
| private_gists | int | X | √ |  | 
| suspended_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| blog | string | X | √ |  | 
| twitter_username | string | X | √ |  | 
| email | string | X | √ |  | 
| public_gists | int | X | √ |  | 
| collaborators | int | X | √ |  | 
| gists_url | string | X | √ |  | 
| received_events_url | string | X | √ |  | 
| text_matches | json | X | √ |  | 
| id | int | X | √ |  | 
| html_url | string | X | √ |  | 
| company | string | X | √ |  | 
| hireable | bool | X | √ |  | 
| bio | string | X | √ |  | 
| following | int | X | √ |  | 
| plan | json | X | √ |  | 
| ldap_dn | string | X | √ |  | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| name | string | X | √ |  | 
| role_name | string | X | √ |  | 
| repos_url | string | X | √ |  | 
| permissions | json | X | √ |  | 
| public_repos | int | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| type | string | X | √ |  | 
| github_teams_selefra_id | string | X | X | fk to github_teams.selefra_id | 
| team_id | int | X | √ |  | 
| location | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| owned_private_repos | int | X | √ |  | 
| disk_usage | int | X | √ |  | 
| url | string | X | √ |  | 
| following_url | string | X | √ |  | 
| followers_url | string | X | √ |  | 
| node_id | string | X | √ |  | 
| avatar_url | string | X | √ |  | 
| organizations_url | string | X | √ |  | 
| starred_url | string | X | √ |  | 
| followers | int | X | √ |  | 
| site_admin | bool | X | √ |  | 
| total_private_repos | int | X | √ |  | 
| two_factor_authentication | bool | X | √ |  | 
| membership | json | X | √ |  | 
| login | string | X | √ |  | 
| subscriptions_url | string | X | √ |  | 


