pre-commit: Check the commit message for spelling errors.
pre-receive: Enforce project coding standards.
post-commit: Email/SMS team members of a new commit.
post-receive: Push the code to production.

hooks files must be executable
```
chmod +x prepare-commit-msg
```

To use global hook scripts place them all outside of your repositories and then point Git at this new folder.

$ git config --global core.hooksPath /path/to/global/hooks