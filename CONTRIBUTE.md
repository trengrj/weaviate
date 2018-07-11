### Thanks for looking into contributing to Weaviate!
Contributing works pretty easy. You can do a pull request or you can commit if you are part of a Weaviate team.

### How we use Gitflow
How we use [Gitflow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow) and how you can contribute following a few steps.

- The master branch is what is released.
- You can create a feature-branch that is named: feature/YOUR-FEATURE-NAME.
- Your feature branch always has the develop branch as a starting point.
- When you are done you can merge your feature into the develop branch _or_ you can request a merge.
- The master branch is protected.

### Tagging your commit

Always add a refference to your issue to your git commit.

For example: `gh-100: This is the commit message`

AKA: smart commits

### Pull Request

If you create a pull request without smart commits, the pull request will be [squashed into](https://blog.github.com/2016-04-01-squash-your-commits/) one git commit.

### Running Weaviate without database

If you work on Weaviate but not need a database. You can run Weaviate like this: `./cmd/weaviate-server/main.go --scheme=http --port=8080 --host=127.0.0.1 --config="dummy"`
