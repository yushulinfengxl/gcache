### …or create a new repository on the command line

```bash
echo "# go-ydb" >> README.md
git init
git add README.md
git commit -m "first commit"
git tag -a v1.0.3-test -m "1.0.3-test 仅用于测试"
git branch -M main
git remote add origin git@github.com:yushulinfengxl/gcache.git
git remote set-url origin git@github.com:yushulinfengxl/gcache.git
git push -u origin main
git push -u origin main v1.0.3-test
```

### …or push an existing repository from the command line

```bash
git remote add origin git@github.com:yushulinfengxl/gcache.git
git branch -M main
git push -u origin main
```