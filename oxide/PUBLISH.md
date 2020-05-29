## Internal publish mechanism

```bash
$ npm install -g vsce
```

```bash
$ vsce login (publisher name)
$ cd myExtension
$ vsce package
# myExtension.vsix generated
$ vsce publish
# <publisherID>.myExtension published to VS Code MarketPlace
```

To revoke token see:
* [dev.azure.com](https://dev.azure.com/)
* [guide](https://code.visualstudio.com/api/working-with-extensions/publishing-extension)

```bash
$ vsce publish minor
$ vsce publish 2.0.1
```
