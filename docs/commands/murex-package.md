# _murex_ Shell Docs

## Command Reference: `murex-package`

> _murex_'s package manager

## Description

_murex_ comes with it's own package manager to make managing plugins easier.

The format of the packages is a directory, typically located at `~/.murex_modules`,
which contains one or more murex scripts. Each script can be it's own module.
ie there are multiple modules that can be grouped together and distributed as a
single package.

The way packages and modules are represented is as a path:
    
    package/module
        
`murex-package` is a package management tool for administrating murex modules
and packages.

## Usage

Install a new package

    murex-package: install uri -> <stdout>
    
Update all packages

    murex-package: update -> <stdout>
    
Enable a package or module which had been disabled

    murex-package: enable package
    
    murex-package: enable package/module
    
Disable a package

    murex-package: disable package
    
    murex-package: disable package/module
    
Import packages from another package database

    murex-package: import [ uri/ | local/path/ ]packages.json -> <stdout>
    
Check status of murex packages

    murex-package: status -> <stdout>

## Flags

* `cd`
    Changes working directory to a package's install location
* `disable`
    Disables a previously enabled package or module
* `enable`
    Enables a previously disabled package or module
* `import`
    Import packages described in a backup package DB from user defined URI or local path
* `install`
    Installs a package from a user defined URI
* `list`
    Returns a list of indexed packages/modules (eg what's enabled or disabled)
* `reload`
    Reloads all enabled modules
* `status`
    Returns the version status of locally installed packages
* `update`
    Updates all installed packages

## Detail

### `murex-package: list`... `enabled` vs `loaded`

`enabled` and `disabled` reads the package status from disk rather than the
package cache in your current _murex_ session (like `runtime` reports). This
because the typical use for `murex-package list enabled|disabled` is to view
which packages and modules will be loaded with any new murex session.

If you wish to view what modules are loaded in a current session then use
`murex-package list loaded` instead. This is also equivalent to using
`runtime --modules`.

## See Also

* [Modules and Packages](../user-guide/modules.md):
  An introduction to _murex_ modules and packages
* [_murex_ Profile Files](../user-guide/profile.md):
  A breakdown of the different files loaded on start up
* [`alias`](../commands/alias.md):
  Create an alias for a command
* [`config`](../commands/config.md):
  Query or define _murex_ runtime settings
* [`function`](../commands/function.md):
  Define a function block
* [`murex-docs`](../commands/murex-docs.md):
  Displays the man pages for _murex_ builtins
* [`private`](../commands/private.md):
  Define a private function block
* [`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_