## assignmentctl ci release

Creates an ENV file to source variables for the selected SCM provider

### Synopsis

The command is meant for usage inside CI pipelines to create release objects
for Gitlab and exports several environment variables in a file that are
required for the job running Gitlab's release-cli or Github's CLI, respectively.

You can run this command outside of CI pipelines, note however that it is highly
dependant on the ENV variables being available in the runner context. You will
have to provide either $CI_COMMIT_TAG with $CI_JOB_ID and $CI_PROJECT_URL or
$GITHUB_REF_NAME for the command to output correct .env files.

```text
assignmentctl ci release [flags]
```

### Options

```text
  -f, --file string   Write the template directly to a file instead of Stdout
  -h, --help          help for release
```

### Options inherited from parent commands

```text
  -v, --verbose   Sets logging verbosity level to high
```

### SEE ALSO

- [assignmentctl ci](assignmentctl_ci.md) - Manage CI integration for assignmentctl

###### Auto generated by spf13/cobra on 21-Sep-2022