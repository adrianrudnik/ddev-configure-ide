# ddev-configure-ide

Simple CLI tool to configure PHPStorm or JetBrains based IDEs against a running [DDEV](https://ddev.com/) project.

Written as a quick workaround for [problems](https://github.com/php-perfect/ddev-intellij-plugin/issues/307)
with the actual [IDE plugin](https://plugins.jetbrains.com/plugin/18813-ddev-integration).

Features supported so far:

- Reconfigure the local DDEV db connection configuration

## How to use

Compile yourself or download one of the precompiled binaries from the [releases](https://github.com/adrianrudnik/ddev-configure-ide/releases) page.

Place the executable in your OS executable path or invoke it with an absolute path.

For Windows may need to append `.exe` to the executable name.

```shell
cd your-project

# Test it on a current project, will print the updated config file to the console
ddev-configure-ide jetbrains autoconfig --dry-run

# Configure the current working directory
ddev-configure-ide jetbrains autoconfig

# Configure a different DDEV project root directory
ddev-configure-ide jetbrains autoconfig --root-path=/home/example/code/myproject

# Help can be shown at any level like this
ddev-configure-ide --help
ddev-configure-ide jetbrains autoconfig --help 
```

Hook it into your DDEV project, inside a custom `.ddev/config.local.yaml` like this:

```yaml
hooks:
  post-start:
    - exec-host: "ddev-configure-ide jetbrains autoconfig"
```

Check the output for any errors.
