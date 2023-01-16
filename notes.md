
> Link to video for more info: https://youtu.be/ny0xLzhvADk

List of command uses in this video:
```bash
chmod +x prepare-commit-msg # make git hook file executable
```

To use global hook scripts, place them all outside of your repositories and then point Git ti the new directory by this command:
```bash
git config --global core.hooksPath ~/Downloads/githooks/hooks # Set global Git hooks in your system for all repo
git config --edit --global # view your Git global config
```

__Note:__ All files in the `hook_example` directory is what I used in the [tutorial video](https://youtu.be/ny0xLzhvADk). If you want to test them, you need to replace them with your repository Git hooks in `.git/hook`

> Look into other branches for more examples through video