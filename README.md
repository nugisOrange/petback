# petback
pembuatan package go untuk tugas chapter05


go get -u all
go mod tidy
git tag                                 #check current version
git tag v0.0.28                          #set tag version
git push origin --tags                  #push tag version to repo
go list -m github.com/petapedia/peda@v0.0.28   #publish to pkg dev, replace ORG/URL with your repo URL
