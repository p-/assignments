## assignmentctl bundle

Bundles an assignment with all additional files inside the assignment's directory

### Synopsis

Bundling compiles all files relevant for an assignment into an archive
format. The backend defaults to zip, but can be set to tarball by
passing the --tar flag. If you want to use tar and gzip, use --gzip.

By default, every bundle includes at least the assignment's PDF from
the ./dist/ directory. If you want to add further files or directories
see the list .spec.bundle.include in your configuration file. It lets
you specify files explicitly, or a glob pattern for multiple files,
e.g. "code/_" or "figures/_.pdf". It is meant to complement the list
of directories to create when using the generate command. The bundle will
preserve the structure of the files included, and will have the PDF
located at the archive's root.

You can customize how the filename for the archive is generated. For this,
you can set .spec.bundle.template to be an arbitrary Golang text template
(including the use of sprig text functions). Just note that this is
limited by what file paths are supported by your operating system, so
don't get too crazy. The map in .spec.bundle.data is passed down to
the template's execution for data binding.

The default archive template is "assignment-{{._id}}.{{._format}}". Note
the \_id field: this is internally augmented from the command's arguments
or the configuration's status field (or, in case of usage of --all, all
available assignments in the repository). "format" is derived from the
selected backend's common file extension, but respects overrides from
the map at .spec.bundle.data, so you can also pick your own file extension
without overriding the entire template.

```text
assignmentctl bundle [flags]
```

### Options

```text
  -a, --all     Bundle all assignments
  -f, --force   Override any existing archives with the same name
      --gzip    Use gzip to encode the archive. Requires --tar to be specified as well
  -h, --help    help for bundle
      --tar     Use tar as a backend for archive bundling
```

### Options inherited from parent commands

```text
  -v, --verbose   Sets logging verbosity level to high
```

### SEE ALSO

- [assignmentctl](assignmentctl.md) - assignments CLI for conveniently templating, building, and bundling course assignment

###### Auto generated by spf13/cobra on 21-Sep-2022