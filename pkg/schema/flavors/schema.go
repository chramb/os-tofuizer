package flavors

//go:generate bash -c "gomodifytags -file $GOPATH/pkg/mod/github.com/gophercloud/gophercloud*/openstack/compute/v2/flavors/results.go  -struct Flavor -add-tags hello | sed -n -e '/type Flavor struct/,/^}/ p' - > flavors.gen.go"
