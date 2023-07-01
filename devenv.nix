{...}: {
  languages.go.enable = true;

  pre-commit = {
    excludes = ["tools/"];
    hooks = {
      gofmt.enable = true;
      gotest.enable = true;
      govet.enable = true;
      revive.enable = true;
      staticcheck.enable = true;
    };
  };
}
