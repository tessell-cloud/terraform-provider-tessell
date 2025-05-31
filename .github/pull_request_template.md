## Mention the type of change (Select one or more)?
- [ ] Improvement
- [ ] Bug Fix
- [ ] CI Change
- [ ] Others


> [!TIP]
> * Select `Improvement` if you are adding a new capability, however small it could be.
> * Select `Bug Fix` if you are fixing a broken capability. It could be regression or a new bug.
> * Select `CI Change` when changes are w.r.t. CI/CD (typically in `.github` folder)
> * If the change is anything not covering these three option, use `Others`

> [!NOTE]
> When selecting `Others`, add all relevant details to help reviewer get the maximum understanding of change, its impact and testing done.

## Improvement Details

> [!TIP]
> Fill this section only if  `Improvement` is selected in above section.
>

### Please share the top level JIRA
<jira | Not Applicable>

> [!NOTE]
> If this PR is part of a bigger project, specify the top level jira number. Otherwise specify `Not Applicable`

### Please share reference to requirement document having sign-off from stake holders
<Google Doc|Confluence>


### Mention Known Limitations or Assumptions
<Description | None>


## Bug Fix Details 

> [!TIP]
> Fill this section only if  `Bug Fix` is selected in above section.
>
### In which environment the issue is found?
- [ ] Development
- [ ] QA
- [ ] Stage
- [ ] Production

### PR that introduced the bug.
<PR | Not Applicable>

> [!TIP]
> A bug is always introduced by a code, so use `git blame` to find how it was introduced. 
> 
> This will also help introspect how it was missed and improve the code quality.
>

### Explain what caused this failure. 
*Add details here*

### Mention Known Limitations or Assumptions
<Description | Not Applicable>


## DB Schema changes

> [!TIP]
> Tick the option, and fill in the following questions if this PR
> has changes related to DB Schema

- [x] This PR has DB Schema changes

> [!WARNING]
> Review your changes to be backward compatible and select this option.
> If this option is not selected, you need approval from [@tessell-merge-approvers](https://github.com/orgs/TessellDevelopment/teams/tessell-merge-approvers)
> 

- [ ] The changes are backward compatible.

### Please share a github gist for the reverting the changes in case of failure?
> [!NOTE]
> Schema changes update data or the data-definition. In case of an issue, we must have
> a way to rollback the changes into the older compatible version.
> While schema can be reverted by adding one more flyway version file (mostly), for data
> we might need custom logic. 
> Have the change ready here such that any failures, it is easy to rollback.
> Make sure you test the rollback use-cases.
> 
<GithubGist>

### Added appropriate DB indexes as applicable

> [!WARNING]
> If the answer is `No`, you need approval from [@tessell-merge-approvers](https://github.com/orgs/TessellDevelopment/teams/tessell-merge-approvers)
> 
<Yes|No|Not Applicable>

### Added Flyway migration is idempotent
> [!WARNING]
> If the answer is `No`, you need approval from [@tessell-merge-approvers](https://github.com/orgs/TessellDevelopment/teams/tessell-merge-approvers)
> 
<Yes|No|Not Applicable>


## API changes

- [x] This pull request has related change in `tessell-api-specifications` repository

### Please share the PR from `tessell-api-specifications` repository
<Github | Not Applicable>


> [!WARNING]
> If the answer is `No`, you need approval from [@tessell-merge-approvers](https://github.com/orgs/TessellDevelopment/teams/tessell-merge-approvers)
> 

- [ ] The API specification is additive only and does not break API backward compatibility
- [ ] The API specification change is NOT additive, but does not break API backward compatibility.
- [ ] The API specification change does not break older versions of terraform provider. 


### If not additive change, Please share the list of PRs for all consumers of the API?
<List of Github PRs>



### TessellOps Compatibility

- [x] The PR has changes that will be consumed by `tessellops`
- [ ] The changes are backward compatible with `tessellops`

### Share the PR from tessellops which takes care of fixing it.
<PR>


## Deployment Changes

> [!NOTE]
> When selected, approval from [@devops-eng](https://github.com/orgs/TessellDevelopment/teams/devops-eng)

- [x] This change requires change in the deployment specification (k8s manifest, etc)
### If yes, please share the reference of the PR taking care of the changes
<Github>


- [x] This change require a custom release-specific update to be run during upgrade.
### If yes, please share the reference of the PR taking care of the changes
<Github>

- [ ] Is there any deployment related change which cannot be taken care in release-specific script and need to run manually
### If yes,  please share the confluence page explaining the steps 
<Confluence>

# If Yes, please share the confluence page explaining the steps 
<Confluence>


## Coding Best Practices Checklist

- [ ] The changes have sufficient logging
- [ ] The changes have sufficient code-comments
- [ ] Appropriate alerts are raised wherever applicable


## Validation / Testing 

### How was this change tested (List the use-cases tested) ?
> [!TIP]
> * Mention scenarios that are covered by unit-testing
> * Explain how was integration testing performed.
> * Explain how upgrade testing was done.

<Description>

- [ ] Upgrade testing is successful in dev environment?
<Yes|No>

### Please share the convoy link for the regression test run
<ConvoyUrl | Not Required>

### Please share the sheet where testcases are added for this change.
> [!TIP]
> 
> The sheet need not be a new sheet everytime, if you are adding new 
> testcases to existing feature, share the link to the respective document
> 
<GoogleSheet | Not Available>

### Please share the pull request for the automation test code for this change
> [!WARNING]
> If PR not available, this PR will need approval from [@tessell-qa-pr-reviewers](https://github.com/orgs/TessellDevelopment/teams/tessell-qa-pr-reviewers) team.

<GithubPR | PR Not Available>
