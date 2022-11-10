# Table: github_organizations

## Primary Keys 

```
org, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| blog | string | X | √ |  | 
| description | string | X | √ |  | 
| public_repos | int | X | √ |  | 
| billing_email | string | X | √ |  | 
| members_can_create_repositories | bool | X | √ |  | 
| members_url | string | X | √ |  | 
| html_url | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| has_repository_projects | bool | X | √ |  | 
| members_can_fork_private_repositories | bool | X | √ |  | 
| url | string | X | √ |  | 
| avatar_url | string | X | √ |  | 
| total_private_repos | int | X | √ |  | 
| members_can_create_public_repositories | bool | X | √ |  | 
| public_members_url | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| two_factor_requirement_enabled | bool | X | √ |  | 
| has_organization_projects | bool | X | √ |  | 
| members_allowed_repository_creation_type | string | X | √ |  | 
| members_can_create_pages | bool | X | √ |  | 
| hooks_url | string | X | √ |  | 
| twitter_username | string | X | √ |  | 
| location | string | X | √ |  | 
| public_gists | int | X | √ |  | 
| collaborators | int | X | √ |  | 
| default_repository_permission | string | X | √ |  | 
| members_can_create_internal_repositories | bool | X | √ |  | 
| members_can_create_private_pages | bool | X | √ |  | 
| events_url | string | X | √ |  | 
| company | string | X | √ |  | 
| email | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| owned_private_repos | int | X | √ |  | 
| plan | json | X | √ |  | 
| login | string | X | √ |  | 
| disk_usage | int | X | √ |  | 
| type | string | X | √ |  | 
| is_verified | bool | X | √ |  | 
| default_repository_settings | string | X | √ |  | 
| members_can_create_private_repositories | bool | X | √ |  | 
| issues_url | string | X | √ |  | 
| repos_url | string | X | √ |  | 
| node_id | string | X | √ |  | 
| id | int | X | √ |  | 
| followers | int | X | √ |  | 
| following | int | X | √ |  | 
| private_gists | int | X | √ |  | 
| members_can_create_public_pages | bool | X | √ |  | 
| org | string | X | √ | `The Github Organization of the resource.` | 


