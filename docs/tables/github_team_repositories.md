# Table: github_team_repositories

## Primary Keys 

```
org, id, team_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| default_branch | string | X | √ |  | 
| blobs_url | string | X | √ |  | 
| deployments_url | string | X | √ |  | 
| subscribers_url | string | X | √ |  | 
| statuses_url | string | X | √ |  | 
| org | string | X | √ | `The Github Organization of the resource.` | 
| code_of_conduct | json | X | √ |  | 
| open_issues | int | X | √ |  | 
| allow_auto_merge | bool | X | √ |  | 
| delete_branch_on_merge | bool | X | √ |  | 
| topics | string_array | X | √ |  | 
| issue_events_url | string | X | √ |  | 
| milestones_url | string | X | √ |  | 
| clone_url | string | X | √ |  | 
| mirror_url | string | X | √ |  | 
| open_issues_count | int | X | √ |  | 
| url | string | X | √ |  | 
| compare_url | string | X | √ |  | 
| events_url | string | X | √ |  | 
| tags_url | string | X | √ |  | 
| downloads_url | string | X | √ |  | 
| hooks_url | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| parent | json | X | √ |  | 
| has_pages | bool | X | √ |  | 
| security_and_analysis | json | X | √ |  | 
| commits_url | string | X | √ |  | 
| stargazers_url | string | X | √ |  | 
| language | string | X | √ |  | 
| assignees_url | string | X | √ |  | 
| branches_url | string | X | √ |  | 
| notifications_url | string | X | √ |  | 
| pulls_url | string | X | √ |  | 
| network_count | int | X | √ |  | 
| has_issues | bool | X | √ |  | 
| visibility | string | X | √ |  | 
| github_teams_selefra_id | string | X | X | fk to github_teams.selefra_id | 
| created_at | timestamp | X | √ |  | 
| team_id | int | X | √ |  | 
| git_url | string | X | √ |  | 
| source | json | X | √ |  | 
| collaborators_url | string | X | √ |  | 
| text_matches | json | X | √ |  | 
| name | string | X | √ |  | 
| issue_comment_url | string | X | √ |  | 
| allow_update_branch | bool | X | √ |  | 
| allow_forking | bool | X | √ |  | 
| license | json | X | √ |  | 
| has_downloads | bool | X | √ |  | 
| git_commits_url | string | X | √ |  | 
| subscription_url | string | X | √ |  | 
| pushed_at | timestamp | X | √ |  | 
| ssh_url | string | X | √ |  | 
| svn_url | string | X | √ |  | 
| permissions | json | X | √ |  | 
| private | bool | X | √ |  | 
| keys_url | string | X | √ |  | 
| languages_url | string | X | √ |  | 
| homepage | string | X | √ |  | 
| auto_init | bool | X | √ |  | 
| organization | json | X | √ |  | 
| allow_rebase_merge | bool | X | √ |  | 
| allow_merge_commit | bool | X | √ |  | 
| disabled | bool | X | √ |  | 
| html_url | string | X | √ |  | 
| watchers | int | X | √ |  | 
| comments_url | string | X | √ |  | 
| merges_url | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| license_template | string | X | √ |  | 
| contents_url | string | X | √ |  | 
| owner | json | X | √ |  | 
| master_branch | string | X | √ |  | 
| forks_count | int | X | √ |  | 
| size | int | X | √ |  | 
| template_repository | json | X | √ |  | 
| archived | bool | X | √ |  | 
| contributors_url | string | X | √ |  | 
| git_refs_url | string | X | √ |  | 
| issues_url | string | X | √ |  | 
| releases_url | string | X | √ |  | 
| trees_url | string | X | √ |  | 
| teams_url | string | X | √ |  | 
| id | int | X | √ |  | 
| full_name | string | X | √ |  | 
| fork | bool | X | √ |  | 
| stargazers_count | int | X | √ |  | 
| has_projects | bool | X | √ |  | 
| forks_url | string | X | √ |  | 
| git_tags_url | string | X | √ |  | 
| labels_url | string | X | √ |  | 
| watchers_count | int | X | √ |  | 
| allow_squash_merge | bool | X | √ |  | 
| use_squash_pr_title_as_default | bool | X | √ |  | 
| is_template | bool | X | √ |  | 
| gitignore_template | string | X | √ |  | 
| archive_url | string | X | √ |  | 
| node_id | string | X | √ |  | 
| subscribers_count | int | X | √ |  | 
| has_wiki | bool | X | √ |  | 
| role_name | string | X | √ |  | 


